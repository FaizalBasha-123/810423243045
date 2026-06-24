# Notification App BE

Gin microservice on port 8081. Fetches priority inbox from external API with weight-based and timestamp sorting.

## Endpoints

- `GET /api/v1/priority-inbox?top=N` — Returns top N sorted notifications

## Sorting

1. By type weight: Placement (3) > Result (2) > Event (1)
2. By timestamp (newest first) when weights tie

## Run

```bash
go run main.go
```
