package service

import "github.com/amirhossein-ka/url_shortener/database"

type Service struct {
    rdb *database.Redis
}

func New(rdb *database.Redis) *Service {
    return &Service{
        rdb: rdb,
    }
}
