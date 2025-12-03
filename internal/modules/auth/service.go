package auth

import (
	"errors"
	"time"

	"AuthService/internal/config" // ajusta el módulo: authService/internal/config

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("credenciales inválidas")
	ErrInvalidOTP         = errors.New("otp inválido o expirado")
)

type AuthService interface {
	Register(req *RegisterRequest) (*User, error)
	Login(req *LoginRequest) (*AuthResponse, error)
	Refresh(refreshToken string) (*AuthResponse, error)
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{repo}
}

// claims personalizados
type AuthClaims struct {
	UserID      uint     `json:"sub"`
	TenantID    uint     `json:"tenant"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

func (s *authService) Register(req *RegisterRequest) (*User, error) {
	tenant, err := s.repo.GetActiveTenantByCode(req.TenantCode)
	if err != nil {
		return nil, err
	}

	role, err := s.repo.GetRoleByName(tenant.ID, req.RoleName)
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		TenantID:     tenant.ID,
		Email:        req.Email,
		PasswordHash: string(hash),
		RoleID:       role.ID,
		IsActive:     true,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(req *LoginRequest) (*AuthResponse, error) {
	tenant, err := s.repo.GetActiveTenantByCode(req.TenantCode)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	u, err := s.repo.GetUserByEmail(tenant.ID, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	// si tiene 2FA habilitado, validamos OTP
	if u.TwoFAEnabled {
		if req.OTP == "" {
			return nil, ErrInvalidOTP
		}
		otp, err := s.repo.FindValidOTP(u.ID, u.TenantID, req.OTP, time.Now())
		if err != nil {
			return nil, ErrInvalidOTP
		}
		_ = s.repo.MarkOTPUsed(otp.ID)
	}

	perms, err := s.repo.GetPermissionsByRole(u.RoleID)
	if err != nil {
		return nil, err
	}
	permKeys := make([]string, len(perms))
	for i, p := range perms {
		permKeys[i] = p.Key
	}

	accessToken, refreshTokenStr, err := s.generateTokens(u, permKeys)
	if err != nil {
		return nil, err
	}

	resp := &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenStr,
		ExpiresIn:    int64(config.JwtExpiry.Seconds()),
	}
	return resp, nil
}

func (s *authService) generateTokens(u *User, perms []string) (string, string, error) {
	now := time.Now()

	claims := AuthClaims{
		UserID:      u.ID,
		TenantID:    u.TenantID,
		Role:        "",
		Permissions: perms,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(config.JwtExpiry)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken := &RefreshToken{
		UserID:    u.ID,
		Token:     generateRandomToken(), // implementa un generador aleatorio
		ExpiresAt: now.Add(config.RefreshTokenTTL),
	}
	if err := s.repo.SaveRefreshToken(refreshToken); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken.Token, nil
}

func (s *authService) Refresh(refreshTokenStr string) (*AuthResponse, error) {
	rt, err := s.repo.GetRefreshToken(refreshTokenStr)
	if err != nil || rt.ExpiresAt.Before(time.Now()) || rt.Revoked {
		return nil, ErrInvalidCredentials
	}

	// cargar usuario
	u := &User{}
	u.ID = rt.UserID

	perms, err := s.repo.GetPermissionsByRole(u.RoleID)
	if err != nil {
		return nil, err
	}
	permKeys := make([]string, len(perms))
	for i, p := range perms {
		permKeys[i] = p.Key
	}

	accessToken, newRefresh, err := s.generateTokens(u, permKeys)
	if err != nil {
		return nil, err
	}

	_ = s.repo.RevokeRefreshToken(rt.ID)

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefresh,
		ExpiresIn:    int64(config.JwtExpiry.Seconds()),
	}, nil
}
