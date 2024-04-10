package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

// GenerateToken generates a random token for the player with length 32
func GenerateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
