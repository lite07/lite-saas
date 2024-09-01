package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashSha256String(value string) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s:%s", configManager.HashSalt, value)))
	hashValue := h.Sum(nil)

	return hex.EncodeToString(hashValue)
}
