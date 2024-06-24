#!/usr/bin/env bash

set -euo pipefail

REPO_ROOT=$(git rev-parse --show-toplevel)
readonly REPO_ROOT

if ! command "fhenix" >/dev/null 2>&1; then
  go install github.com/friendly-fhir/fhenix@latest
fi

artifact_dir="${REPO_ROOT}/dist"
readonly artifact_dir

fhenix run                                                                     \
  .fhenix/r4/config.yaml                                                       \
  --output "${artifact_dir}"                                                   \
  --fhirig-cache "${REPO_ROOT}/.fhenix/cache"

rm -r "${REPO_ROOT}/r4"
mv "${artifact_dir}/r4" "${REPO_ROOT}/r4"
go fmt ./r4/...
rm -r "${artifact_dir}"
