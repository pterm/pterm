#!/usr/bin/env bash

set -Eeuo pipefail

source "$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)/animation-common.sh"

usage() {
	cat <<EOF
Usage: scripts/animate.sh [--dry-run] [--example <section/example>]

Options:
  --dry-run              Record examples locally, write golden files, and skip publishing/docs updates.
  --example <path>       Record one example from _examples and skip publishing/docs updates.
  -h, --help             Show this help text.
EOF
}

default_jobs() {
	if [[ -n "${ANIMATE_JOBS:-}" ]]; then
		printf '%s\n' "$ANIMATE_JOBS"
	elif command -v nproc >/dev/null 2>&1; then
		nproc
	else
		printf '4\n'
	fi
}

find_examples() {
	local example="$1"

	if [[ -n "$example" ]]; then
		printf '%s\0' "$REPO_ROOT/_examples/$example/main.go"
		return
	fi

	find "$REPO_ROOT/_examples" -mindepth 3 -maxdepth 3 -name main.go -print0 | sort -z
}

record_examples() {
	local workdir="$1"
	local manifest="$2"
	local max_jobs="$3"
	local selected_example="$4"
	local manifest_dir="$workdir/manifests"
	local logs_dir="$workdir/logs"
	local status_dir="$workdir/status"
	local active=0
	local total=0
	local queued=0
	local example_dir rel key log_file status_file
	local -a examples=()

	mkdir -p "$manifest_dir" "$logs_dir" "$status_dir" "$REPO_ROOT/testdata/raw"

	if [[ -n "$selected_example" && ! -f "$REPO_ROOT/_examples/$selected_example/main.go" ]]; then
		die "Example not found: _examples/$selected_example"
	fi

	mapfile -d '' examples < <(find_examples "$selected_example")
	total="${#examples[@]}"
	(( total > 0 )) || die "No examples found under _examples"

	info "Recording $total examples with $max_jobs concurrent jobs"

	for main_file in "${examples[@]}"; do
		example_dir="$(dirname "$main_file")"
		rel="${example_dir#"$REPO_ROOT/_examples/"}"
		key="$(example_key "$rel")"
		log_file="$logs_dir/${key}.log"
		status_file="$status_dir/${key}.status"
		queued=$((queued + 1))

		printf '%s[%03d/%03d]%s %s\n' "$MAGENTA" "$queued" "$total" "$RESET" "$rel"

		(
			if "$REPO_ROOT/scripts/record-example.sh" "$example_dir" "$manifest_dir" >"$log_file" 2>&1; then
				printf '%s\tok\t%s\n' "$rel" "$log_file" >"$status_file"
			else
				printf '%s\tfail\t%s\n' "$rel" "$log_file" >"$status_file"
			fi
		) &

		active=$((active + 1))
		if (( active >= max_jobs )); then
			wait -n
			active=$((active - 1))
		fi
	done

	while (( active > 0 )); do
		wait -n
		active=$((active - 1))
	done

	local failures=0
	local successes=0
	local statuses=0
	local status rel_state result log
	for status in "$status_dir"/*.status; do
		[[ -e "$status" ]] || continue
		statuses=$((statuses + 1))
		IFS=$'\t' read -r rel_state result log <"$status"
		if [[ "$result" == "fail" ]]; then
			failures=$((failures + 1))
			print_file_group "Failed example: $rel_state" "$log"
		else
			successes=$((successes + 1))
		fi
	done

	if (( statuses != total )); then
		failures=$((failures + total - statuses))
		error "Only $statuses of $total examples reported a status"
	fi

	if (( failures > 0 )); then
		die "$failures example recording(s) failed"
	fi

	(( successes == total )) || die "Only $successes of $total examples were recorded"
	cat "$manifest_dir"/*.tsv | sort >"$manifest"
}

main() {
	local jobs workdir manifest
	local dry_run=0
	local selected_example=""

	while (($#)); do
		case "$1" in
			--dry-run)
				dry_run=1
				export ANIMATE_PUBLISH=0
				export ANIMATE_UPDATE_UNITTESTCOUNT=0
				;;
			--example)
				shift
				[[ $# -gt 0 && -n "$1" ]] || die "--example requires a path like logger/default"
				selected_example="${1#_examples/}"
				selected_example="${selected_example%/}"
				selected_example="${selected_example%/main.go}"
				export ANIMATE_PUBLISH=0
				export ANIMATE_UPDATE_UNITTESTCOUNT=0
				;;
			-h | --help)
				usage
				exit 0
				;;
			*)
				die "Unknown argument: $1"
				;;
		esac
		shift
	done

	group "Animation environment"
	require_cmd go
	require_cmd docker
	require_cmd python3
	require_cmd realpath

	jobs="$(default_jobs)"
	[[ "$jobs" =~ ^[0-9]+$ && "$jobs" -gt 0 ]] || die "ANIMATE_JOBS must be a positive number"

	info "Repository: $REPO_ROOT"
	info "Concurrent jobs: $jobs"
	info "Publish GIFs: ${ANIMATE_PUBLISH:-1}"
	info "VHS Docker image: ${VHS_DOCKER_IMAGE:-ghcr.io/charmbracelet/vhs:latest}"
	info "Dry run: $dry_run"
	if [[ -n "$selected_example" ]]; then
		info "Selected example: $selected_example"
	fi
	end_group

	workdir="$(mktemp -d)"
	trap "rm -rf '$workdir'" EXIT
	manifest="$workdir/published.tsv"

	group "Record and publish examples"
	record_examples "$workdir" "$manifest" "$jobs" "$selected_example"
	end_group

	if (( dry_run == 1 )) || [[ -n "$selected_example" ]]; then
		success "Local animation run finished; GIFs and golden files were generated"
		return
	fi

	"$REPO_ROOT/scripts/render-animation-docs.sh" "$manifest"

	if [[ "${ANIMATE_UPDATE_UNITTESTCOUNT:-1}" == "1" ]]; then
		"$REPO_ROOT/scripts/update-unittestcount.sh"
	else
		warn "Skipping unittest count update because ANIMATE_UPDATE_UNITTESTCOUNT=$ANIMATE_UPDATE_UNITTESTCOUNT"
	fi

	success "Animation pipeline finished"
}

main "$@"