# Configuration Management Service

A Go-based service to manage application configurations with validation, versioning, and rollback.

## Features
- Create, update, rollback, and fetch configurations
- JSON schema validation
- Version tracking

## Installation
### Prerequisite
- Go 1.21+
- Makefile

### How to Build and Run
Make sure your environment has Go installed. If you haven't installed Go in your environment, then you should [install Go](https://go.dev/doc/install) first.

1. **Clone** the project and get into the project folder
    ```
    git clone https://github.com/bwindrasmoro/config-management-service.git && cd config-management-service
    ```

2. **Build** the binary
    - Using **Makefile**
        ```
        make build
        ```
    or
    - **Linux/MAC**
      ```
      go build -o config-management-service main.go
      ```
    - **Windows**
      ```
      go build -o config-management-service.exe main.go
      ```

3. **Run** the binary

    This application will be run with port **3001**
    - Using **Makefile**
        ```
        make run
        ```
    or
    - **Linux/MAC**
        ```
        ./config-management-service
        ```
    - **Windows**
        ```
        config-management-service.exe 
        ```

4. **Optional**: Running without create binary
    ```
    go run main.go
    ```


5. Run unit testing

    Using **Makefile**
    ```
    make test
    ```
    or
    ```
    go test ./test
    ```

## API Documentation

<details>
<summary>OpenAPI Spec</summary>

```yaml
openapi: 3.0.3
info:
  title: Config Management Service
  version: "1.0"
  description: This service handle create, update, and listing configuration
servers:
  - url: http://localhost:3001/api/v1/config
paths:
  /:
    post:
      summary: Create new config
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                config_name:
                  type: string
                  example: "payment_config"
                data:
                  $ref: "#/components/schemas/PaymentConfig"
      responses:
        "200":
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: string
                    example: "Success"
                  data:
                    $ref: "#/components/schemas/PaymentConfig"
                  version:
                    type: integer
                    example: 1
                  message:
                    type: string
                    example: "Config sucessfully created"
        "400":
          content:
            application/json:
              schema:
                oneOf:
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "One or more field are required"
                      data:
                        $ref: "#/components/schemas/MissingField"
                  - type: object
                    properties:
                    status:
                      type: string
                      example: "Fail"
                    message:
                      type: string
                      example: "Invalid request body"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Unsupported config type: not_exists_schema"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "invalid config data"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Config already exists"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Failed to unmarshal data"                
  /{schema}:
    post:
      summary: Update existing config
      parameters:
        - name: schema
          in: path
          required: true
          description: Schema Name
          schema:
            type: string
            example: "payment_config"
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/PaymentConfig"                  
      responses:
        "200":
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: string
                    example: "Success"
                  data:
                    $ref: "#/components/schemas/PaymentConfig"
                  version:
                    type: integer
                    example: 1
                  message:
                    type: string
                    example: "payment_config sucessfully updated"
        "400":
          content:
            application/json:
              schema:
                oneOf:
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "One or more field are required"
                      data:
                        $ref: "#/components/schemas/MissingField"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Schema not exists"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Config for this schema not exists"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Failed to parse config"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Unsupported config type: not_exists_schema"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "invalid config data"
                
    get:
      summary: Retrieve latest config of the schema
      parameters:
        - name: schema
          in: path
          required: true
          description: Schema Name
          schema:
            type: string
            example: "payment_config"
      responses:
        "200":
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: string
                    example: "Success"
                  data:
                    $ref: "#/components/schemas/PaymentConfig"
                  version:
                    type: integer
                    example: 1
        "400":
          content:
            application/json:
              schema:
                oneOf:
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Schema not exists"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Configuration for payment_config schema not found"                
  /:schema/rollback/:version:
    post:
      summary: Rollback to previous version
      parameters:
        - name: schema
          in: path
          required: true
          description: Schema Name
          schema:
            type: string
            example: "payment_config"
        - name: version
          in: path
          required: true
          description: Config version
          schema:
            type: integer
            example: 2
      responses:
        "200":
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: string
                    example: "Success"
                  data:
                    $ref: "#/components/schemas/PaymentConfig"
                  version:
                    type: integer
                    example: 1
                  message:
                    type: string
                    example: "payment_config sucessfully updated"
        "400":
          content:
            application/json:
              schema:
                oneOf:
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Old Version must be Integer"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Schema not exists"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Cannot rollback payment_config schema: rollback version 17 not exists"
  /:schema/versions:
    get:
      summary: Retrieve all history versions config of the schema
      parameters:
        - name: schema
          in: path
          required: true
          description: Schema Name
          schema:
            type: string
            example: "payment_config"
      responses:
        "200":
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: string
                    example: "Success"
                  data:
                    type: array
                    items:
                      $ref: "#/components/schemas/PaymentConfig"
        "400":
          content:
            application/json:
              schema:
                oneOf:
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Schema not exists"
                  - type: object
                    properties:
                      status:
                        type: string
                        example: "Fail"
                      message:
                        type: string
                        example: "Configuration for payment_config schema not found"
                
components:
  schemas:
    PaymentConfig:
      type: object
      properties:
        max_limit:
          type: int
          example: 1000
        enabled:
          type: boolean
          example: false
    MissingField:
      type: array
      items:
        field_name:
          type: string
          example: "field_name"

```
</details>

