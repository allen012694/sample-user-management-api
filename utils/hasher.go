package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256WithSalt(text string, salt string) string {
	hashed := sha256.Sum256([]byte(salt + text))
	return hex.EncodeToString(hashed[:])
}
