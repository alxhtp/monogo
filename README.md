# Monogo

Monogo is a monolithic Go application that exposes a robust, production-ready REST API for user management and more. Built with Clean Architecture principles, it leverages PostgreSQL for data persistence and is designed for scalability, maintainability, and developer productivity.

---

## Table of Contents
- [Project Overview](#project-overview)
- [Architecture](#architecture)
- [Main Features](#main-features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Local Development](#local-development)
  - [Running with Docker](#running-with-docker)
  - [Database Migrations](#database-migrations)
- [Configuration & Environment Variables](#configuration--environment-variables)
- [API Usage Examples](#api-usage-examples)
- [Testing](#testing)
- [Contribution Guidelines](#contribution-guidelines)
- [License](#license)

---

## Project Overview

Monogo is a Go-based monolithic backend that provides a RESTful API on port **8080**. It is designed for rapid development and easy deployment, following Clean Architecture to ensure separation of concerns and testability. The project uses Go modules for dependency management and supports both local and containerized development workflows.

## Architecture

- **Entry Point:** [`cmd/main.go`](cmd/main.go)
- **Clean Architecture Layers:**
  - **Handlers/Controllers:** HTTP request handling and routing
  - **Usecases/Services:** Business logic and application rules
  - **Repositories:** Data access and persistence (PostgreSQL)
  - **Entities/Domain Models:** Core business objects
- **Configuration:** Environment variables (with optional `.env` file)
- **Database:** PostgreSQL (see [`migration/files/`](migration/files/))
- **API Documentation:** Swagger/OpenAPI (`docs/swagger.yaml`, `docs/swagger.json`)
- **Containerization:** Docker & Docker Compose for local and production

## Main Features
- RESTful API (JSON) on port 8080
- User CRUD endpoints (see [API Usage Examples](#api-usage-examples))
- Clean Architecture for maintainability
- PostgreSQL integration
- Database migrations (SQL files)
- Environment-based configuration
- Makefile for common tasks (build, run, lint, test, migrate, docker, clean)
- Docker & Docker Compose support
- Swagger/OpenAPI documentation
- Input validation and error handling

## Getting Started

### Prerequisites
- [Go 1.22+](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)
- [Make](https://www.gnu.org/software/make/)
- [PostgreSQL](https://www.postgresql.org/) (if running DB locally)

### Local Development

1. **Clone the repository:**
   ```sh
   git clone https://github.com/alxhtp/monogo.git
   cd monogo
   ```
2. **Set up environment variables:**
   ```sh
   cp .env.example .env  # Edit as needed
   ```
3. **Build and run the app:**
   ```sh
   make start
   # or, to run directly from source
   make run
   ```
4. **Access the API:**
   - API: http://localhost:8080/v1/
   - Swagger UI: http://localhost:8080/swagger/index.html (if enabled)

### Running with Docker

1. **Start the app and database:**
   ```sh
   make start-docker
   # or
   docker-compose up --build
   ```
2. **Stop the stack:**
   ```sh
   make docker-down
   # or
   docker-compose down
   ```

### Database Migrations

- Migration files are in [`migration/files/`](migration/files/).
- Migrations use [sql-migrate](https://github.com/rubenv/sql-migrate) (install if needed).

**Apply migrations:**
```sh
cd migration
make migrate-up env=local
```

**Rollback last migration:**
```sh
cd migration
make migrate-down env=local
```

**Check migration status:**
```sh
cd migration
make migrate-status env=local
```

---

## Configuration & Environment Variables

Configuration is loaded from environment variables (optionally via `.env`). Key variables include:

| Variable                        | Default         | Description                                 |
|----------------------------------|-----------------|---------------------------------------------|
| `APP_NAME`                      | app             | Application name                            |
| `APP_ENVIRONMENT`               | development     | Environment (development/production)        |
| `APP_DEBUG`                     | false           | Enable debug mode                           |
| `APP_PORT`                      | 8080            | Port to run the API server                  |
| `APP_HOST`                      | localhost       | Host for the API server                     |
| `DB_HOST`                       | localhost       | Database host                               |
| `DB_PORT`                       | 5432            | Database port                               |
| `DB_NAME`                       | app             | Database name                               |
| `DB_USER`                       | user            | Database user                               |
| `DB_PASSWORD`                   | password        | Database password                           |
| `DB_SSL_MODE`                   | disable         | PostgreSQL SSL mode                         |
| `JWT_SECRET_KEY`                | (required)      | JWT secret key                              |
| `JWT_ACCESS_TOKEN_EXPIRY_IN_HOURS` | 1           | JWT access token expiry (hours)             |
| `JWT_REFRESH_TOKEN_EXPIRY_IN_DAYS` | 7           | JWT refresh token expiry (days)             |
| `SWAGGER_USERNAME`              | (required)      | Swagger UI basic auth username              |
| `SWAGGER_PASSWORD`              | (required)      | Swagger UI basic auth password              |
| ...                             |                 | See [`config/config.go`](config/config.go)  |

---

## API Usage Examples

### Create User
```sh
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice",
    "email": "alice@example.com",
    "metadata": {
      "sex": "female",
      "address": "123 Main St",
      "phone": "+1234567890"
    }
  }'
```

### Get Users (with filters)
```sh
curl "http://localhost:8080/v1/users?ids=...&name=Alice&email=alice@example.com&status=1&sex=female&address=123+Main+St&phone=%2B1234567890&include-deleted=false&show-count=true&offset=0&limit=10&order-by=+created_at"
```

### Get User by ID
```sh
curl http://localhost:8080/v1/users/{id}
```

### Update User
```sh
curl -X PUT http://localhost:8080/v1/users/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice Updated",
    "email": "alice.updated@example.com"
  }'
```

### Delete User
```sh
curl -X DELETE http://localhost:8080/v1/users/{id}
```

---

## Testing

- **Run tests:**
  ```sh
  go test ./...
  ```
- **Lint code:**
  ```sh
  make lint
  ```

---

## Contribution Guidelines

1. Fork the repository and create your branch from `master`.
2. Ensure code is formatted (`go fmt`), linted, and tested.
3. Write clear, descriptive commit messages.
4. Add tests for new features and bug fixes.
5. Open a pull request and describe your changes.

---

## License

This project is licensed under the [Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0.html).
