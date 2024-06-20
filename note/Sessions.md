如何使用 Go 中流行的 gorilla/sessions 包将数据存储在会话 cookie 中。
Cookies是存储在用户浏览器中的一小段数据，并在每次请求时发送到我们的服务器。例如，我们可以存储用户是否登录我们的网站，并确定他到底是谁
此示例中，我们将仅允许经过身份验证的用户在 `/secret` 页面上查看我们的机密消息。要访问它，首先必须访问 `/login` 以获取有效的会话cookie，该cookie使他登录。此外，他可以访问 `/logout` 以撤销他对我们秘密信息的访问权限。

```go
// sessions.go
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/sessions"
)

var (
    // key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
    key = []byte("super-secret-key")
    store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name")

    // Check if user is authenticated
    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    // Print secret message
    fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name")

    // Authentication goes here
    // ...

    // Set user as authenticated
    session.Values["authenticated"] = true
    session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name")

    // Revoke users authentication
    session.Values["authenticated"] = false
    session.Save(r, w)
}

func main() {
    http.HandleFunc("/secret", secret)
    http.HandleFunc("/login", login)
    http.HandleFunc("/logout", logout)

    http.ListenAndServe(":8080", nil)
}
```