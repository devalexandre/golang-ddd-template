# LLM Context - cmd

## Responsibility

This folder contains the application entrypoint.

It is responsible for starting the application and wiring dependencies.

## Allowed Content

- `main.go`
- application bootstrap
- dependency wiring
- configuration loading
- server startup

## Not Allowed

- Business rules
- SQL queries
- Domain validation
- Large application logic

## AI Agent Rules

- Keep `main.go` small.
- Move business behavior to domain or application services.
- Move infrastructure setup to dedicated packages when it grows.
- Do not implement domain logic in this folder.
