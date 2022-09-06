package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

type AES struct{ cipher cipher.AEAD }

const keyLen = 32

var (
	ErrInvalidKey        = errors.New("invalid_key")
	ErrInvalidCiphertext = errors.New("invalid_ciphertext")
)

func NewAES(key []byte) (*AES, error) {
	if len(key) != keyLen {
		return nil, ErrInvalidKey
	}

	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create block cipher: %w", err)
	}

	g, err := cipher.NewGCM(b)
	if err != nil {
		return nil, fmt.Errorf("failed to create AED cipher: %w", err)
	}

	return &AES{cipher: g}, nil
}

func (a *AES) Encrypt(plaintext []byte) ([]byte, error) {
	nonce, err := RandomBytes(uint(a.cipher.NonceSize()))
	if err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	return a.cipher.Seal(nonce, nonce, plaintext, nil), nil
}

func (a *AES) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < a.cipher.NonceSize() {
		return nil, ErrInvalidCiphertext
	}

	return a.cipher.Open(nil,
		ciphertext[:a.cipher.NonceSize()],
		ciphertext[a.cipher.NonceSize():],
		nil,
	)
}
