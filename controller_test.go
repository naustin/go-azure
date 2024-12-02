package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context){
}

func TestValidMetric(t *testing.T) {

    // Switch to test mode so you don't get such noisy output
    gin.SetMode(gin.TestMode)

    // Setup your router, just like you did in your main function, and
    // register your routes
    r := gin.Default()
    r.POST("/process-alert", RouteBySignalType)

    // Create the mock request you'd like to test. Make sure the second argument
    // here is the same as one of the routes you defined in the router setup
    // block!
	file, err := os.Open("./testdata/metric1_gov.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

    req, err := http.NewRequest(http.MethodPost, "/process-alert", file)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    // Create a response recorder so you can inspect the response
    w := httptest.NewRecorder()

    // Perform the request
    r.ServeHTTP(w, req)

    // Check to see if the response was what you expected
    if w.Code != http.StatusOK {
        t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
    }

	body, _ := io.ReadAll(w.Body)
	
	t.Log(string(body))
}