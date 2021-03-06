// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Firat Atmaca",
            "url": "frtatmaca",
            "email": "frtatmaca@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/deeplinktourl": {
            "post": {
                "description": "Deep Link To Url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deeplinktourl"
                ],
                "summary": "Deep Link To Url",
                "parameters": [
                    {
                        "description": "Link Converter Input",
                        "name": "DeepLinkToUrl",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UriResponse"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UriResponse"
                        }
                    }
                }
            }
        },
        "/urltodeeplink": {
            "post": {
                "description": "Deep Link To Url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urltodeeplink"
                ],
                "summary": "Url To Deep Link",
                "parameters": [
                    {
                        "description": "Link Converter Input",
                        "name": "UrlToDeepLink",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UriResponse"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UriResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.UriResponse": {
            "type": "object",
            "properties": {
                "deepLink": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Go Restful API with Swagger",
	Description: "Simple swagger implementation in Go HTTP",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
