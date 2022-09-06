package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(context.Context, string) ([]byte, error)
	Set(context.Context, string, []byte, time.Duration) error
	Del(context.Context, string) error
}

const (
	DefaultExpiry = time.Minute * 30
)
