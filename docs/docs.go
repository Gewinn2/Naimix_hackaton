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
        "/company/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Create a new company",
                "parameters": [
                    {
                        "description": "Company creation request data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.CreateCompanyRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entities.CreateCompanyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/company/{id}/members": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Get members of a company by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Company ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of company members",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.CompanyMemberInfo"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/cosmogram": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cosmogram"
                ],
                "summary": "Get cosmogram",
                "parameters": [
                    {
                        "description": "Get cosmogram data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.CosmogramRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.CosmogramApiResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User login data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.LoginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.LoginUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Sign up user",
                "parameters": [
                    {
                        "description": "User registration data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tarot/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tarot"
                ],
                "summary": "Get a tarot card by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tarot card ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Details of the tarot card",
                        "schema": {
                            "$ref": "#/definitions/entities.TaroCard"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "List of users",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/entities.User"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Aspect": {
            "type": "object",
            "properties": {
                "aspect": {
                    "type": "string"
                },
                "orbit": {
                    "type": "number"
                },
                "p1_name": {
                    "type": "string"
                },
                "p2_name": {
                    "type": "string"
                }
            }
        },
        "entities.CompanyMemberInfo": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "integer"
                },
                "surname": {
                    "type": "string"
                },
                "third_name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "entities.CosmogramApiResponseBody": {
            "type": "object",
            "properties": {
                "aspects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Aspect"
                    }
                },
                "data": {
                    "$ref": "#/definitions/entities.Data"
                },
                "is_destiny_sign": {
                    "type": "boolean"
                },
                "score": {
                    "type": "integer"
                },
                "score_description": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "entities.CosmogramRequestBody": {
            "type": "object",
            "properties": {
                "candidate": {
                    "$ref": "#/definitions/entities.User"
                },
                "staff": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.User"
                    }
                }
            }
        },
        "entities.CreateCompanyRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "entities.CreateCompanyResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "entities.CreateUserRequest": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string",
                    "example": "1990-01-01"
                },
                "email": {
                    "type": "string",
                    "example": "john.doe@example.com"
                },
                "name": {
                    "type": "string",
                    "example": "John"
                },
                "password": {
                    "type": "string",
                    "example": "P@ssw0rd123"
                },
                "phone_number": {
                    "type": "string",
                    "example": "+1234567890"
                },
                "role": {
                    "type": "string",
                    "example": "HR manager"
                },
                "surname": {
                    "type": "string",
                    "example": "Doe"
                },
                "third_name": {
                    "type": "string",
                    "example": "Middle"
                }
            }
        },
        "entities.CreateUserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiRXhwaXJlc0F0IjoxNzMxNzE1ODM0LCJJc3N1ZWRBciI6MTczMTY5NDIzNH0sInVzZXJfaWQiOjJ9.4PZY_8iDHCeACmPs5woqtowvCrrByTB4Gadr6oLdlZg"
                },
                "email": {
                    "type": "string",
                    "example": "john.doe@example.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "refresh_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiRXhwaXJlc0F0IjoxNzM0MzAwNjM0LCJJc3N1ZWRBciI6MTczMTY5NDIzNH0sInVzZXJfaWQiOjJ9.o9i-1GBcLtUR0SpDNiBuHfzQWHBt465bhxcC_X2WqBY"
                }
            }
        },
        "entities.Data": {
            "type": "object",
            "properties": {
                "first_subject": {
                    "$ref": "#/definitions/entities.Subject2"
                },
                "second_subject": {
                    "$ref": "#/definitions/entities.Subject2"
                }
            }
        },
        "entities.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Invalid data provided"
                }
            }
        },
        "entities.LoginUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john.doe@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "P@ssw0rd123"
                },
                "phone_number": {
                    "type": "string",
                    "example": "+1234567890"
                }
            }
        },
        "entities.LoginUserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiRXhwaXJlc0F0IjoxNzMxNzE1ODM0LCJJc3N1ZWRBciI6MTczMTY5NDIzNH0sInVzZXJfaWQiOjJ9.4PZY_8iDHCeACmPs5woqtowvCrrByTB4Gadr6oLdlZg"
                },
                "email": {
                    "type": "string",
                    "example": "john.doe@example.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "refresh_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiRXhwaXJlc0F0IjoxNzM0MzAwNjM0LCJJc3N1ZWRBciI6MTczMTY5NDIzNH0sInVzZXJfaWQiOjJ9.o9i-1GBcLtUR0SpDNiBuHfzQWHBt465bhxcC_X2WqBY"
                }
            }
        },
<<<<<<< HEAD
        "entities.TaroCard": {
            "type": "object",
            "properties": {
                "direct_meaning": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "reverse_meaning": {
=======
        "entities.Planet": {
            "type": "object",
            "properties": {
                "abs_pos": {
                    "type": "number"
                },
                "element": {
                    "type": "string"
                },
                "emoji": {
                    "type": "string"
                },
                "house": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "point_type": {
                    "type": "string"
                },
                "position": {
                    "type": "number"
                },
                "quality": {
                    "type": "string"
                },
                "retrograde": {
                    "type": "boolean"
                },
                "sign": {
                    "type": "string"
                },
                "sign_num": {
                    "type": "integer"
                }
            }
        },
        "entities.Subject2": {
            "type": "object",
            "properties": {
                "chiron": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "city": {
                    "type": "string"
                },
                "day": {
                    "type": "integer"
                },
                "hour": {
                    "type": "integer"
                },
                "houses_system_identifier": {
                    "type": "string"
                },
                "houses_system_name": {
                    "type": "string"
                },
                "iso_formatted_local_datetime": {
                    "type": "string"
                },
                "iso_formatted_utc_datetime": {
                    "type": "string"
                },
                "julian_day": {
                    "type": "number"
                },
                "jupiter": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "lat": {
                    "type": "number"
                },
                "lng": {
                    "type": "number"
                },
                "local_time": {
                    "type": "number"
                },
                "mars": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "mean_lilith": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "mean_node": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "mean_south_node": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "mercury": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "minute": {
                    "type": "integer"
                },
                "month": {
                    "type": "integer"
                },
                "moon": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "name": {
                    "type": "string"
                },
                "nation": {
                    "type": "string"
                },
                "neptune": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "perspective_type": {
                    "type": "string"
                },
                "pluto": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "saturn": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "sidereal_mode": {
                    "type": "string"
                },
                "sun": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "tz_str": {
                    "type": "string"
                },
                "uranus": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "utc_time": {
                    "type": "number"
                },
                "venus": {
                    "$ref": "#/definitions/entities.Planet"
                },
                "year": {
                    "type": "integer"
                },
                "zodiac_type": {
>>>>>>> 8c51950ffbe0fe76694f0ab70623b52b48f2ac05
                    "type": "string"
                }
            }
        },
        "entities.User": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string",
                    "example": "1990-01-01"
                },
                "email": {
                    "type": "string",
                    "example": "john.doe@example.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "John"
                },
                "password": {
                    "type": "string",
                    "example": "P@ssw0rd123"
                },
                "phone_number": {
                    "type": "string",
                    "example": "+1234567890"
                },
                "role": {
                    "type": "string",
                    "example": "user"
                },
                "surname": {
                    "type": "string",
                    "example": "Doe"
                },
                "third_name": {
                    "type": "string",
                    "example": "Petrovich"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Harmony Compass API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
