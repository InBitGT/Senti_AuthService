// internal/config/config.go
package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	JwtSecret       []byte
	JwtExpiry       time.Duration
	RefreshTokenTTL time.Duration
	OTPTTL          time.Duration
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error cargando .env: %v\n", err)
	}

	JwtSecret = []byte(os.Getenv("JWT_SECRET"))

	JwtExpiry, err = time.ParseDuration(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		fmt.Printf("Error parseando JWT_EXPIRES_IN: %v\n", err)
		JwtExpiry = 24 * time.Hour
	}

	RefreshTokenTTL, err = time.ParseDuration(os.Getenv("REFRESH_TOKEN_TTL"))
	if err != nil {
		fmt.Printf("Error parseando REFRESH_TOKEN_TTL: %v\n", err)
		RefreshTokenTTL = 7 * 24 * time.Hour
	}

	OTPTTL, err = time.ParseDuration(os.Getenv("OTP_TTL"))
	if err != nil {
		fmt.Printf("Error parseando OTP_TTL: %v\n", err)
		OTPTTL = 5 * time.Minute
	}
}
