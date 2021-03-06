package main

import (
	"log"
	"net/http"
	"strings"
)

const dir = "example/html"

func main() {
	fs := http.FileServer(http.Dir(dir))
	log.Print("Serving " + dir + " on http://localhost:8080")
	http.ListenAndServe(":8080", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}

		if strings.HasSuffix(req.URL.Path, ".css") {
			resp.Header().Set("content-type", "text/css")
		}

		if strings.HasSuffix(req.URL.Path, ".jpg") {
			resp.Header().Set("content-type", "image/jpeg")
		}

		if strings.HasSuffix(req.URL.Path, ".mp4") {
			resp.Header().Set("content-type", "video/mp4")
		}

		fs.ServeHTTP(resp, req)
	}))
}
