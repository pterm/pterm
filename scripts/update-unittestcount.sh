#!/usr/bin/env bash

set -Eeuo pipefail

source "$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)/animation-common.sh"

tmp="$(mktemp)"
trap 'rm -f "$tmp"' EXIT

group "Update unittest count"
info "Counting tests with go test -json"

set +e
go test -json -p 1 ./... >"$tmp"
test_status=$?
set -e

if (( test_status != 0 )); then
	warn "go test failed while counting tests; keeping the partial count for now"
fi

count="$(awk '/"Action":"run"/ {count++} END {print count + 0}' "$tmp")"
info "Found $count unit test runs"

badge_file="$(mktemp)"
inline_file="$(mktemp)"
trap 'rm -f "$tmp" "$badge_file" "$inline_file"' EXIT

printf '<img src="https://img.shields.io/badge/Unit_Tests-%s-magenta?style=flat-square" alt="Unit Tests">' "$count" >"$badge_file"
printf '`%s`' "$count" >"$inline_file"

replace_between "unittestcount" "$REPO_ROOT/README.md" "$badge_file"
replace_between "unittestcount2" "$REPO_ROOT/README.md" "$inline_file"
end_group

success "Updated README unit test count to $count"
