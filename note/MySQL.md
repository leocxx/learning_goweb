Go 编程语言附带了一个名为“database/sql”的方便包，用于查询各种 SQL 数据库。这很有用，因为它将所有常见的 SQL 功能抽象到一个 API 中供您使用。Go 不包括的是数据库驱动程序。在 Go 中，数据库驱动程序是一个包，它实现了特定数据库（在我们的例子中为 MySQL）的低级细节。正如您可能已经猜到的那样，这对于保持向前兼容很有用。因为在创建所有 Go 包时，作者无法预见将来会使用每个数据库，并且支持每个可能的数据库将需要大量的维护工作

安装：
```sh
go get -u github.com/go-sql-driver/mysql
```

#### 连接MySQL
请导入 `database/sql` 和 包 `go-sql-driver/mysql` 并打开连接，如下所示：
```go
import "database/sql"
import _ "go-sql-driver/mysql"

db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")

```
#### 创建第一个数据表
```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at DATETIME,
    PRIMARY KEY (id)
);
```
现在我们有了 SQL 命令，我们可以使用该 `database/sql` 包在 MySQL 数据库中创建表：
```go
query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
_, err := db.Exec(query)
```
#### 插入第一个用户
```sql
INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)
```
在 Go 中使用此 SQL 查询，并在表中插入新行：
```go
import "time"

username := "johndoe"
password := "secret"
createdAt := time.Now()

result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
```
要为您的用户获取新创建的 ID，只需像这样获取它：
```go
userID, err := result.LastInsertId()
```

#### 查询用户表
```sql
SELECT id, username, password, created_at FROM users WHERE id = ?
```
在 Go 中，我们首先声明一些变量来存储我们的数据，然后查询单个数据库行，如下所示：
```go
var (
    id        int
    username  string
    password  string
    createdAt time.Time
)


query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
```

#### 查询所有用户

```sql
SELECT id, username, password, created_at FROM users
```
 Go 中，我们首先声明一些变量来存储我们的数据，然后查询单个数据库行，如下所示：
 ```go
type user struct {
    id        int
    username  string
    password  string
    createdAt time.Time
}

rows, err := db.Query(`SELECT id, username, password, created_at FROM users`) // check err
defer rows.Close()

var users []user
for rows.Next() {
    var u user
    err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt) // check err
    users = append(users, u)
}
err := rows.Err() // check err
```

#### 删除
```go
_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1) // check err
```



 