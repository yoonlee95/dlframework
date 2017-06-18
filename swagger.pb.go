package dlframework 

const (
dlframework_swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "CarML DLFramework",
    "version": "1.0.0",
    "description": "TODO... fillme."
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/framework/{framework_name}/info": {
      "get": {
        "operationId": "GetFrameworkManifest",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/dlframeworkFrameworkManifest"
            }
          }
        },
        "parameters": [
          {
            "name": "framework_name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "framework_version",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "Registry"
        ]
      }
    },
    "/v1/framework/{framework_name}/model/{model_name}/info": {
      "post": {
        "operationId": "GetFrameworkModelManifest",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/dlframeworkModelManifest"
            }
          }
        },
        "parameters": [
          {
            "name": "framework_name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "model_name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkGetFrameworkModelManifestRequest"
            }
          }
        ],
        "tags": [
          "Registry"
        ]
      }
    },
    "/v1/framework/{framework_name}/model/{model_name}/predict": {
      "post": {
        "operationId": "Predict",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "framework_name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "model_name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictRequest"
            }
          }
        ],
        "tags": [
          "Predictor"
        ]
      }
    },
    "/v1/framework/{framework_name}/models": {
      "get": {
        "operationId": "GetFrameworkModels",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/dlframeworkGetModelManifestsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "framework_name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "framework_version",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "Registry"
        ]
      }
    },
    "/v1/frameworks": {
      "get": {
        "operationId": "GetFrameworkManifests",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/dlframeworkGetFrameworkManifestsResponse"
            }
          }
        },
        "tags": [
          "Registry"
        ]
      }
    },
    "/v1/model/{model_name}/info": {
      "post": {
        "operationId": "GetModelManifest",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/dlframeworkModelManifest"
            }
          }
        },
        "parameters": [
          {
            "name": "model_name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkGetModelManifestRequest"
            }
          }
        ],
        "tags": [
          "Registry"
        ]
      }
    },
    "/v1/models": {
      "get": {
        "operationId": "GetModelManifests",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/dlframeworkGetModelManifestsResponse"
            }
          }
        },
        "tags": [
          "Registry"
        ]
      }
    }
  },
  "definitions": {
    "ModelManifestModel": {
      "type": "object",
      "properties": {
        "base_url": {
          "type": "string"
        },
        "weights_path": {
          "type": "string"
        },
        "graph_path": {
          "type": "string"
        },
        "is_archive": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "TypeParameter": {
      "type": "object",
      "properties": {
        "value": {
          "$ref": "#/definitions/protobufStruct"
        }
      }
    },
    "dlframeworkContainerHardware": {
      "type": "object",
      "properties": {
        "gpu": {
          "type": "string"
        },
        "cpu": {
          "type": "string"
        }
      }
    },
    "dlframeworkErrorStatus": {
      "type": "object",
      "properties": {
        "ok": {
          "type": "boolean",
          "format": "boolean"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "dlframeworkFrameworkManifest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "container": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/dlframeworkContainerHardware"
          }
        }
      }
    },
    "dlframeworkGetFrameworkManifestsResponse": {
      "type": "object",
      "properties": {
        "manifests": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkFrameworkManifest"
          }
        }
      }
    },
    "dlframeworkGetFrameworkModelManifestRequest": {
      "type": "object",
      "properties": {
        "framework_name": {
          "type": "string"
        },
        "framework_version": {
          "type": "string"
        },
        "model_name": {
          "type": "string"
        },
        "model_version": {
          "type": "string"
        }
      }
    },
    "dlframeworkGetModelManifestRequest": {
      "type": "object",
      "properties": {
        "model_name": {
          "type": "string"
        },
        "model_version": {
          "type": "string"
        }
      }
    },
    "dlframeworkGetModelManifestsResponse": {
      "type": "object",
      "properties": {
        "manifests": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkModelManifest"
          }
        }
      }
    },
    "dlframeworkModelManifest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "framework": {
          "$ref": "#/definitions/dlframeworkFrameworkManifest"
        },
        "container": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/dlframeworkContainerHardware"
          }
        },
        "description": {
          "type": "string"
        },
        "reference": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "license": {
          "type": "string"
        },
        "inputs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkModelManifestType"
          }
        },
        "output": {
          "$ref": "#/definitions/dlframeworkModelManifestType"
        },
        "before_preprocess": {
          "type": "string"
        },
        "preprocess": {
          "type": "string"
        },
        "after_preprocess": {
          "type": "string"
        },
        "before_postprocess": {
          "type": "string"
        },
        "postprocess": {
          "type": "string"
        },
        "after_postprocess": {
          "type": "string"
        },
        "model": {
          "$ref": "#/definitions/ModelManifestModel"
        },
        "attributes": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    },
    "dlframeworkModelManifestType": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "parameters": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/TypeParameter"
          }
        }
      }
    },
    "dlframeworkPredictRequest": {
      "type": "object",
      "properties": {
        "model_name": {
          "type": "string"
        },
        "model_version": {
          "type": "string"
        },
        "framework_name": {
          "type": "string"
        },
        "framework_version": {
          "type": "string"
        },
        "limit": {
          "type": "integer",
          "format": "int32"
        },
        "data": {
          "type": "string",
          "format": "byte"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "dlframeworkPredictResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "features": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkPredictionFeature"
          }
        },
        "error": {
          "$ref": "#/definitions/dlframeworkErrorStatus"
        }
      }
    },
    "dlframeworkPredictionFeature": {
      "type": "object",
      "properties": {
        "index": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "probability": {
          "type": "number",
          "format": "float"
        }
      }
    },
    "protobufListValue": {
      "type": "object",
      "properties": {
        "values": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/protobufValue"
          },
          "description": "Repeated field of dynamically typed values."
        }
      },
      "description": "'ListValue' is a wrapper around a repeated field of values.\n\nThe JSON representation for 'ListValue' is JSON array."
    },
    "protobufNullValue": {
      "type": "string",
      "enum": [
        "NULL_VALUE"
      ],
      "default": "NULL_VALUE",
      "description": "'NullValue' is a singleton enumeration to represent the null value for the\n'Value' type union.\n\n The JSON representation for 'NullValue' is JSON 'null'.\n\n - NULL_VALUE: Null value."
    },
    "protobufStruct": {
      "type": "object",
      "properties": {
        "fields": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/protobufValue"
          },
          "description": "Unordered map of dynamically typed values."
        }
      },
      "description": "'Struct' represents a structured data value, consisting of fields\nwhich map to dynamically typed values. In some languages, 'Struct'\nmight be supported by a native representation. For example, in\nscripting languages like JS a struct is represented as an\nobject. The details of that representation are described together\nwith the proto support for the language.\n\nThe JSON representation for 'Struct' is JSON object."
    },
    "protobufValue": {
      "type": "object",
      "properties": {
        "null_value": {
          "$ref": "#/definitions/protobufNullValue",
          "description": "Represents a null value."
        },
        "number_value": {
          "type": "number",
          "format": "double",
          "description": "Represents a double value."
        },
        "string_value": {
          "type": "string",
          "description": "Represents a string value."
        },
        "bool_value": {
          "type": "boolean",
          "format": "boolean",
          "description": "Represents a boolean value."
        },
        "struct_value": {
          "$ref": "#/definitions/protobufStruct",
          "description": "Represents a structured value."
        },
        "list_value": {
          "$ref": "#/definitions/protobufListValue",
          "description": "Represents a repeated 'Value'."
        }
      },
      "description": "'Value' represents a dynamically typed value which can be either\nnull, a number, a string, a boolean, a recursive struct value, or a\nlist of values. A producer of value is expected to set one of that\nvariants, absence of any variant indicates an error.\n\nThe JSON representation for 'Value' is JSON value."
    }
  }
}
`
swagger_info = `{
  "info": {
    "title": "CarML DLFramework",
    "description": "TODO... fillme.",
    "version": "1.0.0"
  }
}
`
)
