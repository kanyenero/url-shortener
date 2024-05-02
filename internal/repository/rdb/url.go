package rdb

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type UrlRepository struct {
	client  *redis.Client
	context context.Context
}

func (r *UrlRepository) GetUrl(hash string) (string, error) {
	return r.client.Get(r.context, hash).Result()
}

func (r *UrlRepository) SetUrl(hash string, url string) error {
	return r.client.Set(r.context, hash, url, 0).Err()
}

func NewUrlRepository(client *redis.Client, context context.Context) *UrlRepository {
	return &UrlRepository{client: client, context: context}
}
