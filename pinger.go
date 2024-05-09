package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var ServiceURL = "http://localhost:8081"

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ServiceURL := os.Getenv("SERVICE_URL")
	proxyURL, _ := url.Parse(ServiceURL)
	proxy := httputil.NewSingleHostReverseProxy(proxyURL)

	r.URL.Host = proxyURL.Host
	r.URL.Scheme = proxyURL.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = proxyURL.Host

	proxy.ServeHTTP(w, r)
}
