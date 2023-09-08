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
        "/admin/show": {
            "get": {
                "tags": [
                    "管理员"
                ],
                "summary": "通过审核",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "获取信息id",
                        "name": "messageId",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AdminChooseTask"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AdminChooseTask": {
            "type": "object",
            "properties": {
                "Organization": {
                    "description": "organization",
                    "type": "integer"
                },
                "Posts": {
                    "description": "posts",
                    "type": "integer"
                },
                "State": {
                    "description": "State",
                    "type": "integer"
                },
                "StuNum": {
                    "description": "StuNum",
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
