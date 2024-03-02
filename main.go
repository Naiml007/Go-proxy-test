package main

import (
  "fmt"
  "net/http"
  "net/http/httputil"
  "net/url"
)

func main() {
  // Set the target URL of your Node.js API
  targetURL, err := url.Parse("https://consumet-api-de70.onrender.com")
  if err != nil {
    panic(err)
  }

  // Create a reverse proxy
  proxy := httputil.NewSingleHostReverseProxy(targetURL)

  // Start a server to handle incoming requests
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // Serve as a reverse proxy
    proxy.ServeHTTP(w, r)
  })

  // Specify the port to listen on
  port := 8080
  fmt.Printf("Proxy server is running on port %d...\n", port)
  err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
  if err != nil {
    panic(err)
  }
}
