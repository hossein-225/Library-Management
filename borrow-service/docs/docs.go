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
        "/borrow": {
            "post": {
                "description": "Allows a user to borrow a book from the library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "borrow"
                ],
                "summary": "Borrow a book",
                "parameters": [
                    {
                        "description": "User ID",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Book ID",
                        "name": "book_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Book borrowed successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "User ID or Book ID cannot be empty",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Book or User not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/borrows/user": {
            "get": {
                "description": "Retrieves the list of books borrowed by a specific user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "borrow"
                ],
                "summary": "Get user's borrowed books",
                "parameters": [
                    {
                        "description": "User ID",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of borrowed books retrieved successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "User ID cannot be empty",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "No borrow records found for this user",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/return": {
            "post": {
                "description": "Allows a user to return a borrowed book to the library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "borrow"
                ],
                "summary": "Return a borrowed book",
                "parameters": [
                    {
                        "description": "User ID",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Book ID",
                        "name": "book_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Book returned successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "User ID or Book ID cannot be empty",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Book or User not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.6",
	Host:             "borrow-service:50053",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Library Management API - borrow-service",
	Description:      "API documentation for the Library Management system - borrow-service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
