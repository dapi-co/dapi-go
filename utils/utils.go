package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"math/rand"
	"net/http"
	"time"

	svix "github.com/svix/svix-webhooks/go"
)

func GenerateRandomString(length int) string {

	result := ""
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}

	return result
}

func DecryptWebhookData(aesKey string, encryptedData string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(aesKey)
	if err != nil {
		return "", err
	}

	encDataBuffer, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	iv := encDataBuffer[:aes.BlockSize]
	encryptedText := encDataBuffer[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(encryptedText))
	mode.XORKeyStream(decrypted, encryptedText)

	return string(decrypted), nil
}

func VerifyWebhook(secret string, payload []byte, headers http.Header) bool {
	wh, _ := svix.NewWebhook(secret)
	if err := wh.Verify(payload, headers); err != nil {
		return false
	}
	return true
}
