package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {

	service1Url, _ := url.Parse("http://localhost:8081/")
	service1Proxy := httputil.NewSingleHostReverseProxy(service1Url)
	http.HandleFunc("/service/", func(w http.ResponseWriter, r *http.Request) {

		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/service/")
		service1Proxy.ServeHTTP(w, r)
	})
	http.ListenAndServe(":8080", nil)
}
