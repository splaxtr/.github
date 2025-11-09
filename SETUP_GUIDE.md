# 🛠️ Org-level `.github` Setup Guide

Use this playbook to create or update the organization-wide `.github` repository so that every project inherits the same workflows, templates, and policies.

## 1. Prerequisites
- GitHub organization admin permissions.
- GitHub CLI (`gh`) or browser access.
- Familiarity with Git workflows and GitHub Actions.

## 2. Create (or sync) the `.github` repository
```bash
# Create the repo if it does not exist
gh repo create <org>/.github --public --description "Org-wide workflows and health files"

# Clone locally
git clone https://github.com/<org>/.github.git
cd .github

git checkout main
```
If the repo already exists, pull the latest changes:
```bash
git fetch origin
git pull origin main
```

## 3. Customize contact and policy files
1. Update contact emails in `ISSUE_TEMPLATE/config.yml`, `SUPPORT.md`, `CONTRIBUTING.md`, and `SECURITY.md` (e.g., replace `support@org`, `security@org`).
2. Translate or localize text if your organization needs additional languages—keep both English and Turkish files in sync.
3. Review funding links (`FUNDING.yml`) and CODEOWNERS (if present) to make sure reviewers are accurate.

## 4. Link reusable workflows to downstream repos
For each project that should consume the shared workflows:
1. Remove or simplify repo-specific workflow logic if you intend to reuse the central ones.
2. Add a workflow referencing the reusable template. Example (`.github/workflows/ci.yml` in the consuming repo):
   ```yaml
   name: CI
   on: [push, pull_request]
   jobs:
     test:
       uses: <org>/.github/.github/workflows/reusable-node-ci.yml@main
       with:
         node_version: '20.x'
         working_directory: '.'
   ```
3. For advanced scenarios (deploy, metrics, security) inspect [WORKFLOWS_GUIDE.md](./WORKFLOWS_GUIDE.md) to supply the correct inputs/secrets.

## 5. Refresh templates & docs
- Ensure issue forms (`ISSUE_TEMPLATE/*.yml`) reflect your preferred severity levels, labels, and assignees.
- Update `PULL_REQUEST_TEMPLATE.md` if reviewer or checklist defaults change.
- Keep `README.md` and `README.tr.md` aligned—update both whenever technical details change.
- Regenerate `FOLDER_STRUCTURE.md` after adding directories/files.

## 6. Validate reusable workflows
1. The repo ships sample apps (`frontend`, `backend`, `mobile`, `bun-app`, `bots/*`, `docker`) that power the validation workflow. Keep them minimal but working.
2. Run the comprehensive validation workflow before merging major changes:
   ```bash
   gh workflow run validate-reusables.yml
   # or trigger via GitHub UI (Actions → Validate Reusable Workflows → Run workflow)
   ```
3. Inspect the `reusable-validation-report` artifact to confirm every job succeeded. Fix failures before rolling out updates to production repositories.

## 7. Roll out changes
1. Commit updates with conventional prefixes (`feat:`, `fix:`, etc.) and open a PR referencing the relevant issue.
2. Ensure `validate-reusables` and other CI workflows pass.
3. Merge via PR, then notify maintainers of downstream repos to pull the latest `.github` changes if needed.

## 8. Keep docs handy
Add quick links to the main README badge row so teammates can jump to:
- [WORKFLOWS_GUIDE.md](./WORKFLOWS_GUIDE.md)
- [FOLDER_STRUCTURE.md](./FOLDER_STRUCTURE.md)
- [SETUP_GUIDE.md](./SETUP_GUIDE.md)
- [CONTRIBUTING.md](./CONTRIBUTING.md)
- [SECURITY.md](./SECURITY.md)

Following these steps ensures every repository in the organization benefits from the same automation, templates, and policies with minimal duplication.
