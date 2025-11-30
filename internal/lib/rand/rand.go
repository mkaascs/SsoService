package rand

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

func GenerateBase64(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than zero")
	}

	result := make([]byte, length)
	if _, err := rand.Read(result); err != nil {
		return "", fmt.Errorf("failed to generate random string: %w", err)
	}

	return base64.RawStdEncoding.EncodeToString(result), nil
}
