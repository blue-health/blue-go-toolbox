package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func RandomBytes(n uint) ([]byte, error) {
	b := make([]byte, n)

	if _, err := rand.Read(b); err != nil {
		return nil, fmt.Errorf("failed to read crypto/rand: %w", err)
	}

	return b, nil
}

func ReadRandomBytes(o []byte) error {
	if _, err := rand.Read(o); err != nil {
		return fmt.Errorf("failed to read crypto/rand: %w", err)
	}

	return nil
}

func RandomString(n uint) (string, error) {
	b, err := RandomBytes(n)
	if err != nil {
		return "", nil
	}

	return hex.EncodeToString(b), nil
}
