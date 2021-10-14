package main

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ipdata struct {
	CountryCode string `json:"countrycode"`
	Timezone    string `json:"timezone"`
}

var data = []ipdata{
	{CountryCode: "SG", Timezone: "Asia/Singapore"},
	{CountryCode: "CA", Timezone: "America/Toronto"},
	{CountryCode: "US", Timezone: "America/New_York"},
}

func isValidIp(ip string) bool {
	return net.ParseIP(ip) != nil
}

func getDataByIp(c *gin.Context) {
	ip := c.Query("ip")
	if isValidIp(ip) {
		c.IndentedJSON(http.StatusOK, data[2])
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid ip or data not found"})
	}
}

func main() {
	router := gin.Default()
	// router.GET("/data", getData)
	router.GET("/data", getDataByIp)

	router.Run("localhost:8080")
}
