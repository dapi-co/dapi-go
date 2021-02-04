package utils

import (
	"math/rand"
	"time"
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
