# LLM Context - internal/infra

## Responsibility

This folder contains technical implementations used by the application.

Infrastructure code should implement contracts defined by the domain or application layers.

## Allowed Content

- Database connections
- Repository implementations
- HTTP clients
- External service adapters
- Cache implementations
- Queue integrations
- Storage integrations
- Logging adapters
- Framework-specific configuration

## Not Allowed

- Core business rules
- Domain decisions
- Entity invariants
- Business validation that belongs to the domain layer

## Common Tasks

Modify this folder when changing:

- SQL queries
- database drivers
- external API clients
- cache behavior
- queue publishing/consuming
- infrastructure configuration

## AI Agent Rules

- Before changing infrastructure for a domain module, read the related domain `llm.md`.
- Infrastructure should depend on domain contracts, not the opposite.
- Keep external service details isolated in this layer.
- Avoid leaking infrastructure types into the domain layer.
