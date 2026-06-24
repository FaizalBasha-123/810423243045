# Notification System Design

## Stage 1

Logging middleware service on port 8083 with health check and request timing middleware.

---

## Stage 2

Notification CRUD service on port 8081 with in-memory store.

---

## Stage 3

Vehicle scheduler service on port 8082 with vehicle and schedule management.

---

## Stage 4

External auth and logging integration via AuthenticateClient() and Log().

---

## Stage 5

Priority inbox endpoint with external API fetch, weight-based sorting (Placement > Result > Event), and timestamp tie-breaker.

---

## Stage 6

Refactored to standalone handler under /api/v1/priority-inbox with stdlib-only service layer.
