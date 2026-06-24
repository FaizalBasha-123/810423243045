# Notification App BE

A Gin microservice that manages notifications with an in-memory store.

## Endpoints

- `POST /notifications` — Create a notification
- `GET /notifications/:id` — Get a notification by ID
- `GET /notifications/user/:userID` — Get all notifications for a user

## Run

```bash
go run main.go
```
