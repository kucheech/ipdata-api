package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
