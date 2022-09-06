package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/blue-health/blue-go-toolbox/crypto"
)

type Encrypted struct {
	inner Cache
	aes   *crypto.AES
}

var _ Cache = (*Encrypted)(nil)

func NewEncrypted(aes *crypto.AES, inner Cache) *Encrypted {
	return &Encrypted{aes: aes, inner: inner}
}

func (e *Encrypted) Get(ctx context.Context, key string) ([]byte, error) {
	var (
		dec []byte
		err error
	)

	if dec, err = e.inner.Get(ctx, key); err == nil && dec != nil {
		dec, err = e.aes.Decrypt(dec)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt data: %w", err)
		}
	}

	return dec, err
}

func (e *Encrypted) Set(ctx context.Context, key string, value []byte, expiry time.Duration) error {
	b, err := e.aes.Encrypt(value)
	if err != nil {
		return fmt.Errorf("failed to encrypt data: %w", err)
	}

	return e.inner.Set(ctx, key, b, expiry)
}

func (e *Encrypted) Del(ctx context.Context, key string) error {
	return e.inner.Del(ctx, key)
}
