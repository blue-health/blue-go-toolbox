package crypto

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func Hash(b []byte) string {
	h := sha256.New()
	h.Write(b)

	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func HashBytes(bs ...[]byte) string {
	h := sha256.New()

	for _, b := range bs {
		h.Write(b)
	}

	return hex.EncodeToString(h.Sum(nil))
}
