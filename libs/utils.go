package libs

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/fusionxx23/ecommerce-go/models"
	"github.com/golang-jwt/jwt"
)

// generateRandomID generates a random ID of the specified length.
func GenerateRandomID(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

var (
	key []byte
)

func ParseJWT(signedString string, claims jwt.MapClaims) (*jwt.Token, error) {
	key = []byte("your-256-bit-secret") // Replace with your actual secret key
	a, err := jwt.ParseWithClaims(signedString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	return a, err
}

func CreateJWT(user models.User) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "auth",
			"sub": user.Email,                                 // subject of the token
			"exp": time.Now().Add(86400 * time.Second).Unix(), // 24 hours expiry
		})
	s, err := t.SignedString(key)
	return s, err
}
