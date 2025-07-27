# Go gPRC Starter

> A fully-featured, production-ready gRPC boilerplate in Go — complete with JWT authentication, TLS, validation, error handling, Envoy REST gateway, and more.

---

## Description

This repository demonstrates a **production-grade gRPC service** built in Go, showcasing best practices in secure communication, modular code structure, and robust error handling. It includes tools and configurations for real-world deployment, local development, and REST-to-gRPC bridging via Envoy.

Whether you're starting a new microservice or modernizing an existing monolith, this boilerplate is a clean foundation for scalable, performant, and secure Go-based services.

---

## Features

### Authentication

- JWT-based auth using `RS256` with PEM key pair (`private.pem` / `public.pem`)
- Tokens include standard and custom claims (`sub`, `iat`, `role`, etc.)
- Server-side verification via `UnaryServerInterceptor`

### Metadata Handling

- Sends and parses custom metadata headers (e.g. Authorization token)
- Uses `context.Context` with outgoing/incoming metadata

### Custom Error Handling

- Structured error response using custom `Error` protobuf
- Wrapped with `google.golang.org/grpc/status` and `codes`
- Uniform client-side error decoding with details

### Field Validation

- Custom validation logic in handler layer
- Example: Rejects empty names with gRPC errors

### TLS / SSL (Secure Transport)

- Self-signed certs (`cert.pem`, `key.pem`) for local testing
- Supports replacement with production-grade certs without code change
- gRPC client configured with `credentials.NewClientTLSFromFile`

### REST API via Envoy Proxy

- REST-to-gRPC translation using Envoy
- Accepts standard HTTP/JSON requests and forwards to gRPC backend
- No BSR (Buf Schema Registry) required

### Context & Deadlines

- All client-server requests use `context.Context`
- Deadline propagation supported with timeouts

### Interceptors

- Unary interceptors for:
  - Logging
  - Authentication
  - Error translation

### Shared Protobuf Definitions

- Shared `proto/` folder between client and server
- Uses `buf` for linting, breaking checks, and codegen

### Build Tooling

- `Makefile` includes:
  - `make proto`: Compile all protos
  - `make run-server`: Start the gRPC server
  - `make run-client`: Run the test client

### Client Example

- Client includes:
  - Auth token generation via `jwt-go`
  - TLS transport
  - Metadata headers
  - Calls for `Create`, `Read`, `Write`, and `Delete`

### Configuration

- Centralized configuration via `config.yml`
- Share JWT keys, TLS cert paths, and server address between components

---

## Quickstart

### Prerequisites

- Go ≥ 1.23
- `buf` (https://buf.build/)
- `make`
- `docker` (for Envoy REST proxy)

### 1. Generate Protobufs

```bash
make proto
```

### 2. Run Server

```bash
make run-server
```

### 3. Run Client

```bash
make run-client
```

### 4. REST Testing via Envoy

**This project includes a Docker-based Envoy proxy to bridge HTTP/JSON REST requests to your gRPC backend**

Steps to Run Envoy:

- Make sure the gRPC server is running on port 50051.
- Build the Envoy Docker image from the provided Dockerfile:

```bash
docker build -t go-grpc-envoy -f envoy/Dockerfile .
```

- Run the container:

```bash
docker run --rm -p 8080:8080 --network host go-grpc-envoy
```

> - **--network host** allows Envoy to communicate with the gRPC server on **localhost:50051**
> - If you're on macOS or Windows, replace **--network host** with appropriate Docker networking, or use **host.docker.internal** in the envoy.yaml.

- Test using curl:

```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Authorization: Bearer <your_jwt>" \
  -H "Content-Type: application/json" \
  -d '{"name": "test", "email": "test@test.com", "number": "1234567890"}'
```

## TLS/SSL Notes

| Environment | Setup                                                      |
| ----------- | ---------------------------------------------------------- |
| Local Dev   | Uses self-signed `cert.pem` and `key.pem`                  |
| Production  | Replace certs in `certs/` folder, no code changes required |
| Client      | Uses server’s `cert.pem` as CA cert                        |

## JWT Auth Notes

| File          | Purpose                   |
| ------------- | ------------------------- |
| `private.pem` | Signing key (client)      |
| `public.pem`  | Verification key (server) |

- Tokens issued by client via RSA
- Claims are validated server-side via custom interceptor
