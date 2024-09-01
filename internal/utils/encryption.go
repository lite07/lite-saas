package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

func EncryptString(str string) (string, error) {
	aes, err := aes.NewCipher([]byte(configManager.EncryptionSecretKey))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}

	encryptedString := gcm.Seal(nonce, nonce, []byte(str), nil)

	return base64.StdEncoding.EncodeToString(encryptedString), nil
}

func DecryptString(encodedStr string) (string, error) {
	str, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return "", err
	}

	aes, err := aes.NewCipher([]byte(configManager.EncryptionSecretKey))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := str[:nonceSize], str[nonceSize:]

	decryptedText, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}

	return string(decryptedText), nil
}
