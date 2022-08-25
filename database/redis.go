package database

import (
	"context"
	"fmt"

	"github.com/amirhossein-ka/url_shortener/config"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	db  *redis.Client
}

func New(cfg config.Config) (*Redis, error) {
    rdb := redis.NewClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%s",cfg.DB.Address, cfg.DB.Port),
        DB: 0,
    })
    res,err := rdb.Ping(context.Background()).Result()
    if err != nil {
        return nil, err
    }
    fmt.Println(res)

    return &Redis{db: rdb}, nil
    
}
