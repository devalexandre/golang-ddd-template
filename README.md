# `golang-ddd-template` Repository Structure

This repository follows the Domain-Driven Design (DDD) pattern and is organized as follows:

## Project Root
- `Dockerfile`: Defines the Docker image build for the project.
- `docker-compose.yml`: Configures Docker container orchestration.
- `go.mod` and `go.sum`: Manage Go project dependencies.
- `cmd/main.go`: The entry point of the application.

## `internal` Directory
Contains the internal logic of the application, divided into several subsections.

### `domain` Subdirectory
Where the core business logic is centralized.
- `internal/domain/user/`
  - `contracts.go`: Interfaces for the user repository and service.
  - `factory.go`: Factory for creating user instances.
  - `repository.go`: Implementation of the user repository.
  - `service.go`: Implementation of the user service.

### `helpers` Subdirectory
Provides helper functions and utilities.
- `internal/helpers/`
  - `config/config.go`: Application configuration management.
  - `errors.go`: Structures and functions related to errors.

### `infra` Subdirectory
Specific infrastructure implementations.
- `internal/infra/pgx/`
  - `connection.go`: Manages database connections.
  - `contracts.go`: Interfaces for the infrastructure layer.

## Tests
Test files are named with the `_test.go` suffix and are located in the same directories as the files they test.



Each part of the structure is designed to keep the code organized, modular, and focused on business rules.


```
├── cmd
│   └── main.go
├── internal
│   ├── domain
│   │   ├── user
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   │   ├── factory.go
│   │   │   ├── contracts.go
│   │   │   ├── repository_test.go
│   │   │   ├── service_test.go
│   │   │   ├── factory_test.go
│   │   │   └── mock.go
│   ├── infra
│   │   ├── database
│   │   │   ├── mysql
│   │   │   │   ├── connect.go
│   │   │   │   ├── contracts.go
│   │   │   │   └── mock.go
│   │   ├── openapi
│   │   │   ├── openapi.go
│   │   │   ├── openapi_test.go
│   │   │   ├── contracts.go
│   │   │   └── mock.go
│   ├── helpers
│   │   ├── errors
│   │   │   ├── errors.go
│   │   │   └── errors_test.go
│   │   ├── logger
│   │   │   ├── logger.go
│   │   │   └── logger_test.go
│   │   ├── config
│   │   │   ├── environment.go
│   │   │   └── environment_test.go

# LLM Context - path/to/folder

## Responsibility

Describe the responsibility of this folder in one or two short paragraphs.

## Important Files

| File | Purpose | Open When |
|---|---|---|
| `example.go` | Describe what this file does | Describe when an agent should open it |

## Main Flows

1. Describe the main input.
2. Describe the main processing step.
3. Describe the main output.
4. Mention important dependencies.

## Architecture Rules

- Rule 1
- Rule 2
- Rule 3

## Dependencies

This folder depends on:

- package A
- package B

This folder should not depend on:

- package X
- package Y

## Common Changes

Use this section to help agents decide what files to open.

### To change business behavior

Open:

- `service.go`
- `entity.go`

### To change persistence

Open:

- `repository.go`

### To change dependency wiring

Open:

- `factory.go`

## Known Pitfalls

- Pitfall 1
- Pitfall 2
- Pitfall 3

## AI Agent Rules

- Read this file before opening source code.
- Open only the files related to the requested change.
- Update this file when responsibilities or flows change.
