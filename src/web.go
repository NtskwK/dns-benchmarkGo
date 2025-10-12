package main

import (
	"bytes"
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

func modifiedIndexHTML() []byte {
	// 构建内联CSS和JS
	cssInline := append(append([]byte(`<style>`), indexCSS...), []byte(`</style>`)...)
	jsInline := append(append([]byte(`<script>`), indexJS...), []byte(`</script>`)...)

	// 组合添加到body末尾
	addition := append(cssInline, jsInline...)

	// 替换</body>为 addition + </body>
	modifiedIndexHTML := bytes.ReplaceAll(indexHTML, []byte(`</body>`), append(addition, []byte(`</body>`)...))

	return modifiedIndexHTML
}

func findSafePort(start, end int) (int, error) {
	for port := start; port <= end; port++ {
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err == nil {
			return port, nil
		}
	}
	return 0, fmt.Errorf("no available port found")
}

func serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			w.Write(modifiedIndexHTML())
			return
		}
	})

	port, err := findSafePort(8000, 8080)
	if err != nil {
		fmt.Println("Error finding available port:", err)
		return err
	}

	fmt.Println("Server starting on :%d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
