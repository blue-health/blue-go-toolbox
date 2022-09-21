package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
)

type InMemory struct{ cache *ristretto.Cache }

const bufferItems = 64

var _ Cache = (*InMemory)(nil)

func NewInMemory(numCounters, maxCost int64) (*InMemory, error) {
	c, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: numCounters,
		MaxCost:     maxCost,
		BufferItems: bufferItems,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to initialize cache: %w", err)
	}

	return &InMemory{cache: c}, nil
}

func (c *InMemory) Get(_ context.Context, key string) ([]byte, error) {
	i, f := c.cache.Get(key)
	if !f {
		return nil, nil
	}

	v, ok := i.([]byte)
	if !ok {
		return nil, nil
	}

	return v, nil
}

func (c *InMemory) GetValue(_ context.Context, key string) (any, error) {
	i, f := c.cache.Get(key)
	if !f {
		return nil, nil
	}

	return i, nil
}

func (c *InMemory) Set(_ context.Context, key string, value []byte, expiry time.Duration) error {
	_ = c.cache.SetWithTTL(key, value, 1, expiry)
	return nil
}

func (c *InMemory) Del(_ context.Context, key string) error {
	c.cache.Del(key)
	return nil
}
