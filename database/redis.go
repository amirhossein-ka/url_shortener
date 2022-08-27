package database

import (
	"context"
	"fmt"

	"github.com/amirhossein-ka/url_shortener/config"
	"github.com/go-redis/redis/v8"
)

type Database interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}

type Redis struct {
	db *redis.Client
}

func New(cfg config.Config) (Database, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.DB.Address, cfg.DB.Port),
		DB:   0,
	})
	res, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(res)

	return &Redis{db: rdb}, nil

}
