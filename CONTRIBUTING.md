# Contributing Guide

Thanks for helping improve this repository! This document explains how we work so you can get productive quickly.

## Branching strategy
- `main` is always deployable. Direct pushes are disabled; use pull requests (PRs).
- Create topic branches from `main` using the format `<type>/<short-description>` (e.g. `feat/runtime-flags`, `fix/login-race`).
- Rebase onto `main` before opening a PR (or before merging) to keep history clean.

## Commit naming
Prefix every commit with one of the following types so history stays searchable:
- `feat:` – new capability or API
- `fix:` – bug fix
- `chore:` – tooling, dependencies, build, infra
- `docs:` – documentation-only changes
- `test:` – tests only (unit/integration/e2e)

Example: `feat: add build artifact caching`.

## Pull request flow
1. Create/choose an issue describing the work. Link it in your PR (`Closes #123`).
2. Ensure lint/tests/workflows succeed locally (see `.github/workflows`).
3. Fill out the PR template sections in English.
4. Add labels for type/component/priority so triage stays consistent.
5. Assign reviewers suggested via CODEOWNERS (GitHub does this automatically; double-check the list).
6. Respond to review feedback with follow-up commits or comments; squash/merge via GitHub once approved.

## Labels
We rely on labels for automation and release notes:
- `type:*` – categorize the change (`type: bug`, `type: feature`, `type: docs`, etc.).
- `component:*` – identify the area affected (`component: backend`, `component: mobile`, etc.).
- `priority:*` – communicate urgency (`priority: critical`, `priority: low`).
- `status:*` – optional workflow states (`status: needs-info`, `status: ready-for-review`).

Please apply the appropriate labels when opening a PR or filing an issue so maintainers can route it quickly.

## Code style & tests
- Follow the linters/formatters configured in the repo (see language-specific READMEs when available).
- Add or update tests whenever functionality changes.
- Keep public API changes backward compatible unless the PR clearly documents a breaking change and migration notes.

## Reporting issues
Use the issue templates under `.github/ISSUE_TEMPLATE` to report bugs, request features, or ask questions. Include repro steps, environment details, and acceptance criteria when relevant.

Thank you for contributing! 🚀 For any questions or support needs, email **ahmetsplaxtr@gmail.com**.
