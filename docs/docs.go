// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/example": {
            "get": {
                "tags": [
                    "example"
                ],
                "summary": "GetList",
                "parameters": [
                    {
                        "description": "param",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.ExampleGetList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.APIResponse-resp_ExampleGetList"
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "example"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "description": "param",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.ExampleUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.APIResponse-string"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "example"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "param",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.ExampleCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.APIResponse-string"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "example"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "description": "param",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.ExampleDelete"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.APIResponse-string"
                        }
                    }
                }
            }
        },
        "/example/:id": {
            "get": {
                "tags": [
                    "example"
                ],
                "summary": "Get",
                "parameters": [
                    {
                        "description": "param",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.ExampleGet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.APIResponse-resp_ExampleGet"
                        }
                    }
                }
            }
        },
        "/metrics": {
            "get": {
                "tags": [
                    "public"
                ],
                "summary": "Metrics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "tags": [
                    "public"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "tags": [
                    "users"
                ],
                "summary": "用戶清單",
                "parameters": [
                    {
                        "description": "param",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UsersGetList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.APIResponse-resp_UsersGetList"
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "users"
                ],
                "summary": "用戶訊息修改",
                "parameters": [
                    {
                        "description": "param",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UsersUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.APIResponse-string"
                        }
                    }
                }
            }
        },
        "/users/:id": {
            "get": {
                "tags": [
                    "users"
                ],
                "summary": "取得用戶訊息",
                "parameters": [
                    {
                        "description": "param",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UsersGet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.APIResponse-resp_UsersGet"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "summary": "用戶登錄並返回授權令牌",
                "responses": {}
            }
        },
        "/users/logout": {
            "post": {
                "summary": "用戶登出",
                "responses": {}
            }
        },
        "/users/register": {
            "post": {
                "tags": [
                    "users"
                ],
                "summary": "用戶註冊",
                "parameters": [
                    {
                        "description": "param",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UsersCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.APIResponse-string"
                        }
                    }
                }
            }
        },
        "/users/{id}/online-status": {
            "post": {
                "summary": "更新指定用戶ID的在線狀態",
                "responses": {}
            }
        }
    },
    "definitions": {
        "req.ExampleCreate": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "創建範例描述",
                    "type": "string",
                    "example": "一個武林高手"
                },
                "name": {
                    "description": "創建範例名",
                    "type": "string",
                    "example": "小明"
                }
            }
        },
        "req.ExampleDelete": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "刪除範例ID",
                    "type": "string",
                    "example": "1"
                }
            }
        },
        "req.ExampleGet": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "範例描述",
                    "type": "string",
                    "example": "取得描述"
                },
                "name": {
                    "description": "範例名",
                    "type": "string",
                    "example": "名字"
                }
            }
        },
        "req.ExampleGetList": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "description": "範例ID列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "1",
                        "2",
                        "3",
                        "4",
                        "5",
                        "6"
                    ]
                },
                "index": {
                    "description": "頁碼",
                    "type": "integer"
                },
                "name": {
                    "description": "範例名",
                    "type": "string",
                    "example": "名字"
                },
                "order": {
                    "description": "排序",
                    "type": "string",
                    "example": "id asc"
                },
                "size": {
                    "description": "筆數",
                    "type": "integer"
                }
            }
        },
        "req.ExampleUpdate": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "description": {
                    "description": "修改範例描述",
                    "type": "string",
                    "example": "繼承小明武功但沒天分的孩子"
                },
                "id": {
                    "description": "修改範例ID",
                    "type": "string",
                    "example": "1"
                },
                "name": {
                    "description": "修改範例名",
                    "type": "string",
                    "example": "小明的孩子"
                }
            }
        },
        "req.UsersCreate": {
            "type": "object",
            "required": [
                "email",
                "password",
                "phone_number",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "phone_number": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "req.UsersGet": {
            "type": "object"
        },
        "req.UsersGetList": {
            "type": "object",
            "properties": {
                "index": {
                    "description": "頁碼",
                    "type": "integer"
                },
                "order": {
                    "description": "排序",
                    "type": "string",
                    "example": "id asc"
                },
                "size": {
                    "description": "筆數",
                    "type": "integer"
                }
            }
        },
        "req.UsersUpdate": {
            "type": "object"
        },
        "resp.APIResponse-resp_ExampleGet": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {
                    "$ref": "#/definitions/resp.ExampleGet"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "resp.APIResponse-resp_ExampleGetList": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {
                    "$ref": "#/definitions/resp.ExampleGetList"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "resp.APIResponse-resp_UsersGet": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {
                    "$ref": "#/definitions/resp.UsersGet"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "resp.APIResponse-resp_UsersGetList": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {
                    "$ref": "#/definitions/resp.UsersGetList"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "resp.APIResponse-string": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "resp.ExampleGet": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "id": {
                    "description": "數據ID",
                    "type": "string"
                },
                "name": {
                    "description": "範例名",
                    "type": "string"
                }
            }
        },
        "resp.ExampleGetList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/resp.ExampleGet"
                    }
                },
                "page": {
                    "$ref": "#/definitions/resp.PageResponse"
                }
            }
        },
        "resp.PageResponse": {
            "type": "object",
            "properties": {
                "index": {
                    "description": "頁碼",
                    "type": "integer"
                },
                "pages": {
                    "description": "總頁數",
                    "type": "integer"
                },
                "size": {
                    "description": "筆數",
                    "type": "integer"
                },
                "total": {
                    "description": "總筆數",
                    "type": "integer"
                }
            }
        },
        "resp.UsersGet": {
            "type": "object"
        },
        "resp.UsersGetList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/resp.UsersGet"
                    }
                },
                "page": {
                    "$ref": "#/definitions/resp.PageResponse"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:9000",
	BasePath:         "/im",
	Schemes:          []string{},
	Title:            "Im",
	Description:      "This is a project im.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
