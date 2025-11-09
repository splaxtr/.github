# Workflows Guide

This guide summarizes every workflow under `.github/workflows`, including purpose, notable inputs, outputs, usage examples, and quick troubleshooting tips.

## workflows/ci-lint.yml
**Purpose:** Composite orchestrator that chains the reusable Node, Bun, Python, Go, Flutter, and Security pipelines to run full-stack checks with a single `uses` block.

**Key inputs:** `node_working_directory`, `python_working_directory`, `go_working_directory`, `flutter_working_directory`, `bun_enabled`, `bun_working_directory`, `security_enabled`, `docker_context`.

**Outputs / artifacts:** Inherits the artifacts uploaded by each chained workflow, and writes a step summary detailing job results.

**Usage:**
```yaml
jobs:
  ci:
    uses: splaxtr/.github/.github/workflows/ci-lint.yml@main
    with:
      node_working_directory: frontend
      python_working_directory: bots/python
      go_working_directory: backend
      flutter_working_directory: mobile
```

**Troubleshooting:** Ensure each referenced directory exists; pass `bun_enabled: false` or `security_enabled: false` if those stacks are absent to avoid job failures.

## workflows/labels.yml
**Purpose:** Keeps repository labels in sync with `.github/labels.yml`, validates the YAML, and optionally reports basic metrics.

**Key inputs:** `labels_file`, `dry_run`, `collect_stats`.

**Outputs / artifacts:** Writes sync results and validation info to the job summary.

**Usage:**
```yaml
jobs:
  labels:
    uses: splaxtr/.github/.github/workflows/labels.yml@main
    with:
      labels_file: .github/labels.yml
      dry_run: false
```

**Troubleshooting:** Requires `contents`, `issues`, and `pull-requests` write permissions; make sure the labels file is valid YAML and follows the expected schema.

## workflows/release-drafter.yml
**Purpose:** Runs Release Drafter with org-standard config to keep draft releases up to date.

**Key inputs:** `config_path` (path to the release drafter config).

**Outputs / artifacts:** Summary entry documenting which config was applied.

**Usage:**
```yaml
jobs:
  release_notes:
    uses: splaxtr/.github/.github/workflows/release-drafter.yml@main
    with:
      config_path: .github/release-drafter.yml
```

**Troubleshooting:** Needs `contents: write` and `pull-requests: read`; ensure the referenced config exists or the workflow will fail early.

## workflows/reusable-bun-ci.yml
**Purpose:** Lints, formats, type-checks, tests (with coverage), builds, and optionally runs E2E checks for Bun projects.

**Key inputs:** `bun_version`, `working_directory`, `lint_command`, `format_command`, `typecheck_command`, `test_command`, `build_command`, `run_e2e_tests`.

**Outputs / artifacts:** Coverage bundle (`bun-coverage-*`) and build artifacts (`bun-build-*`) uploaded via `actions/upload-artifact`.

**Usage:**
```yaml
jobs:
  bun:
    uses: splaxtr/.github/.github/workflows/reusable-bun-ci.yml@main
    with:
      bun_version: '1.0.0'
      working_directory: apps/bun-service
```

**Troubleshooting:** Provide a pre-generated `bun.lockb`; the workflow runs `bun install --frozen-lockfile`, which fails if the lockfile is missing.

## workflows/reusable-deploy-production.yml
**Purpose:** Orchestrates deploy steps for backend (Go), frontend (Node), mobile (Flutter), and bots, including optional smoke tests and DB migrations.

**Key inputs:** `environment`, `changed_services`, `force_deploy`, `node_version`, `python_version`, `go_version`, `flutter_version`, `flutter_channel`, `docker_context`, `run_smoke_tests`.

**Outputs / artifacts:** Step summary containing per-service deploy status; reuses artifacts created within each internal job (e.g., mobile build uploads).

**Usage:**
```yaml
jobs:
  deploy:
    uses: splaxtr/.github/.github/workflows/reusable-deploy-production.yml@main
    with:
      environment: production
      changed_services: backend,frontend
      docker_context: backend
```

**Troubleshooting:** Provide the necessary secrets (e.g., `KUBECONFIG`, deploy tokens, Vercel credentials); without them the workflow will log a warning and skip real deployments.

## workflows/reusable-docker-build.yml
**Purpose:** Builds (and optionally pushes) Docker images to GHCR with metadata labels and GHA cache support.

**Key inputs:** `docker_context`, `dockerfile`, `platforms`, `push_image`, `build_args`.

**Outputs / artifacts:** Publishes multi-tag images to GHCR and emits a summary with the image name and tags.

**Usage:**
```yaml
jobs:
  docker:
    uses: splaxtr/.github/.github/workflows/reusable-docker-build.yml@main
    with:
      docker_context: services/api
      dockerfile: Dockerfile
      push_image: true
```

**Troubleshooting:** Requires `packages: write` permission to push; when `push_image: false`, the step still builds locally but skips registry auth.

## workflows/reusable-flutter-ci.yml
**Purpose:** Runs Flutter formatting, analysis, unit/widget tests (with coverage), and optional multi-platform builds.

**Key inputs:** `flutter_version`, `flutter_channel`, `working_directory`, `build_targets`, `run_integration_tests`.

**Outputs / artifacts:** Coverage file (`flutter-coverage-*`) plus per-platform build artifacts (APK/Web) for enabled targets.

**Usage:**
```yaml
jobs:
  flutter:
    uses: splaxtr/.github/.github/workflows/reusable-flutter-ci.yml@main
    with:
      working_directory: apps/mobile
      build_targets: android,web
```

**Troubleshooting:** Ensure `pubspec.yaml` and platform folders exist; disable unused targets by omitting them from `build_targets` to save time.

