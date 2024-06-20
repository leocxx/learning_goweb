 #### 介绍
 `net/http` 的软件包为 HTTP 协议提供了许多功能。它做得不好的一件事是复杂的请求路由，例如将请求 url 分割为单个参数。可以通过gorilla/mux包解决
 
 `gorilla/mux` 是一个适应 Go 默认 HTTP 路由器的包。它具有许多功能，可以提高编写 Web 应用程序时的生产力。它还符合 Go 的默认请求处理程序签名 `func (w http.ResponseWriter, r *http.Request)` ，因此该包可以与其他 HTTP 库（如中间件或现有应用程序）混合和修改。使用命令 `go get` 从 GitHub 安装包
```sh
go get -u github.com/gorilla/mux
```

#### 创建新的路由器
首先创建一个新的请求路由器。路由器是 Web 应用程序的主路由器，稍后将作为参数传递给服务器。它将接收所有 HTTP 连接并将其传递给您将在其上注册的请求处理程序。您可以像这样创建一个新路由器：
```go
r := mux.NewRouter()
```

#### 注册请求处理程序
一旦你有了一个新的路由器，你就可以像往常一样注册请求处理程序。唯一的区别是，您不是在路由器 `http.HandleFunc(...)` 上调用，而是像这样调用 `HandleFunc` ： `r.HandleFunc(...)`

#### URL参数
路由器的最大优势 `gorilla/mux` 是能够从请求 URL 中提取分段

#### 设置HTTP服务器的路由器

有没有想过 `nil` 在 `http.ListenAndServe(":80", nil)` 是什么？它是HTTP服务器主路由器的参数。默认情况下，它是 `nil` ，这意味着使用 `net/http` 软件包的默认路由器。要使用您自己的路由器，请将 替换 `nil` 为 路由器的变量 `r` 。

```go
http.ListenAndServe(":80", r)
```


#### `gorilla/mux` 路由器的特点

##### 方法
将请求处理程序限制为特定的 HTTP 方法。
```go
r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
```

##### 主机名和子域
将请求处理程序限制为特定主机名或子域。
```go
r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")
```


##### Schemes
将请求处理程序限制为 http/https。
```go
r.HandleFunc("/secure", SecureHandler).Schemes("https")
r.HandleFunc("/insecure", InsecureHandler).Schemes("http")
```

##### 路径前缀和子路由器
将请求处理程序限制为特定路径前缀。
```go
bookrouter := r.PathPrefix("/books").Subrouter()
bookrouter.HandleFunc("/", AllBooks)
bookrouter.HandleFunc("/{title}", GetBook)
```
