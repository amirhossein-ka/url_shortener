package service

import (
	"context"
)

func (s *ServiceImpl) ShortUrlAddress(ctx context.Context, url string) (string, error) {
	rString := randomString(randomStringLenght)

	if err := s.rdb.Set(ctx, rString, url); err != nil {
		return "", err
	}

	return rString, nil

}

func (s *ServiceImpl) GetUrlAddr(ctx context.Context, key string) (string, error) {
	url, err := s.rdb.Get(ctx, key)
	if err != nil {
		return "", err
	}

	return url, nil
}
