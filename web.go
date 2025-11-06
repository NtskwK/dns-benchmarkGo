package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed web/dist/index.html
var indexHTML []byte

//go:embed web/dist/static/js/index.js
var indexJS []byte

//go:embed web/dist/static/css/index.css
var indexCSS []byte

func ServeOn(port int, resultJson string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			w.Write(indexHTML)
			return
		} else if r.URL.Path == "/static/js/index.js" {
			w.Header().Set("Content-Type", "application/javascript")
			w.Write(indexJS)
			return
		} else if r.URL.Path == "/static/css/index.css" {
			w.Header().Set("Content-Type", "text/css")
			w.Write(indexCSS)
			return
		} else if r.URL.Path == "/api/data" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(resultJson))
			return
		}
	})

	fmt.Printf("Server starting on :%d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
