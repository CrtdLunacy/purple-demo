package hash

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func generateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func HashString(s string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}
	combined := salt + s
	hash := sha256.Sum256([]byte(combined))
	hashString := hex.EncodeToString(hash[:])

	// Combine the salt and the hash into one string
	combinedResult := salt + ":" + hashString
	return combinedResult, nil
}
