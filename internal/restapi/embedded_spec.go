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
    "title": "Library API",
    "version": "1.0"
  },
  "paths": {
    "/books": {
      "get": {
        "summary": "Get list of all books.",
        "operationId": "getAllBooks",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "books": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/book"
                  }
                }
              }
            }
          },
          "500": {
            "description": "Internal error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "summary": "Create book record.",
        "operationId": "createBook",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/book"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/book"
            }
          },
          "400": {
            "description": "Bad argument.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/books/{id}": {
      "get": {
        "summary": "Returns book by id.",
        "operationId": "getBook",
        "parameters": [
          {
            "type": "string",
            "description": "book id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/book"
            }
          },
          "404": {
            "description": "Not found.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete book record.",
        "operationId": "deleteBook",
        "parameters": [
          {
            "type": "string",
            "description": "book id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "404": {
            "description": "Not found.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "summary": "Update book record.",
        "operationId": "updateBook",
        "parameters": [
          {
            "type": "string",
            "description": "book id",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "bookUpdate",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "rating": {
                  "type": "integer"
                },
                "status": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad argument.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not found.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "book": {
      "type": "object",
      "required": [
        "title",
        "author",
        "publisher",
        "publication_date"
      ],
      "properties": {
        "author": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "publication_date": {
          "type": "string",
          "format": "date"
        },
        "publisher": {
          "type": "string"
        },
        "rating": {
          "type": "string",
          "enum": [
            "rate1",
            "rate2",
            "rate3"
          ]
        },
        "status": {
          "type": "string",
          "enum": [
            "checked_in",
            "checked_out"
          ]
        },
        "title": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
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
    "title": "Library API",
    "version": "1.0"
  },
  "paths": {
    "/books": {
      "get": {
        "summary": "Get list of all books.",
        "operationId": "getAllBooks",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "books": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/book"
                  }
                }
              }
            }
          },
          "500": {
            "description": "Internal error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "summary": "Create book record.",
        "operationId": "createBook",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/book"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/book"
            }
          },
          "400": {
            "description": "Bad argument.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/books/{id}": {
      "get": {
        "summary": "Returns book by id.",
        "operationId": "getBook",
        "parameters": [
          {
            "type": "string",
            "description": "book id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/book"
            }
          },
          "404": {
            "description": "Not found.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete book record.",
        "operationId": "deleteBook",
        "parameters": [
          {
            "type": "string",
            "description": "book id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "404": {
            "description": "Not found.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "summary": "Update book record.",
        "operationId": "updateBook",
        "parameters": [
          {
            "type": "string",
            "description": "book id",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "bookUpdate",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "rating": {
                  "type": "integer"
                },
                "status": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad argument.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not found.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "book": {
      "type": "object",
      "required": [
        "title",
        "author",
        "publisher",
        "publication_date"
      ],
      "properties": {
        "author": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "publication_date": {
          "type": "string",
          "format": "date"
        },
        "publisher": {
          "type": "string"
        },
        "rating": {
          "type": "string",
          "enum": [
            "rate1",
            "rate2",
            "rate3"
          ]
        },
        "status": {
          "type": "string",
          "enum": [
            "checked_in",
            "checked_out"
          ]
        },
        "title": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
}
