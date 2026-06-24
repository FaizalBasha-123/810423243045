# Campus Evaluation BE

Three Go/Gin microservices:

- **logging-middleware** (port 8083) — Request timing middleware, health check, external log push
- **notification-app-be** (port 8081) — Priority inbox fetcher with weight + timestamp sorting
- **vehicle-scheduler-be** (port 8082) — Vehicle and schedule CRUD with in-memory store

## Run

```powershell
cd <service-dir> && go run main.go
```

## Services

| Service | Port | Auth | Endpoints |
|---------|------|------|-----------|
| logging-middleware | 8083 | Bearer token | GET /health |
| notification-app-be | 8081 | Bearer token | GET /api/v1/priority-inbox?top=N |
| vehicle-scheduler-be | 8082 | — | POST /vehicles, GET /vehicles, POST /schedules |
