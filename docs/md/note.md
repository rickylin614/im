
# ricky-im

## 需要初始化的物件皆需要宣告DI

## code規範

1. http請求GET / POST / PUT / DELETE皆使用JSON
2. handler層可呼叫service層
3. service層可呼叫repository層
4. 需要初始化物件時, 皆需使用dig創建

## nunu

- 根據模板快速創立文件
    - 模板放置 template / nunu 目錄底下

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