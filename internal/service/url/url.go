package url

import (
	"context"
	"errors"
	"url-shortener/internal/repository"
)

type UrlService struct {
	repository *repository.UrlRepository
	context    context.Context
}

func (service *UrlService) GetUrl(hash string) (string, error) {
	if hash == "" {
		return "", errors.New("hash is empty")
	}

	return (*service.repository).GetUrl(hash)
}

func (service *UrlService) SetUrl(url string) (string, error) {
	if url == "" {
		return "", errors.New("url is empty")
	}

	hash := ""
	err := (*service.repository).SetUrl(hash, url)
	return hash, err
}

func NewUrlService(repository *repository.UrlRepository, context context.Context) *UrlService {
	return &UrlService{
		repository: repository,
		context:    context,
	}
}
