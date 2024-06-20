 `html/template` 的软件包为 HTML 模板提供了丰富的模板语言。它主要用于 Web 应用程序，以在客户端的浏览器中以结构化方式显示数据。Go 模板语言的一大好处是数据的自动转义。无需担心 XSS 攻击，因为 Go 会解析 HTML 模板并在将其显示到浏览器之前转义所有输入。
#### 第一个模板
 Go 编写模板非常简单。此示例显示了一个 TODO 列表，该列表在 HTML 中编写为无序列表 （ul）。渲染模板时，传入的数据可以是 Go 的任何类型的数据结构。它可以是一个简单的字符串或一个数字，甚至可以是嵌套的数据结构，如下例所示。要访问模板中的数据，最靠前的变量是 `{{.}}` access by 。大括号内的点称为管道和数据的根元素。
 ```go
data := TodoPageData{
    PageTitle: "My TODO list",
    Todos: []Todo{
        {Title: "Task 1", Done: false},
        {Title: "Task 2", Done: true},
        {Title: "Task 3", Done: true},
    },
}
```

```html
<h1>{{.PageTitle}}</h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>
```

#### 控制结构
模板语言包含一组丰富的控件结构，用于呈现 HTML。在这里，您将获得最常用的概述。


#### 从文件解析模板
模板可以从磁盘上的字符串或文件解析。通常情况下，模板是磁盘的参数，此示例演示如何执行此操作。在此示例中，与 Go 程序位于同一目录中的模板文件名为 `layout.html` 。

```go
tmpl, err := template.ParseFiles("layout.html")
// or
tmpl := template.Must(template.ParseFiles("layout.html"))
```

##### 在请求处理程序中执行模板
从磁盘解析模板后，就可以在请求处理程序中使用它了。该 `Execute` 函数接受用于 `io.Writer` 写出模板和 `interface{}` to 将数据传递到模板中。当在 `http.ResponseWriter` 上调用函数时，Content-Type is 标头会在 HTTP 响应中自动设置为 `Content-Type: text/html; charset=utf-8` 。
```go
func(w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, "data goes here")
}
```


