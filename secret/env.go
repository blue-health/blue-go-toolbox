package secret

import (
	"context"
	"encoding/base64"
	"errors"
	"os"
)

type EnvSource struct{}

var _ Source = (*EnvSource)(nil)

var ErrSecretNotFound = errors.New("secret_not_found")

func NewEnvSource() *EnvSource { return &EnvSource{} }

func (s *EnvSource) Get(_ context.Context, name string) (Secret, error) {
	if v := os.Getenv(name); v != "" {
		b, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			return []byte(v), nil
		}

		return b, nil
	}

	return nil, ErrSecretNotFound
}

func (s *EnvSource) Close() error { return nil }
