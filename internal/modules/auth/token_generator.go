package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func generateRandomToken() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
