# GEMINI.md - Project Context & Instructions

## Project Overview

**go-boilerplate** is a comprehensive Go-based boilerplate designed for building scalable and maintainable RESTful APIs. It follows a Clean Architecture approach, separating concerns into distinct layers: Presentation (REST), Application (Use Cases), Domain, and Infrastructure (Repositories/Persistence).

### Core Technologies

- **Language:** Go 1.25.1
- **Web Framework:** [Fiber v2](https://gofiber.io/)
- **ORM:** [GORM](https://gorm.io/) with PostgreSQL
- **Caching:** [Redis](https://redis.io/)
- **Configuration:** [Viper](https://github.com/spf13/viper) (YAML) and [Netflix go-env](https://github.com/Netflix/go-env) (Environment Variables)
- **Migrations:** [golang-migrate](https://github.com/golang-migrate/migrate)
- **Documentation:** OpenAPI/Swagger (via `swagger-cli`)
- **Logging:** [Zerolog](https://github.com/rs/zerolog)
- **Testing:** [Testify](https://github.com/stretchr/testify) and [Mockgen](https://github.com/golang/mock)

## Project Architecture

The project is organized into the following layers within the `src/` directory:

- **`src/rest/`**: Presentation layer. Contains Fiber handlers, route definitions, and middleware.
- **`src/usecases/`**: Application layer. Implements business logic and defines Data Transfer Objects (DTOs).
- **`src/domain/`**: Domain layer. Contains core business entities and domain-specific logic.
- **`src/repo/`**: Infrastructure layer (Data Access). Implements the repository pattern for database and cache interactions.
- **`src/models/`**: Database schema definitions used by GORM.
- **`src/di/`**: Dependency Injection. Uses `sync.OnceValue` to provide singleton instances of services, use cases, and repositories.
- **`src/pkg/`**: Shared infrastructure and utility packages (e.g., database connections, Redis client, custom HTTP utilities).

## Building and Running

### Prerequisites

- Go 1.25.1+
- Docker & Docker Compose (for Postgres and Redis)
- `golang-migrate` CLI (for manual migration tasks)
- `mockgen` (for generating mocks)
- `swagger-cli` (for generating documentation)

### Key Commands

All major tasks are managed via the `Makefile`:

| Task                 | Command                               | Description                                       |
| :------------------- | :------------------------------------ | :------------------------------------------------ |
| **Run Server**       | `make serve`                          | Starts the API server using `cmd/main.go`.        |
| **Test**             | `make test`                           | Runs all tests in the project.                    |
| **Lint**             | `make lint`                           | Runs `golangci-lint` with the `--fix` flag.       |
| **Migrations Up**    | `make migration-up`                   | Applies all pending database migrations.          |
| **Migrations Down**  | `make migration-down`                 | Reverts the last applied migration.               |
| **New Migration**    | `make migration-generate name=<name>` | Generates a new SQL migration file.               |
| **Generate Mocks**   | `make gen-mock`                       | Generates mocks for all files in `src/repo/`.     |
| **Generate Swagger** | `make swagger-generate`               | Bundles OpenAPI specs into a single Swagger YAML. |

## Development Conventions

### Coding Style

- Follow standard Go idioms and `gofmt` formatting.
- Use `golangci-lint` for static analysis (config found in `.golangci.yml`).
- Dependency Injection should be managed through the `src/di/` package to ensure consistent instantiation and testability.

### Database & Migrations

- Use GORM for database interactions.
- Schema changes must be performed via migrations in `src/migrations/`.
- Use the `make migration-generate` command to create new migration files.

### Testing

- Write unit tests for use cases and domain logic.
- Use `testify/assert` or `testify/require` for assertions.
- Use `mockgen` to create mocks for repositories when testing use cases. Mocks are stored in `src/repo/mocks/`.

### Documentation

- The project uses OpenAPI/Swagger.
- Update `docs/openapi.yaml` when changing API contracts.
- Run `make swagger-generate` to update the compiled documentation at `docs/compile/swagger.yaml`.
