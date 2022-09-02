package database

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

var (
	ErrNotFound error = errors.New("key not found")
)

const HashName string = "urls"

// Set saves url and shorted path to database/cache permanently
func (r *Redis) Set(ctx context.Context, key, value string) error {
	if err := r.db.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	v, err := r.db.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrNotFound
		}
		return "", err
	}

	return v, nil
}
