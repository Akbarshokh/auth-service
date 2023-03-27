// Code generated by swaggo/swag. DO NOT EDIT
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
        "/check-token": {
            "post": {
                "description": "This endpoint verifies token is active or not",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Auth Service"
                ],
                "summary": "Checking token with Access Token",
                "parameters": [
                    {
                        "description": "Access Token",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CheckTokenReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.CheckTokenRes"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        },
        "/get-token": {
            "post": {
                "description": "This endpoint verifies token is active or not and generates new access token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Auth Service"
                ],
                "summary": "Checking token with Refresh Token",
                "parameters": [
                    {
                        "description": "Access Token",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GetTokenReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.GetTokenRes"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        },
        "/sign-in": {
            "post": {
                "description": "API for Sign In",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Auth Service"
                ],
                "summary": "Sign In using client_id, email, and access_token",
                "parameters": [
                    {
                        "description": "Client ID",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignInReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        },
        "/sign-up": {
            "post": {
                "description": "API for Sign Up",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Auth Service"
                ],
                "summary": "Sign Up",
                "parameters": [
                    {
                        "description": "Client ID",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SignUpRes"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CheckTokenReq": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "models.CheckTokenRes": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                }
            }
        },
        "models.GetTokenReq": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.GetTokenRes": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "active": {
                    "type": "boolean"
                },
                "client_id": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.SignInReq": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "client_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "models.SignUpReq": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "device_num": {
                    "type": "string"
                },
                "device_type": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                }
            }
        },
        "models.SignUpRes": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "active": {
                    "type": "boolean"
                },
                "client_id": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "rest.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errorCode": {
                    "type": "integer"
                },
                "errorNote": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
