基本的 HTTP 服务器需要处理一些关键工作:
1. 处理动态请求：处理来自浏览网站、登录其帐户或发布图像的用户的传入请求。
2. 提供静态资源：向浏览器提供 JavaScript、CSS 和图像，为用户创建动态体验。
3. 接受连接：HTTP 服务器必须侦听特定端口才能接受来自 Internet 的连接。

#### 处理动态请求

该 `net/http` 软件包包含接受请求和动态处理请求所需的所有实用程序。我们可以向函数 `http.HandleFunc` 注册一个新的处理程序。它的第一个参数采用要匹配的路径，并采用要执行的函数作为第二个参数。在此示例中：当有人浏览您的网站 （ `http://example.com/` ），他或她将收到一条好消息。
```go
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to my website!")
})
```
对于动态方面，包含 `http.Request` 有关请求及其参数的所有信息。您可以使用 来 `r.URL.Query().Get("token")` 读取 GET 参数，也可以使用 `r.FormValue("email")` 来读取 POST 参数（HTML 表单中的字段）。
#### 提供静态资源

为了提供 JavaScript、CSS 和图像等静态资产，我们使用内置 `http.FileServer` 的 URL 路径。为了使文件服务器正常工作，它需要知道从哪里提供文件。我们可以这样做：
```go
fs := http.FileServer(http.Dir("static/"))
```
一旦我们的文件服务器到位，我们只需要将 url 路径指向它，就像我们对动态请求所做的那样。需要注意的一点是：为了正确提供文件，我们需要剥离一部分 url 路径。通常，这是我们的文件所在的目录的名称。
```go
http.Handle("/static/", http.StripPrefix("/static/", fs))
```
#### 接受连接

完成我们基本HTTP服务器的最后一件事是，在端口上侦听以接受来自Internet的连接。正如你所猜到的，Go 也有一个内置的 HTTP 服务器，我们可以快速启动失败。启动后，您可以在浏览器中查看您的 HTTP 服务器。
 ```go
http.ListenAndServe(":80", nil)
```
