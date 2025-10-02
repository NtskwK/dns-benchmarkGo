package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"net/http"
)

const distRoot = "./web/dist"

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

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			w.Write(modifiedIndexHTML())
			return
		}
		// Serve static files
		fs := http.FileServer(http.Dir(distRoot))
		fs.ServeHTTP(w, r)
	})

	fmt.Println("Server starting on :8000")
	http.ListenAndServe(":8000", nil)
}
