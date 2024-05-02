package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"url-shortener/internal/mock"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
	urlService "url-shortener/internal/service/url"
)

var ctx = context.Background()

func TestUrlHandler_Get(t *testing.T) {
	var repo repository.UrlRepository
	repo = mock.NewUrlRepository(ctx)
	var svc service.UrlService
	svc = urlService.NewUrlService(&repo, ctx)
	hdl := NewUrlHandler(&svc)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /s/", hdl.Get)
	srv := httptest.NewServer(mux)
	srv.Config.Addr = ":8081"
	defer srv.Close()

	positiveTestCases := []struct {
		uri            string
		expectedUrl    string
		expectedStatus int
	}{
		{uri: "/s/88bc09d7", expectedUrl: "http://google.com/?q=golang", expectedStatus: http.StatusOK},
	}

	negativeTestCases := []struct {
		uri            string
		expectedUrl    string
		expectedStatus int
	}{
		{uri: "/s/", expectedUrl: "", expectedStatus: http.StatusBadRequest},
		{uri: "/s/88bc09d~", expectedUrl: "", expectedStatus: http.StatusBadRequest},
		{uri: "/s/88bc09d70", expectedUrl: "", expectedStatus: http.StatusBadRequest},
		{uri: "/s/88bc09d7/88bc09d7", expectedUrl: "", expectedStatus: http.StatusBadRequest},
		{uri: "/s/88BC09D7", expectedUrl: "", expectedStatus: http.StatusNotFound},
	}

	t.Run("positive", func(t *testing.T) {
		for i, tc := range positiveTestCases {
			t.Run("positive "+strconv.Itoa(i), func(t *testing.T) {
				result, err := srv.Client().Get(srv.URL + tc.uri)
				if err != nil {
					t.Error(err)
					return
				}
				if result.StatusCode != tc.expectedStatus {
					t.Errorf("expected status %d, got %d", tc.expectedStatus, result.StatusCode)
					return
				}
			})
		}
	})
	t.Run("negative", func(t *testing.T) {
		for i, tc := range negativeTestCases {
			t.Run("negative "+strconv.Itoa(i), func(t *testing.T) {
				result, err := srv.Client().Get(srv.URL + tc.uri)
				if err != nil {
					t.Error(err)
					return
				}
				if result.StatusCode != tc.expectedStatus {
					t.Errorf("expected status %d, got %d", tc.expectedStatus, result.StatusCode)
					return
				}
			})
		}
	})
}

func TestUrlHandler_Set(t *testing.T) {
	var repo repository.UrlRepository
	repo = mock.NewUrlRepository(ctx)
	var svc service.UrlService
	svc = urlService.NewUrlService(&repo, ctx)
	hdl := NewUrlHandler(&svc)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /a/", hdl.Set)
	srv := httptest.NewServer(mux)
	srv.Config.Addr = ":8080"
	defer srv.Close()

	positiveTestCases := []struct {
		uri            string
		expectedStatus int
	}{
		{uri: "/a/?url=http%3A%2F%2Fgoogle.com%2F%3Fq%3Dgolang", expectedStatus: http.StatusOK},
	}

	negativeTestCases := []struct {
		uri            string
		expectedStatus int
	}{
		{uri: "/a/", expectedStatus: http.StatusBadRequest},
		{uri: "/a/subdir/?url=http%3A%2F%2Fgoogle.com%2F%3Fq%3Dgolang", expectedStatus: http.StatusBadRequest},
		{uri: "/a/?noturl=http%3A%2F%2Fgoogle.com%2F%3Fq%3Dgolang", expectedStatus: http.StatusBadRequest},
	}

	t.Run("positive", func(t *testing.T) {
		for i, tc := range positiveTestCases {
			t.Run("positive "+strconv.Itoa(i), func(t *testing.T) {
				result, err := srv.Client().Get(srv.URL + tc.uri)
				if err != nil {
					t.Error(err)
					return
				}
				if result.StatusCode != tc.expectedStatus {
					t.Errorf("expected status %d, got %d", tc.expectedStatus, result.StatusCode)
					return
				}
			})
		}
	})
	t.Run("negative", func(t *testing.T) {
		for i, tc := range negativeTestCases {
			t.Run("negative "+strconv.Itoa(i), func(t *testing.T) {
				result, err := srv.Client().Get(srv.URL + tc.uri)
				if err != nil {
					t.Error(err)
					return
				}
				if result.StatusCode != tc.expectedStatus {
					t.Errorf("expected status %d, got %d", tc.expectedStatus, result.StatusCode)
					return
				}
			})
		}
	})
}
