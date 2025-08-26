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

1. Clone the project and get into the project folder
    ```
    git clone https://github.com/bwindrasmoro/config-management-service.git && cd config-management-service
    ```

2. Build the binary
    ```
    make build
    ```
    or
    ```
    go build -o config-management-service main.go (for Linux/Mac)
    ```
    ```
    go build -o config-management-service.exe main.go (for Windows)
    ```

3. Run the binary
   ```
   make run
   ```
   or
   ```
   ./config-management-service (for Linux/Mac)
   ```
   ```
   config-management-service.exe (for Linux/Mac)
   ```


4. Run unit testing
   ```
   make test
   ```
   or
   ```
   go test ./test
   ```
