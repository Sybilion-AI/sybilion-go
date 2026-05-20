#!/usr/bin/env bash
# Regenerate the Go client at api/ from vendored openapi/openapi.yaml and refresh defaults_gen.go.
# Requires Docker.
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SPEC="/local/openapi/openapi.yaml"
IMG="${OPENAPI_GENERATOR_IMAGE:-openapitools/openapi-generator-cli:v7.11.0}"
MOD="go.sybilion.dev/sybilion"
UID_LOCAL="$(id -u)"
GID_LOCAL="$(id -g)"
GEN_USER="${GEN_USER:-${UID_LOCAL}:${GID_LOCAL}}"

run_gen() {
  local generator="$1" out_dir="$2"
  shift 2
  docker run --rm --user "${GEN_USER}" -v "${ROOT}:/local" "${IMG}" generate \
    -i "${SPEC}" -g "${generator}" -o "/local/${out_dir}" "$@"
}

# Clean output so stale models from older specs cannot linger.
RESTORE_IGNORE=""
if [[ -f "${ROOT}/api/.openapi-generator-ignore" ]]; then
  RESTORE_IGNORE="$(mktemp "${TMPDIR:-/tmp}/openapi-ignore.XXXXXX")"
  cp "${ROOT}/api/.openapi-generator-ignore" "${RESTORE_IGNORE}"
fi
rm -rf "${ROOT}/api"
mkdir -p "${ROOT}/api"
if [[ -n "${RESTORE_IGNORE}" ]]; then
  mv "${RESTORE_IGNORE}" "${ROOT}/api/.openapi-generator-ignore"
fi

# packageName=sybilionapi gives the generated Go package the import alias users see.
run_gen go api \
  --additional-properties=packageName=sybilionapi,withGoMod=false,generateInterfaces=false,enumClassPrefix=true

if [[ -d "${ROOT}/api/docs" ]]; then
  find "${ROOT}/api/docs" -name '*.md' -type f -print0 | xargs -0 perl -pi -e "s|github.com/GIT_USER_ID/GIT_REPO_ID|${MOD}/api|g"
fi

rm -rf "${ROOT}/api/test" "${ROOT}/api/.travis.yml" "${ROOT}/api/git_push.sh" 2>/dev/null || true

if [[ ! -w "${ROOT}/api/client.go" ]]; then
  docker run --rm -v "${ROOT}:/local" alpine:3.19 sh -c "chown -R ${UID_LOCAL}:${GID_LOCAL} /local/api" || true
fi

(cd "${ROOT}" && go run ./cmd/embed-go-defaults)
gofmt -w "${ROOT}/defaults_gen.go"

rm -rf "${ROOT}/api/test" "${ROOT}/api/git_push.sh" "${ROOT}/api/.travis.yml" 2>/dev/null || true

echo "sdk-generate-go: complete."
