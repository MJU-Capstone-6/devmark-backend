// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "MilkyMilky0116",
            "url": "https://milkymilky0116.github.io",
            "email": "sjlee990129@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/auth/kakao": {
            "post": {
                "description": "retrive user info from kakao oauth. if user exists in our service, then return access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "retrive user info from kakao oauth",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.GetKakaoInfoResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    }
                }
            }
        },
        "/api/v1/category": {
            "post": {
                "description": "카테고리를 생성합니다. *워크스페이스에 카테고리를 등록하는 API는 따로 존재합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Create Category",
                "parameters": [
                    {
                        "description": "body to create category",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/category.CreateCategoryParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repository.Category"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    }
                }
            }
        },
        "/api/v1/category/:id": {
            "put": {
                "description": "카테고리를 업데이트 합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Update Category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body to update category",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/repository.UpdateCategoryParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repository.Category"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    }
                }
            },
            "delete": {
                "description": "카테고리를 삭제 합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Delete Category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repository.Category"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    }
                }
            }
        },
        "/api/v1/invitecode": {
            "post": {
                "description": "워크스페이스의 초대코드를 생성합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "invitecode"
                ],
                "summary": "Generate Invite Code from workspace",
                "parameters": [
                    {
                        "description": "body to Generate Invite code",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/repository.CreateInviteCodeParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repository.InviteCode"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    }
                }
            }
        },
        "/api/v1/workspace": {
            "post": {
                "description": "워크스페이스를 생성합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspace"
                ],
                "summary": "워크스페이스 생성",
                "parameters": [
                    {
                        "description": "Workspace id",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/workspace.CreateWorkspaceParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repository.Workspace"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    }
                }
            }
        },
        "/api/v1/workspace/:id": {
            "get": {
                "description": "워크스페이스를 조회합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspace"
                ],
                "summary": "워크스페이스 조회",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Workspace id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repository.Workspace"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    }
                }
            },
            "put": {
                "description": "워크스페이스 정보를 업데이트 합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspace"
                ],
                "summary": "워크스페이스 업데이트",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Workspace id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Workspace info",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/repository.UpdateWorkspaceParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repository.Workspace"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    }
                }
            },
            "delete": {
                "description": "워크스페이스를 삭제합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspace"
                ],
                "summary": "워크스페이스 삭제",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Workspace id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repository.Workspace"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    }
                }
            }
        },
        "/api/v1/workspace/:id/join": {
            "post": {
                "description": "워크스페이스에 참가합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspace"
                ],
                "summary": "워크스페이스 참가",
                "parameters": [
                    {
                        "description": "Workspace join info",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/workspace.JoinWorkspaceParam"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Workspace id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repository.Workspace"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customerror.CustomError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "category.CreateCategoryParam": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "customerror.CustomError": {
            "type": "object",
            "properties": {
                "data": {},
                "error_code": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "pgtype.InfinityModifier": {
            "type": "integer",
            "enum": [
                1,
                0,
                -1
            ],
            "x-enum-varnames": [
                "Infinity",
                "Finite",
                "NegativeInfinity"
            ]
        },
        "pgtype.Timestamp": {
            "type": "object",
            "properties": {
                "infinityModifier": {
                    "$ref": "#/definitions/pgtype.InfinityModifier"
                },
                "time": {
                    "description": "Time zone will be ignored when encoding to PostgreSQL.",
                    "type": "string"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        },
        "pgtype.Timestamptz": {
            "type": "object",
            "properties": {
                "infinityModifier": {
                    "$ref": "#/definitions/pgtype.InfinityModifier"
                },
                "time": {
                    "type": "string"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        },
        "repository.Category": {
            "type": "object",
            "properties": {
                "created_at": {
                    "$ref": "#/definitions/pgtype.Timestamptz"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "$ref": "#/definitions/pgtype.Timestamptz"
                }
            }
        },
        "repository.CreateInviteCodeParams": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "workspace_id": {
                    "type": "integer"
                }
            }
        },
        "repository.InviteCode": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "created_at": {
                    "$ref": "#/definitions/pgtype.Timestamptz"
                },
                "expired_at": {
                    "$ref": "#/definitions/pgtype.Timestamptz"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "$ref": "#/definitions/pgtype.Timestamptz"
                },
                "workspace_id": {
                    "type": "integer"
                }
            }
        },
        "repository.UpdateCategoryParams": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "repository.UpdateWorkspaceParams": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "repository.Workspace": {
            "type": "object",
            "properties": {
                "created_at": {
                    "$ref": "#/definitions/pgtype.Timestamp"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "$ref": "#/definitions/pgtype.Timestamp"
                }
            }
        },
        "responses.GetKakaoInfoResponse": {
            "type": "object",
            "properties": {
                "access_key": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "workspace.CreateWorkspaceParam": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "workspace.JoinWorkspaceParam": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Devmark API",
	Description:      "Bookmark service with automatic classification.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
