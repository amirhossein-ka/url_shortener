package service

import (
	"context"

	"github.com/amirhossein-ka/url_shortener/database"
)

type Service interface {
	ShortUrlAddress(ctx context.Context, url string) (string, error)
	GetUrlAddr(ctx context.Context, key string) (string, error)
}

type ServiceImpl struct {
	rdb database.Database
}

func New(rdb database.Database) Service {
	return &ServiceImpl{
		rdb: rdb,
	}
}
