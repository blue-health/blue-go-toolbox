package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
)

type HMAC struct{ key []byte }

func NewHMAC(key []byte) *HMAC { return &HMAC{key: key} }

func (h *HMAC) Generate(message []byte) []byte {
	m := hmac.New(sha256.New, h.key)
	m.Write(message)

	return m.Sum(nil)
}

func (h *HMAC) Validate(message, digest []byte) bool {
	return hmac.Equal(digest, h.Generate(message))
}
