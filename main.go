package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"time"
)

func getDataByIp(c *gin.Context) {
	ip := c.Query("ip")
	if isValidIp(ip) {
		result := fetchDataFromIpApi(ip)
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid ip or data not found"})
	}
}

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
