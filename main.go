package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BRO3886/url-shortener/api"
	"github.com/BRO3886/url-shortener/config"
	"github.com/BRO3886/url-shortener/repository/cache"
	"github.com/BRO3886/url-shortener/repository/database"
	"github.com/BRO3886/url-shortener/shortener"
	"github.com/gin-gonic/gin"
)

var (
	configFile string
	port       = flag.String("port", "8080", "port to listen on")
	db         = flag.String("db", "redis", "databse to use (redis, mongo)")
)

func init() {
	flag.Parse()
	switch config.Env {
	case "DEPLOY":
		configFile = config.DeployConfigFile
		gin.SetMode(gin.ReleaseMode)
	default:
		configFile = config.LocalConfigFile
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	err := config.NewConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	repo := getRepo()
	svc := shortener.NewRedirectService(repo)
	handler := api.NewHandler(svc)
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.GET("/:code", handler.Redirect)
	r.POST("/", handler.SetRedirect)

	log.Fatal(r.Run(fmt.Sprintf(":%s", *port)))
}

func getRepo() shortener.RedirectRepository {
	switch *db {
	default:
		redisConf := config.Config.Redis
		repo, err := cache.NewRepository(redisConf.Addr, redisConf.UserName, redisConf.Password)
		if err != nil {
			log.Fatal(err)
		}

		return repo
	case "redis":
		redisConf := config.Config.Redis
		repo, err := cache.NewRepository(redisConf.Addr, redisConf.UserName, redisConf.Password)
		if err != nil {
			log.Fatal(err)
		}

		return repo
	case "mongo":
		mongoConf := config.Config.Mongo
		url := fmt.Sprintf("mongodb://%s:%s@%s", mongoConf.User, mongoConf.Password, mongoConf.Addr)
		repo, err := database.NewRepository(url, mongoConf.Database, mongoConf.Timeout)
		if err != nil {
			log.Fatal(err)
		}

		return repo
	}
}
