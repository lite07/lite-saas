package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSha256String(value string) string {
	h := sha256.New()
	h.Write([]byte(value))
	hashValue := h.Sum(nil)

	return hex.EncodeToString(hashValue)
}
