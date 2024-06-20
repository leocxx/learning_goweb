如何提供静态文件，如 CSS、JavaScript 或特定目录中的图像。
```go
package main

import "net/http"

func main() {
    fs := http.FileServer(http.Dir("assets/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8080", nil)
}
```
```console
$ tree assets/
assets/
└── css
    └── styles.css
```
