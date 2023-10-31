// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/users": {
            "get": {
                "tags": [
                    "user"
                ],
                "summary": "Search Users",
                "operationId": "search-users",
                "responses": {
                    "200": {
                        "description": "userList",
                        "schema": {
                            "$ref": "#/definitions/v1.dataResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create user",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "request form",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.createUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "id",
                        "schema": {
                            "$ref": "#/definitions/v1.idResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "tags": [
                    "user"
                ],
                "summary": "Get user",
                "operationId": "get-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of user",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "user"
                ],
                "summary": "Delete user",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of user",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "NoContent",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user",
                "operationId": "update-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of user",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update form",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.updateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "NoContent",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "display_name": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "v1.createUserRequest": {
            "type": "object",
            "properties": {
                "display_name": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "v1.dataResponse": {
            "type": "object",
            "properties": {
                "data": {}
            }
        },
        "v1.errResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "v1.idResponse": {
            "type": "object",
            "properties": {
                "user_id": {}
            }
        },
        "v1.updateUserRequest": {
            "type": "object",
            "properties": {
                "display_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:3333",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "TrueConf Backend Test",
	Description:      "Тех. Задание TrueConf",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
