# go-microservices

This repository contains a **microservices architecture** implemented in Go, including:

- **SSO service** (`sso`) – a gRPC authentication and authorization service.
- **Protos** (`protos`) – Protocol Buffers definitions for gRPC communication.
- **URL Shortener** (`url-shortener`) – a web service that integrates with the SSO service via gRPC.

---

## Project Structure


### `sso`
The SSO service provides:

- User registration and login.
- Admin role verification.
- Integration with SQLite for storage.
- gRPC API exposed for external clients.

### `protos`
- Contains all `.proto` files used to generate gRPC clients and servers.
- Includes generated Go files in `gen/go`.
- Ensures type-safe communication between services.

### `url-shortener`
- A simple URL shortening service.
- Integrates with the SSO gRPC service for user authentication and authorization.
- Stores shortened URLs in SQLite.
- Demonstrates microservice interaction using gRPC.

---

### Prerequisites
- Go 1.25+
- `protoc` 33.x (Protocol Buffers compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins
- SQLite (for local development)
