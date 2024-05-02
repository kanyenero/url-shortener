package app

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net"
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

	fmt.Println(handlers)
}
