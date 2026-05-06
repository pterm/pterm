#!/usr/bin/env bash

set -Eeuo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd -- "$SCRIPT_DIR/.." && pwd)"

if [[ -t 1 && -z "${NO_COLOR:-}" ]] || [[ "${GITHUB_ACTIONS:-}" == "true" && -z "${NO_COLOR:-}" ]]; then
	RESET=$'\033[0m'
	BOLD=$'\033[1m'
	DIM=$'\033[2m'
	RED=$'\033[31m'
	GREEN=$'\033[32m'
	YELLOW=$'\033[33m'
	BLUE=$'\033[34m'
	MAGENTA=$'\033[35m'
	CYAN=$'\033[36m'
else
	RESET=""
	BOLD=""
	DIM=""
	RED=""
	GREEN=""
	YELLOW=""
	BLUE=""
	MAGENTA=""
	CYAN=""
fi

group() {
	if [[ "${GITHUB_ACTIONS:-}" == "true" ]]; then
		printf '::group::%s\n' "$*"
	else
		printf '\n%s%s%s\n' "$BOLD" "$*" "$RESET"
	fi
}

end_group() {
	if [[ "${GITHUB_ACTIONS:-}" == "true" ]]; then
		printf '::endgroup::\n'
	fi
}

info() {
	printf '%s==>%s %s\n' "$CYAN" "$RESET" "$*"
}

success() {
	printf '%sOK%s %s\n' "$GREEN" "$RESET" "$*"
}

warn() {
	if [[ "${GITHUB_ACTIONS:-}" == "true" ]]; then
		printf '::warning::%s\n' "$*"
	fi
	printf '%sWARN%s %s\n' "$YELLOW" "$RESET" "$*"
}

error() {
	if [[ "${GITHUB_ACTIONS:-}" == "true" ]]; then
		printf '::error::%s\n' "$*"
	fi
	printf '%sERROR%s %s\n' "$RED" "$RESET" "$*" >&2
}

die() {
	error "$*"
	exit 1
}

require_cmd() {
	local cmd="$1"
	command -v "$cmd" >/dev/null 2>&1 || die "Missing required command: $cmd"
}

example_key() {
	local rel="$1"
	printf '%s' "${rel//\//_}"
}

replace_between() {
	local marker="$1"
	local file="$2"
	local insert_file="$3"

	python3 - "$marker" "$file" "$insert_file" <<'PY' || die "Could not replace marker in file"
import sys
from pathlib import Path

marker, file_name, insert_name = sys.argv[1:]
start = f"<!-- {marker}:start -->"
end = f"<!-- {marker}:end -->"

path = Path(file_name)
text = path.read_text()
insert = Path(insert_name).read_text()

start_index = text.find(start)
if start_index == -1:
    raise SystemExit(f"missing marker: {start}")

content_start = start_index + len(start)
end_index = text.find(end, content_start)
if end_index == -1:
    raise SystemExit(f"missing marker: {end}")

old = text[content_start:end_index]
if old.startswith("\n") and not insert.startswith("\n"):
    insert = "\n" + insert
if old.endswith("\n") and not insert.endswith("\n"):
    insert += "\n"

path.write_text(text[:content_start] + insert + text[end_index:])
PY
}

print_file_group() {
	local title="$1"
	local file="$2"

	group "$title"
	sed 's/^/  /' "$file" || true
	end_group
}
