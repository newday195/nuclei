package input

import "github.com/xeipuuv/gojsonschema"

var schemaLoader = gojsonschema.NewBytesLoader(schema)
var schema = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "https://projectdiscovery.io/nuclei-template.schema.json",
    "title": "Nuclei Templates YAML format validation",
    "description": "A validation format for validating nuclei-template files.",
    "required": [ "id", "info" ],
    "type": "object",
    "properties": {
      "id": {
        "type": "string",
        "pattern": "^[a-z0-9\\-]{1,32}$"
      },
      "info": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string"
          },
          "author": {
            "type": "string"
          },
          "severity": {
            "type": "string"
          },
          "description": {
            "type": "string"
          }
        }
      }
    }
  }
`)
