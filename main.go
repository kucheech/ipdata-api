package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ipdata struct {
	CountryCode string `json:"countryCode"`
	Timezone    string `json:"timezone"`
}

func fetchDataFromIpApi(ip string) ipdata {
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=countryCode,timezone", ip)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		return ipdata{}
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject ipdata
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}

func isValidIp(ip string) bool {
	return net.ParseIP(ip) != nil
}

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
	router := gin.Default()
	router.GET("/data", getDataByIp)

	router.Run("localhost:8080")
}
