package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

func EncryptString(str string) string {
	aes, err := aes.NewCipher([]byte(configManager.EncryptionSecretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	encryptedString := gcm.Seal(nonce, nonce, []byte(str), nil)

	return base64.StdEncoding.EncodeToString(encryptedString)
}

func DecryptString(encodedStr string) string {
	str, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		panic(err)
	}

	aes, err := aes.NewCipher([]byte(configManager.EncryptionSecretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := str[:nonceSize], str[nonceSize:]

	decryptedText, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		panic(err)
	}

	return string(decryptedText)
}
