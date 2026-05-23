# LLM Context - Project Root

## Project Purpose

This project is a Go DDD template designed to organize applications into clear layers: entrypoint, domain, infrastructure, and shared helpers.

## How AI Agents Should Navigate This Project

Before opening `.go` files, always read the nearest `llm.md` file.

Recommended navigation order:

1. Read `/llm.md`
2. Read the `llm.md` inside the target top-level folder
3. Read the `llm.md` inside the specific module folder
4. Open source code only when the markdown file indicates it is necessary

## Main Folders

- `cmd/`: application entrypoint.
- `internal/domain/`: business rules, entities, domain services, and repository contracts.
- `internal/infra/`: database, external services, technical adapters, and infrastructure implementations.
- `internal/helpers/`: shared utility functions.

## AI Agent Rules

- Do not modify a folder before reading its `llm.md`.
- Do not place business logic inside infrastructure code.
- Do not place SQL queries inside domain services.
- Do not change public contracts without checking dependent code.
- If behavior changes, update the related `llm.md`.
- If a new module is created, create an `llm.md` for it.
