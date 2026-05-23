# AI Agent Policy

This repository is optimized for AI-assisted development.

Every relevant folder should contain an `llm.md` file.

AI agents must use these files as semantic maps before reading source code.

## Required Workflow

1. Read the closest `llm.md`.
2. Identify the files related to the task.
3. Open only the necessary source files.
4. Apply the change.
5. Update the related `llm.md` if the behavior, structure, or responsibility changed.
6. Run the relevant tests.

## Goal

Reduce unnecessary context usage and prevent architectural violations.
