package docs

import (
	"os"

	"github.com/swaggo/swag"
)

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "soberkoder@swagger.io"
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
        "/CreateData": {
            "post": {
                "description": "傳入以下參數 \u003cbr\u003e Id: 透過DB查出 \u003cbr\u003e TenderID: 標案號 \u003cbr\u003e Accountcode: 銀行代碼 \u003cbr\u003e Account: 銀行代碼 \u003cbr\u003e Name: 戶名 \u003cbr\u003e Currency: 幣別 \u003cbr\u003e Branch: 分行 \u003cbr\u003e Amount: 金額 \u003cbr\u003e Status: 憑證狀態",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DApp"
                ],
                "summary": "https://tender-chain.fishlab.com.tw/CreateData",
                "parameters": [
                    {
                        "description": "CreateData",
                        "name": "dapp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Data"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Data already exists"
                    }
                }
            }
        },
        "/GetHistory": {
            "get": {
                "description": "傳入以下參數 \u003cbr\u003e Id: 透過DB查出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DApp"
                ],
                "summary": "https://tender-chain.fishlab.com.tw/GetHistory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id",
                        "name": "Id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/ReadData": {
            "get": {
                "description": "傳入以下參數 \u003cbr\u003e Id: 透過DB查出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DApp"
                ],
                "summary": "https://tender-chain.fishlab.com.tw/ReadData",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id",
                        "name": "Id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Data"
                        }
                    }
                }
            }
        },
        "/UpdateData": {
            "post": {
                "description": "\u003cbr\u003e Id: 透過DB查出 \u003cbr\u003e TenderID: 標案號 \u003cbr\u003e Accountcode: 銀行代碼 \u003cbr\u003e Account: 銀行代碼 \u003cbr\u003e Name: 戶名 \u003cbr\u003e Currency: 幣別 \u003cbr\u003e Branch: 分行 \u003cbr\u003e Amount: 金額 \u003cbr\u003e Status: 憑證狀態",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DApp"
                ],
                "summary": "https://tender-chain.fishlab.com.tw/UpdateData",
                "parameters": [
                    {
                        "description": "UpdateData",
                        "name": "dapp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Data"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Data": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "string"
                },
                "TenderID": {
                    "type": "string"
                },
                "Accountcode": {
                    "type": "string"
                },
                "Account": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Currency": {
                    "type": "string"
                },
                "Amount": {
                    "type": "string"
                },
                "Branch": {
                    "type": "string"
                },
                "Status": {
                    "type": "string"
                }
                
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             os.Getenv("ENDPOINT_URL"),
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "TenderChain Dapp API",
	Description:      "This is a sample serice for managing dapp",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
