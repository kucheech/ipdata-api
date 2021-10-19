package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
