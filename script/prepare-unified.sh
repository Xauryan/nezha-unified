#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
export GOTOOLCHAIN=go1.26.6
mkdir -p cmd/dashboard/docs cmd/dashboard/admin-dist cmd/dashboard/user-dist
# Swagger must be importable while its generator examines the controller.
if [[ ! -f cmd/dashboard/docs/docs.go ]]; then
  printf 'package docs\n' > cmd/dashboard/docs/docs.go
fi
# These files satisfy go:embed during unit tests only. make build replaces them
# with a locally built admin app and checksum-verified public themes.
for theme in admin user; do
  if [[ ! -f "cmd/dashboard/${theme}-dist/index.html" ]]; then
    printf 'test build input\n' > "cmd/dashboard/${theme}-dist/test-placeholder.txt"
  fi
done
go run github.com/swaggo/swag/cmd/swag@v1.16.6 init --pd -d cmd/dashboard -g main.go -o cmd/dashboard/docs
