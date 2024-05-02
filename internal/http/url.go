package http

import (
	"net/http"
	"url-shortener/internal/service"
)

type UrlHandler struct {
	urlService *service.UrlService
}

func CreateUrlHandler(urlService *service.UrlService) *UrlHandler {
	return &UrlHandler{urlService}
}

func (handler *UrlHandler) Get(writer http.ResponseWriter, request *http.Request) {

}

func (handler *UrlHandler) Set(writer http.ResponseWriter, request *http.Request) {

}
