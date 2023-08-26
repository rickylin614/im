
# ricky-im

## 需要初始化的物件皆需要宣告DI

## code規範

1. http請求GET / POST / PUT / DELETE皆使用JSON
2. handler層可呼叫service層
3. service層可呼叫repository層
4. 需要初始化物件時, 皆需使用dig創建

## nunu

- 根據模板快速創立文件
    - 模板放置template/nunu目錄底下

```sh
# 安裝
go install gitlab.paradise-soft.com.tw/backend/sloth/cmd/nunu@v1.0.2

# 以下為建立模板`user`範例
# handler
nunu create handler user
# service
nunu create service user
# repository
nunu create repository user
# model
nunu create model user
# all (create handler,service,repository,model)
nunu create all user

# 添加`user`至provider範例
nunu append user
```

## po層model定義

1. 使用`https://github.com/xxjwxc/gormt`從DB產出對應struct
2. 從DB查詢出DDL，使用`http://sql2struct.atotoa.com/`等三方DDL to struct工具 (註: po只保留gorm tag)

## vscode setting

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "apis",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "./cmd/apis/main.go",
            "cwd": "./"
        }
    ]
}
```

## swagger install

- `go install github.com/swaggo/swag/cmd/swag@latest`

### swagger local path

- http://localhost:8800/swagger/index.html

### swagger generate

- `make docs`

## code example

### db read or write transaction demo
```
// Start transaction based on default replicas db
tx := DB.Clauses(dbresolver.Read).Begin()

// Start transaction based on default sources db
tx := DB.Clauses(dbresolver.Write).Begin()
```


## Binding 使用

- binding tag的使用範例

在 Gin 框架中，`Binder` 的主要工作就是解析請求並將其映射到定義的數據結構中。

1. `required`：該欄位必須有值，不能為空。
2. `eq`：該欄位的值必須等於指定的值。
3. `ne`：該欄位的值必須不等於指定的值。
4. `lt`：該欄位的值必須小於指定的值。
5. `lte`：該欄位的值必須小於或等於指定的值。
6. `gt`：該欄位的值必須大於指定的值。
7. `gte`：該欄位的值必須大於或等於指定的值。
8. `email`：該欄位的值必須是有效的電子郵件地址。
9. `len`：該欄位的值的長度必須等於指定的值。
10. `min`：該欄位的值的長度必須至少為指定的值。
11. `max`：該欄位的值的長度必須最多為指定的值。
12. `regexp`：該欄位的值必須符合指定的正則表達式。

下面是一個如何在 Go 結構中使用 `validator` 套件的例子：

```go
type MyData struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}
```

在這個例子中，`Email` 必須是必填的且必須是一個有效的電子郵件地址，`Password` 也是必填的且其長度必須至少為 8。這些標籤就是透過 `validator` 套件來進行驗證的。
