// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/v1/area/{id}": {
            "get": {
                "description": "根据地区ID 查询 地区数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Area"
                ],
                "summary": "返回 地区数据",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "area id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.AreaService"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "serializer.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "service.AreaService": {
            "type": "object",
            "properties": {
                "area_id": {
                    "type": "integer"
                },
                "area_name": {
                    "type": "string"
                },
                "city_code": {
                    "type": "string"
                },
                "city_level": {
                    "type": "integer"
                },
                "parent_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/main.go",
	Schemes:          []string{},
	Title:            "go_ctry API",
	Description:      "This is a go_ctry Server pets",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}