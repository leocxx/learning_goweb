  Go 是一种电池内置的编程语言，并且已经内置了 Web 服务器。标准库中的 `net/http` 包包含有关 HTTP 协议的所有功能。这包括（除其他外）HTTP 客户端和 HTTP 服务器。
## 注册请求处理程序
首先，创建一个处理程序，用于接收来自浏览器、HTTP 客户端或 API 请求的所有传入 HTTP 连接。Go 中的处理程序是具有以下签名的函数：
```go
func (w http.ResponseWriter, r *http.Request)
```
该函数接收两个参数：
`http.ResponseWriter`： 这是您编写 text/html 响应的地方
`http.Request` ：有关此 HTTP 请求的所有信息，包括 URL 或标头字段等内容
将请求处理程序注册到默认 HTTP 服务器：
```go
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
})
```

## 监听HTTP 连接
单独的请求处理程序不能接受来自外部的任何 HTTP 连接。HTTP 服务器必须侦听端口才能将连接传递到请求处理程序。由于端口 80 在大多数情况下是 HTTP 流量的默认端口，因此此服务器也将监听它。

以下代码将启动 Go 的默认 HTTP 服务器，并监听端口 80 上的连接。您可以浏览浏览器并 `http://localhost/` 查看服务器处理您的请求。

```go
http.ListenAndServe(":80", nil)
```

整体代码如下：
```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    http.ListenAndServe(":80", nil)
}
```
