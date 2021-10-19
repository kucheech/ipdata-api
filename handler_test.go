// https://craig-childs.medium.com/testing-gin-json-responses-1f258ce3b0b1
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGetDataByIp(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Build our expected body
	body := gin.H{
		"countryCode": "US",
		"timezone":    "America/New_York",
	}

	// Grab our router
	router := gin.Default()
	router.GET("/data", getDataByIp)

	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/data?ip=8.8.8.8")

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Grab the value & whether or not it exists
	value, exists := response["countryCode"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["countryCode"], value)

	value, exists = response["timezone"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["timezone"], value)
}

func TestGetDataByIpInvalid(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Build our expected body
	body := gin.H{
		"message": "invalid ip or data not found",
	}

	// Grab our router
	router := gin.Default()
	router.GET("/data", getDataByIp)

	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/data?ip=x.x.x.x")

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Grab the value & whether or not it exists
	value, exists := response["message"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["message"], value)
}
