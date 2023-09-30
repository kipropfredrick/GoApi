package util

import (
	"fmt"
	"math/rand"
	// "os"
	"time"

	"golang.org/x/crypto/bcrypt"
)


var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))
//helper methods
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func RanndStrin(length int) int {
  return rand.Int()
}