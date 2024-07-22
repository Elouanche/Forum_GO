package main

import (
	"crypto/sha256"
	"encoding/hex"
)

// hashPassword hache un mot de passe en utilisant SHA-256
func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
