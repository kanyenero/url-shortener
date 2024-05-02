package rdb

import "context"

type UrlRepository struct {
	context context.Context
}

func (r *UrlRepository) GetUrl(hash string) (string, error) {
	return "", nil
}

func (r *UrlRepository) SetUrl(hash string, url string) error {
	return nil
}

func NewUrlRepository(context context.Context) *UrlRepository {
	return &UrlRepository{
		context: context,
	}
}
