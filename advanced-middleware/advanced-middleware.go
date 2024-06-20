package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// 记录所有请求及其路径和处理时间
func Logging() Middleware {

	// 创建一个新的中间件

	return func(f http.HandlerFunc) http.HandlerFunc {

		// 定义 http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// 调用下一个中间件
			f(w, r)
		}
	}
}

// 方法确保 URL 只能通过特定方法请求，否则返回 400
func Method(m string) Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(w, r)
		}
	}
}

// 应用中间件
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}
