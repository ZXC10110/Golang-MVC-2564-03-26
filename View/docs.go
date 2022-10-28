package View

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.3",
  "info": {
    "title": "Golang MVC Structure",
    "version": "1.0.11"
  },
  "tags": [
    {
      "name": "picture",
      "description": "Testing of MVC's structure"
    }
  ],
  "paths": {
    "/uploadPicture": {
      "post": {
        "tags": [
          "picture"
        ],
        "summary": "create picture",
        "requestBody": {
          "description": "create picture",
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/Picture"
              }
            },
            "application/xml": {
              "schema": {
                "$ref": "#/components/schemas/Picture"
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "Upload Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Picture"
                }
              }
            }
          }
        }
      }
    },
    "/viralPicture": {
      "get": {
        "tags": [
          "picture"
        ],
        "summary": "list all picture",
        "responses": {
          "201": {
            "description": "list Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Picture"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "Picture": {
        "type": "object",
        "properties": {
          "pic_no": {
            "type": "integer",
            "format": "int64"
          },
          "reach": {
            "type": "integer",
            "format": "int64"
          },
          "like_qty": {
            "type": "integer",
            "format": "int64"
          },
          "share_qty": {
            "type": "integer",
            "format": "int64"
          },
          "comment_qty": {
            "type": "integer",
            "format": "int64"
          }
        },
        "xml": {
          "name": "Picture"
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
	Title:            "Golang MVC",
	Description:      "This is a sample server todo server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
