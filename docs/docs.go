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
        }
    },
    "definitions": {
        "models.Example": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "integer"
                }
            }
        },
        "models.Page": {
            "type": "object",
            "properties": {
                "order": {
                    "description": "排序",
                    "type": "string"
                },
                "pageIndex": {
                    "description": "頁碼",
                    "type": "integer"
                },
                "size": {
                    "description": "筆數",
                    "type": "integer"
                },
                "total": {
                    "description": "總筆數",
                    "type": "integer"
                },
                "totalPage": {
                    "description": "總頁數",
                    "type": "integer"
                }
            }
        },
        "req.ExampleCreate": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "req.ExampleDelete": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "req.ExampleGet": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "req.ExampleGetList": {
            "type": "object",
            "properties": {
                "order": {
                    "description": "排序",
                    "type": "string"
                },
                "pageIndex": {
                    "description": "頁碼",
                    "type": "integer"
                },
                "size": {
                    "description": "筆數",
                    "type": "integer"
                },
                "total": {
                    "description": "總筆數",
                    "type": "integer"
                },
                "totalPage": {
                    "description": "總頁數",
                    "type": "integer"
                }
            }
        },
        "req.ExampleUpdate": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
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
                    "type": "string"
                },
                "name": {
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
                        "$ref": "#/definitions/models.Example"
                    }
                },
                "page": {
                    "$ref": "#/definitions/models.Page"
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
