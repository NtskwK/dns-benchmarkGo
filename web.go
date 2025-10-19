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

func injectLocalStorageScript(resultJson string) []byte {
	script := fmt.Sprintf(`
		localStorage.setItem("dnsAnalyzerData", %q); 
	`, resultJson)
	return []byte(script)
}

func injectJavaScript(js string) []byte {
	script := fmt.Sprintf(`
		<script>
			%s
		</script>
	`, js)
	return []byte(script)
}

func modifiedIndexHTML(resultJson string) []byte {
	// 构建内联CSS和JS
	cssInline := append(append([]byte(`<style>`), indexCSS...), []byte(`</style>`)...)
	jsContent := append(injectLocalStorageScript(resultJson), indexJS...)
	jsInline := append(append([]byte(`<script>`), jsContent...), []byte(`</script>`)...)
	// injectInline := injectLocalStorageScript(resultJson)

	// 组合添加到body末尾
	addition := append(jsInline, cssInline...)
	// addition := append(cssInline, append(injectInline, jsInline...)...)

	// 替换</body>为 addition + </body>
	modifiedIndexHTML := bytes.ReplaceAll(indexHTML, []byte(`</body>`), append(addition, []byte(`</body>`)...))

	return modifiedIndexHTML
}

func ServeOn(port int, resultJson string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			w.Write(modifiedIndexHTML(resultJson))
			return
		}
	})

	fmt.Printf("Server starting on :%d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
