# Photon Swagger Starter

A Swagger UI integration starter for the Photon framework, providing easy API documentation capabilities.


## Installation

```bash
go get github.com/Phofuture/photon-swagger-starter
```

## Usage

Import the package in your application:

```go
import _ "github.com/Phofuture/photon-swagger-starter"
```

The starter will automatically register itself with the Photon core framework.

## Configuration

Configure Swagger settings in your application configuration

example:

```yaml
swagger:
  enabled: true
  context-path: "/api"
  directory-path: "./docs"
```

### Configuration Options

- `enabled`: Enable or disable Swagger UI (default: false)
- `context-path`: Base path for Swagger endpoints
- `directory-path`: Directory containing swagger.json/swagger.yaml files

## Endpoints

Once configured, Swagger UI will be available at:

- `{context-path}/swagger/*` - Swagger UI interface
- `{context-path}/swagger/swagger.json` - JSON swagger specification
- `{context-path}/swagger/swagger.yaml` - YAML swagger specification

## Dependencies

- [Fiber v2](https://github.com/gofiber/fiber) - Web framework
- [Fiber Swagger](https://github.com/gofiber/swagger) - Swagger middleware
- [Photon Core Starter](https://github.com/Phofuture/photon-core-starter) - Core framework

## Requirements

- Go 1.23.0 or later

## License

This project is part of the Phofuture organization.