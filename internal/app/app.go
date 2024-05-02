package app

import (
	"context"
	"fmt"
	"log"
	"url-shortener/internal/config"
	httpHandlers "url-shortener/internal/http"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
	urlService "url-shortener/internal/service/url"
)

var ctx = context.Background()

func Run(configPath string) {
	cfg, err := config.Initialize(configPath)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	var repositories repository.Repositories
	switch cfg.Database {
	case "":
	default:
		log.Fatalf(fmt.Sprintf("Database '%s' is not supported.", cfg.Database))
	}

	services := service.Services{
		UrlService: urlService.NewUrlService(&repositories.UrlRepository, ctx),
	}

	handlers := httpHandlers.Handlers{
		UrlHandler: *httpHandlers.CreateUrlHandler(&services.UrlService),
	}

	fmt.Println(handlers)
}
