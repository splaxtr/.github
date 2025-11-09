<div align="center">
  <h1>splaxtr/.github</h1>
  <p>Centralized workflows, issue/PR templates, and community health files shared across all repositories.</p>
</div>

<p align="center">
  <a href="./README.tr.md"><img src="https://img.shields.io/badge/Language-TR-red?style=for-the-badge" alt="Türkçe"></a>
  <a href="./WORKFLOWS_GUIDE.md"><img src="https://img.shields.io/badge/Docs-Workflows%20Guide-6f42c1?style=for-the-badge&logo=githubactions" alt="Workflows Guide"></a>
  <a href="./SETUP_GUIDE.md"><img src="https://img.shields.io/badge/Docs-Setup%20Guide-0ea5e9?style=for-the-badge" alt="Setup Guide"></a>
  <a href="./FOLDER_STRUCTURE.md"><img src="https://img.shields.io/badge/Docs-Folder%20Structure-16a34a?style=for-the-badge" alt="Folder Structure"></a>
  <a href="./CONTRIBUTING.md"><img src="https://img.shields.io/badge/Docs-Contributing-f97316?style=for-the-badge" alt="Contributing"></a>
  <a href="./SECURITY.md"><img src="https://img.shields.io/badge/Docs-Security%20Policy-ef4444?style=for-the-badge" alt="Security Policy"></a>
</p>

## 🇬🇧 English · [🇹🇷 Türkçe](README.tr.md)

### 📚 Table of Contents
1. [Overview](#overview)
2. [Repository Contents](#repository-contents)
3. [Reusable Workflows](#reusable-workflows)
4. [Documentation](#documentation)
5. [Quick Start](#quick-start)
6. [Validation & Sample Projects](#validation--sample-projects)
7. [Support & Contact](#support--contact)

---

## Overview
This repository houses every shared GitHub asset used across the splaxtr organization. It keeps workflows, issue/PR templates, policies, and helper scripts in one place so downstream projects can stay consistent without duplicating configuration. The structure also includes lightweight sample applications (Node, Python, Go, Flutter, Docker) that power automated validation of reusable workflows.

## Repository Contents
| Category | What you get |
| --- | --- |
| 🔄 Workflows | Reusable CI/CD pipelines for Node, Bun, Python, Go, Flutter, Docker builds, deployments, metrics, security scans, and an aggregate validator. |
| 📋 Templates | GitHub Issue forms (bug, feature, question) plus a streamlined PR template and CODEOWNERS-aware reviewer guidance. |
| 📜 Policies | CONTRIBUTING, SECURITY, SUPPORT, CODE_OF_CONDUCT, FUNDING, and language-specific variants to keep every repo compliant. |
| 🏗️ Sample projects | `frontend`, `backend`, `mobile`, `bun-app`, `bots/*`, and `docker` skeletons that act as fixtures for CI validation. |
| 🧰 Tooling | Dependabot, label definitions, release drafter config, workflow configs, and helper scripts for bootstrap automation. |

## Reusable Workflows
| Workflow | Description |
| --- | --- |
| [reusable-node-ci](./workflows/reusable-node-ci.yml) | Matrixed lint/test/build pipeline for npm, yarn, or pnpm projects. |
| [reusable-bun-ci](./workflows/reusable-bun-ci.yml) | Bun-specific linting, coverage, build, and optional E2E stage. |
| [reusable-python-ci](./workflows/reusable-python-ci.yml) | Pip install, Ruff/Black checks, pytest coverage, and artifact uploads. |
| [reusable-go-ci](./workflows/reusable-go-ci.yml) | GolangCI-Lint, race-aware tests, binary builds, and optional Docker image. |
| [reusable-flutter-ci](./workflows/reusable-flutter-ci.yml) | Flutter format/analyze/test plus Android/Web build matrices. |
| [reusable-docker-build](./workflows/reusable-docker-build.yml) | Cached Buildx builds with GHCR pushes and metadata tagging. |
| [reusable-deploy-production](./workflows/reusable-deploy-production.yml) | Multi-service deploy orchestration (backend, frontend, mobile, bots, DB). |
| [reusable-metrics-ci](./workflows/reusable-metrics-ci.yml) | LOC/coverage/file-count collection with SVG summary artifacts. |
| [reusable-monorepo-ci](./workflows/reusable-monorepo-ci.yml) | Fan-out pipeline that runs per-stack CI inside monorepos. |
| [reusable-security](./workflows/reusable-security.yml) | npm audit + Trivy filesystem scans + CodeQL analysis. |
| [ci-lint](./workflows/ci-lint.yml) | Composite workflow that chains the reusable stacks for a full validation run. |
| [validate-reusables](./workflows/validate-reusables.yml) | Repository-only job that exercises every reusable workflow using the sample projects. |

See **[WORKFLOWS_GUIDE.md](./WORKFLOWS_GUIDE.md)** for input details, matrices, and troubleshooting notes.

## Documentation
- [WORKFLOWS_GUIDE.md](./WORKFLOWS_GUIDE.md) – Comprehensive reference for each workflow, inputs, outputs, and usage examples.
- [FOLDER_STRUCTURE.md](./FOLDER_STRUCTURE.md) – Full directory tree with explanations for each file/folder, including templates and validation assets.
- [SETUP_GUIDE.md](./SETUP_GUIDE.md) – Step-by-step instructions for creating/updating an org-level `.github` repo, customizing contacts, and wiring workflows.
- [CONTRIBUTING.md](./CONTRIBUTING.md) – Branching model, commit conventions, PR expectations, and label usage.
- [SECURITY.md](./SECURITY.md) – Disclosure policy and the `security@org` contact address.
- Bonus references: [labels-README.md](./labels-README.md), [SUPPORT.md](./SUPPORT.md), [SUPPORT.tr.md](./SUPPORT.tr.md).

## Quick Start
1. **Reference the workflows:** add a workflow in your project that reuses one of the templates, e.g.
   ```yaml
   name: CI
   on: [push, pull_request]
   jobs:
     build:
       uses: splaxtr/.github/.github/workflows/reusable-node-ci.yml@main
       with:
         node_version: '20.x'
         working_directory: '.'
   ```
2. **Leverage templates:** copy issue/PR templates automatically by keeping your repository public under the same organization—GitHub falls back to these defaults when local overrides are absent.
3. **Follow the guides:** update contact emails, policies, and docs as described in SETUP_GUIDE before rolling out to production repositories.

## Validation & Sample Projects
- The `validate-reusables.yml` workflow runs on push/dispatch and depends on the sample folders (`frontend`, `backend`, `mobile`, `bun-app`, `bots/node`, `bots/python`, `docker`). Keep them lightweight but functional so changes to workflows get exercised automatically.
- To test locally, run the same workflows via `act` or trigger `workflow_dispatch` runs in a fork before merging large changes.

## Support & Contact
- Questions: open a discussion or file a “Question” issue using the provided template.
- Bugs/features: use the corresponding issue form with as much detail as possible.
- Security: report privately to **security@org** as outlined in [SECURITY.md](./SECURITY.md).
- General support: [SUPPORT.md](./SUPPORT.md) lists email and response expectations.

Thanks for keeping the platform consistent! 🚀
