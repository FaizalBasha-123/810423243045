# Logging Middleware Service

A standalone Gin microservice that provides a custom HTTP request logging middleware.

## Endpoints

- `GET /health` — Health check

## Middleware

The `Logger` middleware in `middleware/logger.go` measures request duration and logs the method, path, client IP, and time taken using the standard `log` package.

## Run

```bash
go run main.go
```
