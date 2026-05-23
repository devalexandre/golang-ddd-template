# LLM Context - internal/domain

## Responsibility

This folder contains the core business logic of the application.

It should represent the business rules independently from frameworks, databases, queues, HTTP clients, or external services.

## Allowed Content

- Entities
- Value objects
- Domain services
- Repository interfaces
- Business rules
- Domain errors
- Domain validations

## Not Allowed

- SQL queries
- Database connections
- HTTP clients
- Environment variable loading
- Framework-specific code
- Infrastructure logging
- External API implementation details

## When To Modify This Folder

Modify this folder when a business rule changes.

Examples:

- changing validation rules
- adding a new domain behavior
- changing entity invariants
- changing repository contracts
- adding domain-specific errors

## When Not To Modify This Folder

Do not modify this folder when the change is only related to:

- database implementation
- external API calls
- cache
- queues
- storage
- HTTP routing
- framework configuration

For those cases, check `internal/infra`.

## AI Agent Rules

- Keep domain code independent from infrastructure.
- Prefer interfaces over concrete infrastructure dependencies.
- Do not import packages from `internal/infra` inside domain packages.
- If a domain contract changes, check all implementations.
