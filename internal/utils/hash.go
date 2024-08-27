package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func HashSha256String(value string) string {
	salt := os.Getenv("HASH_SALT")
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s:%s", salt, value)))
	hashValue := h.Sum(nil)

	return hex.EncodeToString(hashValue)
}