### Endpoint, Request and Response

#### Base URL: ```http://localhost:3001/api/v1/config```
#### Endpoints
- **POST /**
    - Description: Create new config
    - Request:
        ```
        {
            "config_name": "payment_config",
            "data": {
                "max_limit": 1000,
                "enabled": true
            }
        }
        ```
    - Response:
        ```
        {
            "status": "Success",
            "data": {
                "enabled": true,
                "max_limit": 1000
            },
            "version": 1,
            "message": "Config sucessfully created"
        }
        ```
- **POST /:schema**
    - Description: Update existing config
    - Request:
        ```
        {
            "max_limit": 2000,
            "enabled": true
        }
        ```
    - Response:
        ```
        {
            "status": "Success",
            "data": {
                "max_limit": 2000,
                "enabled": true
            },
            "version": 2,
            "message": "payment_config sucessfully updated"
        }
        ```
- **POST /:schema/rollback/:version**
    - Description: Rollback to previous version
    - Parameters:
        - schema:
            - description: schema name
            - type: string
        - version:
            - description: selected version
            - type: integer
    - Response:
        ```
        {
            "status": "Success",
            "data": {
                "max_limit": 2000,
                "enabled": true
            },
            "version": 5,
            "message": "payment_config sucessfully updated"
        }
        ```
- **GET /:schema**
    - Description: Retrieve latest config of the schema
    - Parameters:
        - schema:
            - description: schema name
            - type: string
    - Response:
        ```
        {
            "status": "Success",
            "data": {
                "max_limit": 4000,
                "enabled": true
            },
            "version": 4
        }
        ```
- **GET /:schema/versions**
    - Description: Retrieve latest config of the schema
    - Parameters:
        - schema:
            - description: schema name
            - type: string
    - Response:
        ```
        {
            "status": "Success",
            "data": [
                {
                    "config": {
                        "enabled": true,
                        "max_limit": 1000
                    },
                    "version": 1
                },
                {
                    "config": {
                        "max_limit": 2000,
                        "enabled": true
                    },
                    "version": 2
                },
                {
                    "config": {
                        "max_limit": 3000,
                        "enabled": true
                    },
                    "version": 3
                },
                {
                    "config": {
                        "max_limit": 4000,
                        "enabled": true
                    },
                    "version": 4
                },
                {
                    "config": {
                        "max_limit": 2000,
                        "enabled": true
                    },
                    "version": 5
                }
            ]
        }
        ```

## Schema Explanation
The schema must be declared using go struct, here is the example:
```go
type PaymentConfig struct {
	MaxLimit int  `json:"max_limit" validation:"required"`
	Enabled  bool `json:"enabled" validation:"required"`
}
```

The schema validation will be based on the validation tag in the struct. 

* **Required** tag means that the field is required, unable to provide data to this field will return status code 400 with a list of variable that must be filled

* Unable to provide data with respected type will return status code 400 with message Invalid request body or Invalid config data


## Improvement
There are several things that can be improved:
1. Authentication/authorization
2. Logging and metric
