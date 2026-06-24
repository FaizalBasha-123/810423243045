git add README.md .gitignore .env go.mod go.sum
git commit -m "chore: initialize project structure and base configurations"

git add logging-middleware/
git commit -m "feat: implement custom logging middleware and health checks"

git add notification-app-be/models/ notification-app-be/config/
git commit -m "feat: define core data structures and environment configuration"

git add notification-app-be/services/ notification-app-be/handlers/
git commit -m "feat: implement in-memory state management and core processing logic"

git add notification-app-be/routes/ notification-app-be/main.go
git commit -m "feat: wire http routes and initialize gin engine"

git add notification-system-design.md
git commit -m "docs: draft system architecture and scalability strategies"