## workflows/reusable-go-ci.yml
**Purpose:** Handles Go module linting, vet/staticcheck, testing with coverage, optional binary builds, and optional Docker builds.

**Key inputs:** `go_version`, `working_directory`, `coverage_threshold`, `build_binary`, `enable_race_detector`, `docker_context`, `dockerfile`, `build_image`.

**Outputs / artifacts:** Coverage file (`go-coverage-*`) and compiled binaries (`go-binary-*`); optional container image pushed to GHCR.

**Usage:**
```yaml
jobs:
  go:
    uses: splaxtr/.github/.github/workflows/reusable-go-ci.yml@main
    with:
      working_directory: services/backend
      coverage_threshold: 80
      build_image: true
```

**Troubleshooting:** Provide a `go.sum` for caching; staticcheck is allowed to warn without failing, but golangci-lint errors will stop the job.

## workflows/reusable-metrics-ci.yml
**Purpose:** Collects LOC, file counts, optional coverage, renders an SVG chart, and publishes a metrics summary for any supported language.

**Key inputs:** `working-directory`, `language`, `coverage-file`.

**Outputs / artifacts:** `profile/metrics.svg` artifact plus a GitHub Step Summary table with computed metrics.

**Usage:**
```yaml
jobs:
  metrics:
    uses: splaxtr/.github/.github/workflows/reusable-metrics-ci.yml@main
    with:
      working-directory: frontend
      language: node
```

**Troubleshooting:** For coverage parsing, supply a supported format (`json`, `xml`, or `lcov`); missing files default to 0% without failing the workflow.

## workflows/reusable-monorepo-ci.yml
**Purpose:** Kicks off language-specific CI jobs (Node, Bun, Python, Go, Flutter) in parallel for monorepo setups, with configurable paths per service.

**Key inputs:** `node_version`, `node_path`, `bun_version`, `bun_path`, `python_version`, `python_path`, `go_version`, `go_path`, `flutter_version`, `flutter_channel`, `flutter_path`.

**Outputs / artifacts:** Per-language coverage artifacts (e.g., `node-coverage-*`, `go-coverage-*`) and summary of job results.

**Usage:**
```yaml
jobs:
  monorepo:
    uses: splaxtr/.github/.github/workflows/reusable-monorepo-ci.yml@main
    with:
      node_path: services/web
      bun_path: services/bot
      python_path: services/py
      go_path: services/api
      flutter_path: apps/mobile
```

**Troubleshooting:** Provide valid paths for each stack or trim unused stacks by pointing them at directories containing minimal projects.

## workflows/reusable-node-ci.yml
**Purpose:** Standard Node.js pipeline that installs dependencies, tests, builds, and uploads artifacts, supporting npm/yarn/pnpm via inputs.

**Key inputs:** `node_version`, `working_directory`, `package_manager`, `lockfile`, `test_command`, `build_command`, `run_tests`, `artifact_path`.

**Outputs / artifacts:** Uploads `node-build-*` artifacts (if available) and emits a step summary with the matrix stage info.

**Usage:**
```yaml
jobs:
  node:
    uses: splaxtr/.github/.github/workflows/reusable-node-ci.yml@main
    with:
      working_directory: frontend
      package_manager: npm
```

**Troubleshooting:** Ensure the selected package manager’s lockfile exists (`npm ci`/`yarn install --frozen-lockfile`/`pnpm install --frozen-lockfile`); missing lockfiles cause install failures.

## workflows/reusable-python-ci.yml
**Purpose:** Performs pip installs, formatting (Black), linting (Ruff), pytest with coverage, and uploads artifacts for Python projects.

**Key inputs:** `python_version`, `working_directory`, `lint_command`, `format_command`, `test_command`, `coverage_file`, `run_tests`, `upload_artifacts`.

**Outputs / artifacts:** Combined artifact containing `.pytest_cache` and the requested coverage file.

**Usage:**
```yaml
jobs:
  python:
    uses: splaxtr/.github/.github/workflows/reusable-python-ci.yml@main
    with:
      working_directory: bots/python
      run_tests: true
```

**Troubleshooting:** Add the necessary `requirements*.txt` files or adjust the install command; missing files are ignored but lint/test commands still run on the tree.

## workflows/reusable-security.yml
**Purpose:** Bundles npm audit, filesystem Trivy scans, and CodeQL analysis into one reusable security pipeline.

**Key inputs:** `working-directory`, `node_version`, `trivy_target`, `codeql_languages`, `npm_audit`, `run_trivy`, `run_codeql`.

**Outputs / artifacts:** Trivy SARIF uploads and CodeQL results published to GitHub Advanced Security.

**Usage:**
```yaml
jobs:
  security:
    uses: splaxtr/.github/.github/workflows/reusable-security.yml@main
    with:
      working-directory: .
      npm_audit: true
      run_trivy: true
      run_codeql: true
```

**Troubleshooting:** Requires `security-events: write`; disable steps you do not need (e.g., set `run_trivy: false`) to reduce runtime or avoid missing prerequisites.

## workflows/validate-reusables.yml
**Purpose:** Repository-only validation workflow that runs every reusable workflow via `uses` syntax with sample inputs, capturing their artifacts and job results.

**Key triggers:** `push`, `workflow_dispatch`.

**Outputs / artifacts:** Consolidated `reusable-validation-report` artifact plus any artifacts emitted by the invoked workflows.

**Usage:**
```yaml
on:
  workflow_dispatch:

jobs:
  validate:
    uses: splaxtr/.github/.github/workflows/validate-reusables.yml@main
```

**Troubleshooting:** Designed for this repo; downstream consumers should run the individual reusable workflows directly rather than nesting this validator.
