#!/usr/bin/env bash

set -Eeuo pipefail

source "$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)/animation-common.sh"

usage() {
	cat <<EOF
Usage: scripts/animate.sh [options]

Records every example in _examples as animation.gif with VHS (in Docker) and
regenerates the example READMEs and the README example section.

Options:
  --example <path>   Record only this example, e.g. table/demo. Can be repeated.
                     Docs are not regenerated.
  --dry-run          Record all examples but leave the docs alone.
  --measure-only     Build and measure examples and print the generated tapes.
                     Needs no Docker and writes no GIFs.
  --publish          Also upload the GIFs to vhs.charm.sh and link those URLs
                     instead of the committed files (ANIMATE_PUBLISH=1).
  -h, --help         Show this help text.

Environment: ANIMATE_JOBS (parallel recordings, default: CPU count) and the
sizing knobs described in scripts/README.md.
EOF
}

default_jobs() {
	if [[ -n "${ANIMATE_JOBS:-}" ]]; then
		printf '%s\n' "$ANIMATE_JOBS"
	elif command -v nproc >/dev/null 2>&1; then
		nproc
	elif command -v sysctl >/dev/null 2>&1; then
		sysctl -n hw.ncpu
	else
		printf '4\n'
	fi
}

# Prints the example directories relative to _examples, NUL separated.
find_examples() {
	local rel

	if (($# > 0)); then
		printf '%s\0' "$@"
		return
	fi

	find "$REPO_ROOT/_examples" -mindepth 3 -maxdepth 3 -name main.go -print0 |
		sort -z |
		while IFS= read -r -d '' main_file; do
			rel="$(dirname "$main_file")"
			printf '%s\0' "${rel#"$REPO_ROOT/_examples/"}"
		done
}

# Builds every example once, in parallel, before anything is recorded. A
# broken example fails here with the compiler output instead of halfway
# through a recording.
build_examples() {
	local bin_dir="$1"
	local logs_dir="$2"
	local max_jobs="$3"
	shift 3
	local cross_compile=0
	local rel key failures=0

	[[ "$(go env GOOS)" == "linux" ]] || cross_compile=1
	mkdir -p "$bin_dir" "$logs_dir"

	export bin_dir logs_dir cross_compile
	printf '%s\0' "$@" | xargs -0 -n 1 -P "$max_jobs" bash -c '
		source "$1"
		rel="$2"
		key="$(example_key "$rel")"
		{
			build_example "$REPO_ROOT/_examples/$rel" "$bin_dir/$key" &&
				if [[ "$cross_compile" == "1" ]]; then
					build_example "$REPO_ROOT/_examples/$rel" "$bin_dir/$key.linux" linux
				fi
		} >"$logs_dir/$key.build.log" 2>&1 || true
	' _ "$SCRIPT_DIR/animation-common.sh"

	for rel in "$@"; do
		key="$(example_key "$rel")"
		if [[ ! -x "$bin_dir/$key" ]]; then
			failures=$((failures + 1))
			print_file_group "Build failed: _examples/$rel" "$logs_dir/$key.build.log"
		fi
	done

	((failures == 0)) || die "$failures example(s) failed to build"
}

record_examples() {
	local workdir="$1"
	local manifest="$2"
	local max_jobs="$3"
	shift 3
	local manifest_dir="$workdir/manifests"
	local logs_dir="$workdir/logs"
	local status_dir="$workdir/status"
	local active=0
	local total="$#"
	local queued=0
	local rel key log_file status_file

	mkdir -p "$manifest_dir" "$logs_dir" "$status_dir"

	for rel in "$@"; do
		key="$(example_key "$rel")"
		log_file="$logs_dir/${key}.log"
		status_file="$status_dir/${key}.status"
		queued=$((queued + 1))

		printf '%s[%03d/%03d]%s %s\n' "$MAGENTA" "$queued" "$total" "$RESET" "$rel"

		(
			if "$SCRIPT_DIR/record-example.sh" "$REPO_ROOT/_examples/$rel" "$manifest_dir" >"$log_file" 2>&1; then
				printf '%s\tok\t%s\n' "$rel" "$log_file" >"$status_file"
			else
				printf '%s\tfail\t%s\n' "$rel" "$log_file" >"$status_file"
			fi
		) &

		active=$((active + 1))
		if ((active >= max_jobs)); then
			wait -n || true
			active=$((active - 1))
		fi
	done

	wait || true

	local failures=0
	local statuses=0
	local status rel_state result log
	for status in "$status_dir"/*.status; do
		[[ -e "$status" ]] || continue
		statuses=$((statuses + 1))
		IFS=$'\t' read -r rel_state result log <"$status"
		if [[ "$result" == "fail" ]]; then
			failures=$((failures + 1))
			print_file_group "Failed example: $rel_state" "$log"
		elif [[ "${ANIMATE_VERBOSE:-0}" == "1" || "${ANIMATE_MEASURE_ONLY:-0}" == "1" ]]; then
			print_file_group "Log: $rel_state" "$log"
		fi
	done

	if ((statuses != total)); then
		failures=$((failures + total - statuses))
		error "Only $statuses of $total examples reported a status"
	fi

	((failures == 0)) || die "$failures example recording(s) failed"

	cat "$manifest_dir"/*.tsv | sort >"$manifest"
}

# Lists every example with its terminal size, in the log and in the GitHub
# job summary.
print_summary() {
	local manifest="$1"
	local rel size source

	group "Terminal sizes"
	while IFS=$'\t' read -r rel _ size source; do
		printf '  %-50s %-8s %s\n' "$rel" "$size" "$source"
	done <"$manifest"
	end_group

	[[ -n "${GITHUB_STEP_SUMMARY:-}" ]] || return 0
	{
		printf '### Example animations\n\n'
		printf '| Example | Terminal | Size from |\n| --- | --- | --- |\n'
		while IFS=$'\t' read -r rel _ size source; do
			printf '| `%s` | %s | %s |\n' "$rel" "$size" "$source"
		done <"$manifest"
	} >>"$GITHUB_STEP_SUMMARY"
}

main() {
	local jobs workdir manifest rel
	local update_docs=1
	local -a selected=() examples=()

	while (($#)); do
		case "$1" in
			--dry-run)
				update_docs=0
				;;
			--example)
				shift
				[[ $# -gt 0 && -n "$1" ]] || die "--example requires a path like table/demo"
				rel="${1#_examples/}"
				rel="${rel%/main.go}"
				rel="${rel%/}"
				selected+=("$rel")
				update_docs=0
				;;
			--measure-only)
				export ANIMATE_MEASURE_ONLY=1
				update_docs=0
				;;
			--publish)
				export ANIMATE_PUBLISH=1
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
	require_cmd python3
	require_cmd timeout
	if [[ "${ANIMATE_MEASURE_ONLY:-0}" != "1" ]]; then
		require_cmd docker
	fi

	jobs="$(default_jobs)"
	[[ "$jobs" =~ ^[0-9]+$ && "$jobs" -gt 0 ]] || die "ANIMATE_JOBS must be a positive number"

	info "Repository: $REPO_ROOT"
	info "Concurrent jobs: $jobs"
	info "VHS Docker image: $VHS_DOCKER_IMAGE"
	info "Publish to vhs.charm.sh: ${ANIMATE_PUBLISH:-0}"
	info "Measure only: ${ANIMATE_MEASURE_ONLY:-0}"
	info "Update docs: $update_docs"
	end_group

	for rel in "${selected[@]}"; do
		[[ -f "$REPO_ROOT/_examples/$rel/main.go" ]] || die "Example not found: _examples/$rel"
	done

	mapfile -d '' examples < <(find_examples "${selected[@]}")
	((${#examples[@]} > 0)) || die "No examples found under _examples"

	workdir="$(mktemp -d)"
	# shellcheck disable=SC2064 # expand now, the variable is local
	trap "rm -rf '$workdir'" EXIT
	manifest="$workdir/manifest.tsv"

	group "Build ${#examples[@]} examples"
	build_examples "$workdir/bin" "$workdir/build-logs" "$jobs" "${examples[@]}"
	end_group

	export ANIMATE_BIN_DIR="$workdir/bin"

	info "Recording ${#examples[@]} examples with $jobs concurrent jobs"
	record_examples "$workdir" "$manifest" "$jobs" "${examples[@]}"
	print_summary "$manifest"

	if ((update_docs == 0)); then
		success "Finished without touching the docs"
		return
	fi

	"$SCRIPT_DIR/render-animation-docs.sh" "$manifest"

	if [[ "${ANIMATE_UPDATE_UNITTESTCOUNT:-1}" == "1" ]]; then
		"$SCRIPT_DIR/update-unittestcount.sh"
	else
		warn "Skipping unittest count update because ANIMATE_UPDATE_UNITTESTCOUNT=$ANIMATE_UPDATE_UNITTESTCOUNT"
	fi

	success "Animation pipeline finished"
}

main "$@"
