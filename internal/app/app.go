package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net"
	"net/http"
	"time"
	"url-shortener/internal/config"
	httpHandlers "url-shortener/internal/http"
	"url-shortener/internal/repository"
	redisDbProvider "url-shortener/internal/repository/rdb"
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
	case "redis":
		redisClient := redis.NewClient(&redis.Options{
			Addr:     net.JoinHostPort(cfg.Redis.Host, cfg.Redis.Port),
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.Db,
		})

		repositories = repository.Repositories{
			UrlRepository: redisDbProvider.NewUrlRepository(redisClient, ctx),
		}
	default:
		log.Fatalf(fmt.Sprintf("Database '%s' is not supported.", cfg.Database))
	}

	services := service.Services{
		UrlService: urlService.NewUrlService(&repositories.UrlRepository, ctx),
	}

	handlers := httpHandlers.Handlers{
		UrlHandler: *httpHandlers.CreateUrlHandler(&services.UrlService),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /a/", handlers.UrlHandler.Set)
	mux.HandleFunc("GET /s/", handlers.UrlHandler.Get)

	httpServer := http.Server{
		Addr:         ":" + cfg.Http.Port,
		Handler:      mux,
		ReadTimeout:  time.Duration(cfg.Http.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Http.WriteTimeout) * time.Second,
	}

	if err = httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("%s", err.Error())
	}
}
