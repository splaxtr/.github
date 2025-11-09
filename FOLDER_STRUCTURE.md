# 🏗️ Repository Structure

This document explains the layout of the `.github` repository so you know where to add or update shared assets.

## Directory Tree
```
.github/
├── ISSUE_TEMPLATE/
│   ├── bug_report.yml
│   ├── feature_request.yml
│   ├── question.yml
│   └── config.yml
├── PULL_REQUEST_TEMPLATE.md
├── WORKFLOWS_GUIDE.md
├── SETUP_GUIDE.md
├── FOLDER_STRUCTURE.md
├── CONTRIBUTING.md
├── SECURITY.md
├── SUPPORT.md / SUPPORT.tr.md
├── README.md / README.tr.md
├── workflows/
│   ├── reusable-*.yml
│   ├── ci-lint.yml
│   ├── labels.yml
│   └── validate-reusables.yml
├── labels.yml & labels-README.md
├── release-drafter.yml & release-drafter-README.md
├── configs/ · docs/ · scripts/
├── dependabot.yml
├── FUNDING.yml
├── CODE_OF_CONDUCT.md
├── LICENSE
├── backend/ · frontend/ · bun-app/ · mobile/ · docker/
└── bots/
    ├── node/
    └── python/
```

## Key directories & files

### ISSUE_TEMPLATE/
- `bug_report.yml`: GitHub form that captures severity, reproduction steps, and environment data for defects.
- `feature_request.yml`: Collects use cases, proposals, and acceptance criteria for roadmap requests.
- `question.yml`: Lightweight template for how-to questions and pointers to Discussions.
- `config.yml`: Disables blank issues and lists contact links (docs, discussions, support email).

### PULL_REQUEST_TEMPLATE.md
Single PR template with Summary, Type of Change, Linked Issues, Implementation Notes, Screenshots, and Checklist sections. Reminds authors to verify CODEOWNERS-suggested reviewers.

### Workflows
Located under `workflows/`:
- `reusable-*.yml`: Language/runtime specific CI pipelines (Node, Bun, Python, Go, Flutter), Docker builds, deployments, metrics, security, and monorepo orchestration.
- `ci-lint.yml`: Composite workflow that chains the reusable jobs for an end-to-end check.
- `labels.yml`: Synchronizes organization-wide labels.
- `validate-reusables.yml`: Runs every reusable workflow using the sample projects shipped in this repo.

### Documentation & policies
- `README.md` / `README.tr.md`: Bilingual overview, table of contents, and quick start.
- `WORKFLOWS_GUIDE.md`: Deep dive into inputs/outputs per workflow.
- `SETUP_GUIDE.md`: Org-level bootstrap/update steps, including validation instructions.
- `CONTRIBUTING.md`: Branching rules, commit tags, PR expectations, and label usage.
- `SECURITY.md`: Disclosure policy and `ahmetsplaxtr@gmail.com` contact.
- `SUPPORT.md` (+ Turkish version): Communication channels and SLA expectations.
- `FOLDER_STRUCTURE.md`: This document.

### Release & automation configs
- `release-drafter.yml` + `release-drafter-README.md`: Automated changelog generation.
- `labels.yml` + `labels-README.md`: Standard label definitions plus usage guidance.
- `dependabot.yml`: Dependency update schedule.
- `FUNDING.yml`: Sponsor links displayed on GitHub.

### Sample applications (for workflow validation)
- `frontend/`: Minimal Node/React-style app.
- `backend/`: Go module with tests and Dockerfile.
- `mobile/`: Flutter app skeleton.
- `bun-app/`: Bun + TypeScript example with tests.
- `bots/node` and `bots/python`: Representative automation scripts for each runtime.
- `docker/`: Simple Node service used to build/push sample containers.

These fixtures are intentionally lightweight and only exist to exercise reusable workflows locally and in CI.

### Utility folders
- `configs/`, `docs/`, `scripts/`: Helper assets and automation scripts used when bootstrapping repositories.
- `CODE_OF_CONDUCT.md`, `LICENSE`: Standard community health defaults inherited by org repositories.

Use this structure to decide where a new template, guide, or workflow should live before opening a PR.
