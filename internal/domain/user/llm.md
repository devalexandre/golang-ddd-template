# LLM Context - internal/domain/user

## Responsibility

This package contains the user domain.

It defines the user entity, repository contract, and user-related business behavior.

## Important Files

| File | Purpose | Open When |
|---|---|---|
| `contracts.go` | Defines user contracts, entities, and repository types | Changing public domain contracts |
| `repository.go` | Contains user persistence behavior or repository implementation | Changing user data access |
| `service.go` | Contains user business rules | Changing user behavior |
| `factory.go` | Creates user-related dependencies | Changing dependency wiring |

## Domain Concepts

### User

Represents an application user.

Expected fields may include:

- `ID`
- `Name`

## Repository Behavior

The user repository is responsible for:

- inserting users
- updating users
- deleting users
- finding one user
- listing users

## Architecture Notes

This package should ideally define repository interfaces only.

Concrete database implementations should live in `internal/infra`.

If this package imports infrastructure code, consider refactoring it so that:

- domain defines the interface
- infrastructure implements the interface
- application wiring connects both layers

## AI Agent Rules

- Do not add SQL directly to domain services.
- Do not add database-specific logic to entities.
- If repository methods change, update all implementations.
- If user behavior changes, update this file.
