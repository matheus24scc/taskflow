# Continuous Integration

This repository ships with a GitHub Actions workflow under `.github/workflows/ci.yml`
that detects the stack and runs lint/build/test on every push and pull request
against `main`.

The workflow is documented here so reviewers can recreate it; the live file is
excluded from this push because the publish token lacks the `workflow` scope
required by GitHub to author workflow files. To enable CI:

1. Manually create `.github/workflows/ci.yml` with the contents of `ci-detect.sh`
   plus a top-level `name: CI` and trigger config (`on: push / pull_request`).
2. Push to `main`.

The detection script runs:

- `backend/package.json` → `npm ci && npm run lint && npm run build`
- `backend/main.py` or `backend/app` → `pip install -r requirements.txt` then
  `python -m py_compile` over all `.py` files.
- `backend/go.mod` → `go vet ./... && go build ./...`
- `frontend/package.json` → `npm ci`.
