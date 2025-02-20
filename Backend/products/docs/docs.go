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
        "/products": {
            "get": {
                "description": "Fetch all products from the system, regardless of the user ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get all products in the system",
                "responses": {
                    "200": {
                        "description": "List of all products grouped by user",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.UserProduct"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new product by parsing form data, uploading images to S3, and saving to the database.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Create a new product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product title",
                        "name": "productTitle",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product description",
                        "name": "productDescription",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "description": "Product price",
                        "name": "productPrice",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Product condition",
                        "name": "productCondition",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product location",
                        "name": "productLocation",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Product image",
                        "name": "productImage",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Product created successfully",
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid userId or form data",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/products/{UserId}": {
            "get": {
                "description": "Fetch all products listed by a user, identified by their user ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get all products for a specific user by user ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "UserId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of products",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Product"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "No products found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/products/{UserId}/{ProductId}": {
            "put": {
                "description": "Update a product's details based on the user ID and product ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update a product by user ID and product ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "UserId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "ProductId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Product Details",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated product",
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Product not found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a product from the system based on the user ID and product ID.",
                "tags": [
                    "Products"
                ],
                "summary": "Delete a product by user ID and product ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "UserId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "ProductId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Product deleted"
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Product not found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ErrorResponse": {
            "description": "Represents an error response when an operation fails.",
            "type": "object",
            "properties": {
                "details": {
                    "description": "Detailed error message with multiple examples",
                    "type": "string",
                    "example": "ProductPrice: cannot be empty or zero, Product not found"
                },
                "error": {
                    "description": "Error message with example value",
                    "type": "string",
                    "example": "Error updaing product"
                }
            }
        },
        "model.Product": {
            "description": "Represents a product for sale in the marketplace.",
            "type": "object",
            "properties": {
                "productCondition": {
                    "description": "Product condition",
                    "type": "integer",
                    "example": 4
                },
                "productDescription": {
                    "description": "Product description",
                    "type": "string",
                    "example": "A high-performance laptop"
                },
                "productId": {
                    "description": "Unique product ID (UUID)",
                    "type": "string",
                    "example": "9b96a85c-f02e-47a1-9a1a-1dd9ed6147bd"
                },
                "productImage": {
                    "description": "Product image URL in GET, Actual product image in PUT",
                    "type": "string",
                    "example": "https://example.com/laptop.jpg"
                },
                "productLocation": {
                    "description": "Location of the product",
                    "type": "string",
                    "example": "University of Florida"
                },
                "productPostDate": {
                    "description": "Product post date (DD-MM-YYYY)",
                    "type": "string",
                    "example": "20-02-2025"
                },
                "productPrice": {
                    "description": "Price of the product",
                    "type": "number",
                    "example": 999.99
                },
                "productTitle": {
                    "description": "Product title",
                    "type": "string",
                    "example": "Laptop"
                }
            }
        },
        "model.UserProduct": {
            "description": "Represents a user along with their associated products in the marketplace.",
            "type": "object",
            "properties": {
                "products": {
                    "description": "List of products owned by the user",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Product"
                    }
                },
                "userId": {
                    "description": "UserID with example value",
                    "type": "integer",
                    "example": 123
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "UniBazaar Products API",
	Description:      "API for managing products in the UniBazaar marketplace for university students.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
