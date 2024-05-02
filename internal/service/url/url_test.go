package url

import (
	"context"
	"testing"
	"url-shortener/internal/mock"
	"url-shortener/internal/repository"
)

var ctx = context.Background()

func newService(ctx context.Context) *UrlService {
	var repo repository.UrlRepository
	repo = mock.NewUrlRepository(ctx)
	return NewUrlService(&repo, ctx)
}

func TestUrlService_SetUrl(t *testing.T) {
	svc := newService(ctx)

	positiveTestCases := []struct {
		url          string
		expectedHash string
	}{
		{url: "http://google.com/?q=golang", expectedHash: "88bc09d7"},
	}

	negativeTestCases := []struct {
		url          string
		expectedHash string
	}{
		{url: "", expectedHash: ""},
		{url: "not url", expectedHash: ""},
	}

	t.Run("positive", func(t *testing.T) {
		t.Parallel()

		for _, tc := range positiveTestCases {
			actualHash, err := svc.SetUrl(tc.url)
			if err != nil {
				t.Errorf("%v", err)
			}
			if actualHash != tc.expectedHash {
				t.Errorf("expected: %s, actual: %s", tc.expectedHash, actualHash)
			}
		}
	})

	t.Run("negative", func(t *testing.T) {
		t.Parallel()

		for _, tc := range negativeTestCases {
			_, err := svc.SetUrl(tc.url)
			if err == nil {
				t.Error("expected error, but got nil")
			}
		}
	})
}
