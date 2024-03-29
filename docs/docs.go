// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Adam Drew",
            "email": "addrew@redhat.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/starter-app-api/v1/db-info": {
            "get": {
                "description": "Query the API to get some database informaiton",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "Get database info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/starter-app-api/v1/hello": {
            "get": {
                "description": "Recieve a greeting from the API",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "Recieve a greeting",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "description": "Determines if application is still alive",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "probes"
                ],
                "summary": "Determines if application is still alive",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/readyz": {
            "get": {
                "description": "Determines readiness of the application",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "probes"
                ],
                "summary": "Determines readiness of the application",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/starter-app-api/v1",
	Schemes:          []string{},
	Title:            "ConsoleDot Go Starter App API",
	Description:      "API Docs for ConsoleDot Go Starter App API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
