# Logging Middleware Service

Gin microservice on port 8083. Logs request duration per endpoint and pushes structured logs to external evaluation server.

## Endpoints

- `GET /health` — Health check (triggers external log)

## Middleware

Logger in `middleware/logger.go` measures and logs method, path, client IP, and duration.

## Run

```bash
go run main.go
```
