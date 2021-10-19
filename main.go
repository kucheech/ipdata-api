package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func main() {
	cache := cache.New(5*time.Minute, 10*time.Minute)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET"}
	router.Use(cors.New(config))

	router.GET("/data", CacheCheck(cache), getDataByIp)

	router.Run("localhost:8080")
}
