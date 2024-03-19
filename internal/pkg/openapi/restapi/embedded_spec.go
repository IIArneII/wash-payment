// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Payment system service for self-service car washes",
    "title": "wash-payment",
    "version": "1.1.0"
  },
  "paths": {
    "/healthCheck": {
      "get": {
        "description": "Checking the server health status.",
        "tags": [
          "Standard"
        ],
        "summary": "Health check",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "ok": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    },
    "/organizations": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "description": "Get a list of organizations",
        "tags": [
          "Organizations"
        ],
        "summary": "Get organizations",
        "operationId": "list",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "default": 1,
            "name": "page",
            "in": "query"
          },
          {
            "maximum": 100,
            "minimum": 1,
            "type": "integer",
            "default": 10,
            "name": "pageSize",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/OrganizationPage"
            }
          },
          "403": {
            "$ref": "#/responses/Forbidden"
          },
          "default": {
            "$ref": "#/responses/InternalError"
          }
        }
      }
    },
    "/organizations/{id}": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "description": "Get information about the specified organization",
        "tags": [
          "Organizations"
        ],
        "summary": "Get organization",
        "operationId": "get",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Organization"
            }
          },
          "403": {
            "$ref": "#/responses/Forbidden"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          },
          "default": {
            "$ref": "#/responses/InternalError"
          }
        }
      }
    },
    "/organizations/{id}/deposit": {
      "post": {
        "security": [
          {
            "authKey": []
          }
        ],
        "description": "Increase the balance of the specified organization by the specified number of kopecks",
        "tags": [
          "Organizations"
        ],
        "summary": "Top up balance",
        "operationId": "deposit",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Deposit"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "403": {
            "$ref": "#/responses/Forbidden"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          },
          "default": {
            "$ref": "#/responses/InternalError"
          }
        }
      }
    },
    "/organizations/{id}/service-prices": {
      "put": {
        "security": [
          {
            "authKey": []
          }
        ],
        "description": "Set prices for services for the specified organization",
        "tags": [
          "Organizations"
        ],
        "summary": "Set service prices",
        "operationId": "setServicePrices",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/ServicePrices"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "403": {
            "$ref": "#/responses/Forbidden"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          },
          "default": {
            "$ref": "#/responses/InternalError"
          }
        }
      }
    },
    "/organizations/{id}/transactions": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "description": "Get a list of transactions for the specified organization",
        "tags": [
          "Organizations"
        ],
        "summary": "Get organization transactions",
        "operationId": "transactions",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "default": 1,
            "name": "page",
            "in": "query"
          },
          {
            "maximum": 100,
            "minimum": 1,
            "type": "integer",
            "default": 10,
            "name": "pageSize",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/TransactionPage"
            }
          },
          "403": {
            "$ref": "#/responses/Forbidden"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          },
          "default": {
            "$ref": "#/responses/InternalError"
          }
        }
      }
    }
  },
  "definitions": {
    "Deposit": {
      "type": "object",
      "required": [
        "amount"
      ],
      "properties": {
        "amount": {
          "description": "Amount in kopecks (RUB * 10^2)",
          "type": "integer",
          "format": "int64",
          "minimum": 1
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Operation": {
      "description": "Type of operation",
      "type": "string",
      "enum": [
        "deposit",
        "debit"
      ]
    },
    "Organization": {
      "type": "object",
      "required": [
        "id",
        "name",
        "displayName",
        "description",
        "balance",
        "servicePrices"
      ],
      "properties": {
        "balance": {
          "description": "Balance in kopecks (RUB * 10^2)",
          "type": "integer",
          "format": "int64"
        },
        "description": {
          "type": "string"
        },
        "displayName": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string"
        },
        "servicePrices": {
          "$ref": "#/definitions/ServicePrices"
        }
      }
    },
    "OrganizationPage": {
      "type": "object",
      "required": [
        "items",
        "page",
        "pageSize",
        "totalPages",
        "totalItems"
      ],
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Organization"
          }
        },
        "page": {
          "type": "integer"
        },
        "pageSize": {
          "type": "integer"
        },
        "totalItems": {
          "type": "integer"
        },
        "totalPages": {
          "type": "integer"
        }
      }
    },
    "Service": {
      "description": "Service, for the use of which the payment was made",
      "type": "string",
      "enum": [
        "payment",
        "bonus",
        "sbp"
      ]
    },
    "ServicePrices": {
      "description": "Prices for services for a specific organization in kopecks (RUB * 10^2)",
      "type": "object",
      "required": [
        "bonus",
        "sbp"
      ],
      "properties": {
        "bonus": {
          "type": "integer",
          "format": "int64"
        },
        "sbp": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "Transaction": {
      "type": "object",
      "required": [
        "id",
        "organizationId",
        "createdAt",
        "operation",
        "sevice",
        "amount"
      ],
      "properties": {
        "amount": {
          "description": "Amount in kopecks (RUB * 10^2)",
          "type": "integer",
          "format": "int64",
          "minimum": 1
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "forDate": {
          "type": "string",
          "format": "date",
          "x-nullable": true
        },
        "groupId": {
          "description": "Group that requested payment for using the service",
          "type": "string",
          "format": "uuid",
          "x-nullable": true
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "operation": {
          "$ref": "#/definitions/Operation"
        },
        "organizationId": {
          "type": "string",
          "format": "uuid"
        },
        "sevice": {
          "$ref": "#/definitions/Service"
        },
        "stationsСount": {
          "description": "Number of stations in the car wash that requested payment for using of the service",
          "type": "integer",
          "minimum": 1,
          "x-nullable": true
        },
        "userId": {
          "description": "The user who credited the organisation's account",
          "type": "string",
          "x-nullable": true
        }
      }
    },
    "TransactionPage": {
      "type": "object",
      "required": [
        "items",
        "page",
        "pageSize",
        "totalPages",
        "totalItems"
      ],
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Transaction"
          }
        },
        "page": {
          "type": "integer"
        },
        "pageSize": {
          "type": "integer"
        },
        "totalItems": {
          "type": "integer"
        },
        "totalPages": {
          "type": "integer"
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Bad request",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Forbidden": {
      "description": "Forbidden",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "InternalError": {
      "description": "Internal error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "NotFound": {
      "description": "Not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "authKey": {
      "description": "Session token inside Authorization header.",
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Payment system service for self-service car washes",
    "title": "wash-payment",
    "version": "1.1.0"
  },
  "paths": {
    "/healthCheck": {
      "get": {
        "description": "Checking the server health status.",
        "tags": [
          "Standard"
        ],
        "summary": "Health check",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "ok": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    },
    "/organizations": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "description": "Get a list of organizations",
        "tags": [
          "Organizations"
        ],
        "summary": "Get organizations",
        "operationId": "list",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "default": 1,
            "name": "page",
            "in": "query"
          },
          {
            "maximum": 100,
            "minimum": 1,
            "type": "integer",
            "default": 10,
            "name": "pageSize",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/OrganizationPage"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/organizations/{id}": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "description": "Get information about the specified organization",
        "tags": [
          "Organizations"
        ],
        "summary": "Get organization",
        "operationId": "get",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Organization"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/organizations/{id}/deposit": {
      "post": {
        "security": [
          {
            "authKey": []
          }
        ],
        "description": "Increase the balance of the specified organization by the specified number of kopecks",
        "tags": [
          "Organizations"
        ],
        "summary": "Top up balance",
        "operationId": "deposit",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Deposit"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/organizations/{id}/service-prices": {
      "put": {
        "security": [
          {
            "authKey": []
          }
        ],
        "description": "Set prices for services for the specified organization",
        "tags": [
          "Organizations"
        ],
        "summary": "Set service prices",
        "operationId": "setServicePrices",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/ServicePrices"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/organizations/{id}/transactions": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "description": "Get a list of transactions for the specified organization",
        "tags": [
          "Organizations"
        ],
        "summary": "Get organization transactions",
        "operationId": "transactions",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "default": 1,
            "name": "page",
            "in": "query"
          },
          {
            "maximum": 100,
            "minimum": 1,
            "type": "integer",
            "default": 10,
            "name": "pageSize",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/TransactionPage"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Deposit": {
      "type": "object",
      "required": [
        "amount"
      ],
      "properties": {
        "amount": {
          "description": "Amount in kopecks (RUB * 10^2)",
          "type": "integer",
          "format": "int64",
          "minimum": 1
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Operation": {
      "description": "Type of operation",
      "type": "string",
      "enum": [
        "deposit",
        "debit"
      ]
    },
    "Organization": {
      "type": "object",
      "required": [
        "id",
        "name",
        "displayName",
        "description",
        "balance",
        "servicePrices"
      ],
      "properties": {
        "balance": {
          "description": "Balance in kopecks (RUB * 10^2)",
          "type": "integer",
          "format": "int64",
          "minimum": 0
        },
        "description": {
          "type": "string"
        },
        "displayName": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string"
        },
        "servicePrices": {
          "$ref": "#/definitions/ServicePrices"
        }
      }
    },
    "OrganizationPage": {
      "type": "object",
      "required": [
        "items",
        "page",
        "pageSize",
        "totalPages",
        "totalItems"
      ],
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Organization"
          }
        },
        "page": {
          "type": "integer"
        },
        "pageSize": {
          "type": "integer"
        },
        "totalItems": {
          "type": "integer"
        },
        "totalPages": {
          "type": "integer"
        }
      }
    },
    "Service": {
      "description": "Service, for the use of which the payment was made",
      "type": "string",
      "enum": [
        "payment",
        "bonus",
        "sbp"
      ]
    },
    "ServicePrices": {
      "description": "Prices for services for a specific organization in kopecks (RUB * 10^2)",
      "type": "object",
      "required": [
        "bonus",
        "sbp"
      ],
      "properties": {
        "bonus": {
          "type": "integer",
          "format": "int64",
          "minimum": 0
        },
        "sbp": {
          "type": "integer",
          "format": "int64",
          "minimum": 0
        }
      }
    },
    "Transaction": {
      "type": "object",
      "required": [
        "id",
        "organizationId",
        "createdAt",
        "operation",
        "sevice",
        "amount"
      ],
      "properties": {
        "amount": {
          "description": "Amount in kopecks (RUB * 10^2)",
          "type": "integer",
          "format": "int64",
          "minimum": 1
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "forDate": {
          "type": "string",
          "format": "date",
          "x-nullable": true
        },
        "groupId": {
          "description": "Group that requested payment for using the service",
          "type": "string",
          "format": "uuid",
          "x-nullable": true
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "operation": {
          "$ref": "#/definitions/Operation"
        },
        "organizationId": {
          "type": "string",
          "format": "uuid"
        },
        "sevice": {
          "$ref": "#/definitions/Service"
        },
        "stationsСount": {
          "description": "Number of stations in the car wash that requested payment for using of the service",
          "type": "integer",
          "minimum": 1,
          "x-nullable": true
        },
        "userId": {
          "description": "The user who credited the organisation's account",
          "type": "string",
          "x-nullable": true
        }
      }
    },
    "TransactionPage": {
      "type": "object",
      "required": [
        "items",
        "page",
        "pageSize",
        "totalPages",
        "totalItems"
      ],
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Transaction"
          }
        },
        "page": {
          "type": "integer"
        },
        "pageSize": {
          "type": "integer"
        },
        "totalItems": {
          "type": "integer"
        },
        "totalPages": {
          "type": "integer"
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Bad request",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Forbidden": {
      "description": "Forbidden",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "InternalError": {
      "description": "Internal error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "NotFound": {
      "description": "Not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "authKey": {
      "description": "Session token inside Authorization header.",
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
