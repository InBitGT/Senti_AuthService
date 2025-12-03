package auth

type RegisterRequest struct {
	TenantCode string `json:"tenant"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RoleName   string `json:"role"`
}

type LoginRequest struct {
	TenantCode string `json:"tenant"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	OTP        string `json:"otp,omitempty"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int64  `json:"expiresIn"`
}
