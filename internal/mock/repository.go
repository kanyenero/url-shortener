package mock

import (
	"context"
	"errors"
)

type UrlRepository struct {
	context context.Context
}

func (r *UrlRepository) GetUrl(hash string) (string, error) {
	switch hash {
	case "88bc09d7":
		return "http://google.com/?q=golang", nil
	default:
		return "", errors.New("invalid hash")
	}
}

func (r *UrlRepository) SetUrl(hash string, url string) error {
	return nil
}

func NewUrlRepository(context context.Context) *UrlRepository {
	return &UrlRepository{
		context: context,
	}
}
