package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	possible := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	s := make([]rune, length)
	for i := range s {
		s[i] = possible[rand.Intn(len(possible))]
	}
	return string(s)
}
