# Vehicle Scheduler BE

Gin microservice on port 8082. Manages vehicles and schedules with in-memory store.

## Endpoints

- `POST /vehicles` — Create a vehicle
- `GET /vehicles` — List all vehicles
- `POST /schedules` — Create a schedule for a vehicle

## Run

```bash
go run main.go
```
