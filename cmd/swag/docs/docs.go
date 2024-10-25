// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/DangPham112000/go-ecommerce-backend-api.git",
        "contact": {
            "name": "Dang Pham",
            "url": "https://google.com",
            "email": "dangpham112000@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/register": {
            "post": {
                "description": "When user is registered send OTP to email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account management"
                ],
                "summary": "User Registration",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/update_password_register": {
            "post": {
                "description": "Update Password Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account management"
                ],
                "summary": "Update Password Register",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdatePasswordRegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/verify_account": {
            "post": {
                "description": "Verify User Login OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account management"
                ],
                "summary": "Verify User Login OTP",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.VerifyInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.RegisterInput": {
            "type": "object",
            "properties": {
                "verify_key": {
                    "type": "string"
                },
                "verify_purpose": {
                    "type": "string"
                },
                "verify_type": {
                    "type": "integer"
                }
            }
        },
        "model.UpdatePasswordRegisterInput": {
            "type": "object",
            "properties": {
                "user_password": {
                    "type": "string"
                },
                "user_token": {
                    "type": "string"
                }
            }
        },
        "model.VerifyInput": {
            "type": "object",
            "properties": {
                "verify_code": {
                    "type": "string"
                },
                "verify_key": {
                    "type": "string"
                }
            }
        },
        "response.ErrorResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "status code",
                    "type": "integer"
                },
                "data": {
                    "description": "returned data"
                },
                "error": {
                    "description": "error message",
                    "type": "string"
                }
            }
        },
        "response.ResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "status code",
                    "type": "integer"
                },
                "data": {
                    "description": "returned data"
                },
                "message": {
                    "description": "error message",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8002",
	BasePath:         "/v1/2024",
	Schemes:          []string{},
	Title:            "Go Ecommerce backend API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}