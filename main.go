package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set the target URL of your Node.js API
	targetURL, err := url.Parse("https://consumet-api-de70.onrender.com")
	if err != nil {
		panic(err)
	}

	// Create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Create a new Gin router
	router := gin.New()

	// Use Gin for handling requests
	router.Use(func(c *gin.Context) {
		// Serve as a reverse proxy
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	// Specify the port to listen on
	port := 8080
	fmt.Printf("Proxy server is running on port %d...\n", port)

	// Start Gin server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		panic(err)
	}
}
