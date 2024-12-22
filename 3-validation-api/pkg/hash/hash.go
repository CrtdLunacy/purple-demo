package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}
