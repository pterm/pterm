#!/usr/bin/env bash

# Compiles and vets every example in _examples.
#
# "go build ./_examples/..." does not work for this: the go tool skips
# directories that start with an underscore when it expands "...", matches
# nothing and still exits 0. So the example packages are listed explicitly.

set -Eeuo pipefail

source "$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)/animation-common.sh"

cd "$REPO_ROOT"

mapfile -t packages < <(
	find _examples -name main.go -not -path '*/testdata/*' -print0 |
		xargs -0 -n 1 dirname |
		sort -u |
		sed 's|^|./|'
)

((${#packages[@]} > 0)) || die "No examples found under _examples"

group "Build ${#packages[@]} examples"
go build "${packages[@]}"
end_group

if [[ "${SKIP_VET:-0}" != "1" ]]; then
	group "Vet ${#packages[@]} examples"
	go vet "${packages[@]}"
	end_group
fi

success "All ${#packages[@]} examples build"
