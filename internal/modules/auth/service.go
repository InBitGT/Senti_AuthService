package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"AuthService/internal/config"
	"AuthService/internal/modules/otp"
	"AuthService/internal/modules/refreshtoken"
	"AuthService/internal/modules/user"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("credenciales inválidas")
	ErrInvalidOTP         = errors.New("otp inválido o expirado")
)

type AuthService interface {
	Register(req *RegisterRequest) (*user.User, error)
	Login(req *LoginRequest) (*AuthResponse, error)
	Refresh(token string) (*AuthResponse, error)

	GenerateOTP(tenantCode, email, channel string) (*otp.UserOTP, error)
	ToggleTwoFA(userID uint, enabled bool) error
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{repo}
}

type AuthClaims struct {
	UserID      uint     `json:"sub"`
	TenantID    uint     `json:"tenant"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

func (s *authService) Register(req *RegisterRequest) (*user.User, error) {
	tenantObj, err := s.repo.GetActiveTenantByCode(req.TenantCode)
	if err != nil {
		return nil, err
	}

	roleObj, err := s.repo.GetRoleByName(tenantObj.ID, req.RoleName)
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		TenantID:     tenantObj.ID,
		Email:        req.Email,
		PasswordHash: string(hash),
		RoleID:       roleObj.ID,
		IsActive:     true,
	}

	if err := s.repo.CreateUser(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *authService) Login(req *LoginRequest) (*AuthResponse, error) {
	tenantObj, err := s.repo.GetActiveTenantByCode(req.TenantCode)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	u, err := s.repo.GetUserByEmail(tenantObj.ID, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	if u.TwoFAEnabled {
		if req.OTP == "" {
			return nil, ErrInvalidOTP
		}

		otpObj, err := s.repo.FindValidOTP(u.ID, u.TenantID, req.OTP, time.Now())
		if err != nil {
			return nil, ErrInvalidOTP
		}
		_ = s.repo.MarkOTPUsed(otpObj.ID)
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

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenStr,
		ExpiresIn:    int64(config.JwtExpiry.Seconds()),
	}, nil
}

func (s *authService) generateTokens(u *user.User, perms []string) (string, string, error) {
	now := time.Now()

	claims := AuthClaims{
		UserID:      u.ID,
		TenantID:    u.TenantID,
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

	rt := &refreshtoken.RefreshToken{
		UserID:    u.ID,
		Token:     generateRandomToken(),
		ExpiresAt: now.Add(config.RefreshTokenTTL),
	}

	if err := s.repo.SaveRefreshToken(rt); err != nil {
		return "", "", err
	}

	return accessToken, rt.Token, nil
}

func (s *authService) Refresh(token string) (*AuthResponse, error) {
	rt, err := s.repo.GetRefreshToken(token)
	if err != nil || rt.Revoked || rt.ExpiresAt.Before(time.Now()) {
		return nil, ErrInvalidCredentials
	}

	u, err := s.repo.GetUserByID(rt.UserID)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

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

/* OTP */

func generateOTP() string {
	n := rand.Intn(899999) + 100000
	return fmt.Sprintf("%06d", n)
}

func (s *authService) GenerateOTP(tenantCode, email, channel string) (*otp.UserOTP, error) {
	t, err := s.repo.GetActiveTenantByCode(tenantCode)
	if err != nil {
		return nil, errors.New("tenant inválido")
	}

	u, err := s.repo.GetUserByEmail(t.ID, email)
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	code := generateOTP()

	o := &otp.UserOTP{
		UserID:    u.ID,
		TenantID:  t.ID,
		Code:      code,
		Channel:   channel,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	if err := s.repo.SaveOTP(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (s *authService) ToggleTwoFA(userID uint, enabled bool) error {
	u, err := s.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	u.TwoFAEnabled = enabled
	return s.repo.UpdateUser(u)
}
