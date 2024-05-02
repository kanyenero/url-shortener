package url

import (
	"context"
	"encoding/hex"
	"errors"
	"hash/adler32"
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

	hasher := adler32.New()
	_, err := hasher.Write([]byte(url))
	if err != nil {
		return "", err
	}

	hash := hex.EncodeToString(hasher.Sum(nil))
	err = (*service.repository).SetUrl(hash, url)
	return hash, err
}

func NewUrlService(repository *repository.UrlRepository, context context.Context) *UrlService {
	return &UrlService{
		repository: repository,
		context:    context,
	}
}
