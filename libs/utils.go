package libs

import (
	"crypto/rand"
	"encoding/hex"
)

// generateRandomID generates a random ID of the specified length.
func GenerateRandomID(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
