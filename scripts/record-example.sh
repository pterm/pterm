#!/usr/bin/env bash

# Records one example as animation.gif with VHS.
#
# Usage: record-example.sh <example-dir> <manifest-dir>
#
# The terminal is sized to fit the example. Non-interactive examples run once
# headless first so their output can be measured; interactive examples (the
# ones with a keys.tape) need a TTY, so they use size.env or a default size.
# See scripts/README.md for every knob.

set -Eeuo pipefail

source "$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)/animation-common.sh"

example_dir="${1:-}"
manifest_dir="${2:-}"

[[ -n "$example_dir" && -n "$manifest_dir" ]] || die "Usage: record-example.sh <example-dir> <manifest-dir>"
[[ -f "$example_dir/main.go" ]] || die "Example has no main.go: $example_dir"

example_dir="$(cd -- "$example_dir" && pwd)"
example_rel="${example_dir#"$REPO_ROOT/_examples/"}"
key="$(example_key "$example_rel")"

min_columns="${ANIMATE_MIN_COLUMNS:-80}"
max_columns="${ANIMATE_MAX_COLUMNS:-120}"
column_padding="${ANIMATE_COLUMN_PADDING:-4}"
min_rows="${ANIMATE_MIN_ROWS:-8}"
default_extra_rows="${ANIMATE_EXTRA_ROWS:-2}"
interactive_columns="${ANIMATE_INTERACTIVE_COLUMNS:-100}"
interactive_rows="${ANIMATE_INTERACTIVE_ROWS:-20}"
measure_timeout="${ANIMATE_MEASURE_TIMEOUT:-60}"
record_timeout="${ANIMATE_RECORD_TIMEOUT:-600}"
size_mode="${ANIMATE_SIZE_MODE:-cells}"

font_size="${VHS_FONT_SIZE:-18}"
line_height="${VHS_LINE_HEIGHT:-1.2}"
padding="${VHS_PADDING:-24}"
framerate="${VHS_FRAMERATE:-30}"
theme="${VHS_THEME:-Catppuccin Macchiato}"
end_sleep="${VHS_END_SLEEP:-4s}"
wait_timeout="${VHS_RUN_TIMEOUT:-300s}"

# Cell size in pixels, only used with ANIMATE_SIZE_MODE=pixels. Taken from
# VHS 0.12.1 recordings with its default font at FontSize 18 and LineHeight
# 1.2, rounded up a little. Change them together with the font settings.
cell_width="${VHS_CELL_WIDTH:-11.5}"
cell_height="${VHS_CELL_HEIGHT:-28}"

gif_file="$example_dir/animation.gif"
manifest_file="$manifest_dir/${key}.tsv"
keys_file="$example_dir/keys.tape"
options_file="$example_dir/options.tape"
size_file="$example_dir/size.env"

container_example_dir="/vhs/_examples/$example_rel"

mkdir -p "$manifest_dir"
stage_dir="$(mktemp -d "${TMPDIR:-/tmp}/pterm-vhs-${key}.XXXXXX")"
container_name="pterm-vhs-${key//[^a-zA-Z0-9_.-]/-}-$$"

cleanup() {
	docker rm --force "$container_name" >/dev/null 2>&1 || true
	rm -rf "$stage_dir"
}
trap cleanup EXIT

clamp() {
	local value="$1" low="$2" high="$3"

	((value < low)) && value="$low"
	((value > high)) && value="$high"
	printf '%s\n' "$value"
}

is_number() {
	[[ "$1" =~ ^[0-9]+$ && "$1" -gt 0 ]]
}

# Reads KEY=VALUE from size.env without sourcing it, so the file can never
# run code.
size_env_value() {
	local name="$1"

	[[ -f "$size_file" ]] || return 0
	sed -nE "s/^[[:space:]]*(export[[:space:]]+)?${name}[[:space:]]*=[[:space:]]*\"?([0-9]+)\"?[[:space:]]*(#.*)?$/\2/p" "$size_file" | tail -n 1
}

# Reads "Set <Name> <number>" from options.tape.
options_tape_value() {
	local name="$1"

	[[ -f "$options_file" ]] || return 0
	sed -nE "s/^[[:space:]]*Set[[:space:]]+${name}[[:space:]]+([0-9]+)[[:space:]]*$/\1/p" "$options_file" | tail -n 1
}

# Runs the example the way it would run in a pipe and writes its output to $2.
run_headless() {
	local columns="$1"
	local output="$2"
	local status=0

	(
		cd "$example_dir"
		unset NO_COLOR FORCE_COLOR LINES
		CI=true COLUMNS="$columns" TERM=xterm-256color PTERM_THEME=macchiato \
			timeout --kill-after=5s "${measure_timeout}s" "$binary" </dev/null >"$output" 2>&1
	) || status=$?

	if ((status == 124 || status == 137)); then
		return 124
	fi

	if ((status != 0)); then
		warn "$example_rel exited with status $status during measuring; using its output anyway"
	fi
}

# Prints "<widest> <rows>" for the output at the given width.
measure() {
	local columns="$1"
	local output="$stage_dir/measure-$columns.txt"

	run_headless "$columns" "$output" || return $?
	python3 "$SCRIPT_DIR/measure-output.py" "$columns" "$output"
}

# The measuring run needs a binary for this machine, the recording one a Linux
# binary for the container. On Linux both are the same file.
binary="$stage_dir/example"
recording_binary="$stage_dir/example.linux"
prebuilt="${ANIMATE_BIN_DIR:+$ANIMATE_BIN_DIR/$key}"

if [[ -n "$prebuilt" && -x "$prebuilt" ]]; then
	cp "$prebuilt" "$binary"
	if [[ -x "$prebuilt.linux" ]]; then
		cp "$prebuilt.linux" "$recording_binary"
	else
		cp "$prebuilt" "$recording_binary"
	fi
else
	info "Building $example_rel"
	build_example "$example_dir" "$binary"
	if [[ "$(go env GOOS)" == "linux" ]]; then
		cp "$binary" "$recording_binary"
	else
		build_example "$example_dir" "$recording_binary" linux
	fi
fi

# Work out the terminal size.

columns="$(size_env_value COLUMNS)"
rows="$(size_env_value ROWS)"
extra_rows="$(size_env_value EXTRA_ROWS)"
extra_rows="${extra_rows:-$default_extra_rows}"
size_source="measured"

# options.tape wins over everything else. VHS refuses a tape that sets both
# Width and Columns (or Height and Rows), so the generated size is left out for
# whatever it already sets.
options_width="$(options_tape_value Width)"
options_height="$(options_tape_value Height)"
options_columns="$(options_tape_value Columns)"
options_rows="$(options_tape_value Rows)"
[[ -n "$options_columns" ]] && columns="$options_columns"
[[ -n "$options_rows" ]] && rows="$options_rows"

if [[ -n "$options_width" || -n "$options_height" ]]; then
	warn "$example_rel sets a pixel size in options.tape; prefer COLUMNS and ROWS in size.env"
fi

if [[ "$example_rel" == interactive_* && ! -f "$keys_file" ]]; then
	warn "$example_rel looks interactive but has no keys.tape, so nothing will answer its prompts"
fi

if [[ -f "$keys_file" ]]; then
	# Interactive examples wait for input, so a headless run would only show
	# the first prompt or hang.
	size_source="interactive default"
	[[ -n "$columns" || -n "$rows" ]] && size_source="size.env"
	columns="${columns:-$interactive_columns}"
	rows="${rows:-$interactive_rows}"
elif [[ -n "$columns" && -n "$rows" ]]; then
	size_source="size.env"
else
	measured=""
	if [[ -n "$columns" ]]; then
		measured="$(measure "$columns")" || measured=""
	else
		# Measure at the narrowest width first, and only widen the terminal
		# when the output needs it. Output that fills whatever width it gets,
		# like a full-width header, then stays at the minimum.
		if measured="$(measure "$min_columns")"; then
			read -r widest _ <<<"$measured"
			columns="$(clamp "$((widest + column_padding))" "$min_columns" "$max_columns")"
			if ((columns != min_columns)); then
				measured="$(measure "$columns")" || measured=""
			fi
		else
			measured=""
		fi
	fi

	if [[ -n "$measured" ]]; then
		read -r widest measured_rows <<<"$measured"
		info "$example_rel output: widest line $widest cells, $measured_rows rows at $columns columns"
		if [[ -z "$rows" ]]; then
			rows="$(( measured_rows + extra_rows ))"
			((rows < min_rows)) && rows="$min_rows"
		else
			size_source="measured + size.env"
		fi
	else
		warn "Could not measure $example_rel (timed out after ${measure_timeout}s); using ${interactive_columns}x${interactive_rows}"
		size_source="fallback"
		columns="${columns:-$interactive_columns}"
		rows="${rows:-$interactive_rows}"
	fi
fi

is_number "$columns" || die "Invalid column count for $example_rel: $columns"
is_number "$rows" || die "Invalid row count for $example_rel: $rows"
info "$example_rel terminal: ${columns}x${rows} ($size_source)"

# Write the tape.

size_settings() {
	if [[ "$size_mode" == "pixels" ]]; then
		[[ -z "$options_width" && -z "$options_columns" ]] &&
			printf 'Set Width %s\n' "$(pixels "$columns" "$cell_width")"
		[[ -z "$options_height" && -z "$options_rows" ]] &&
			printf 'Set Height %s\n' "$(pixels "$rows" "$cell_height")"
	else
		[[ -z "$options_width" && -z "$options_columns" ]] && printf 'Set Columns %s\n' "$columns"
		[[ -z "$options_height" && -z "$options_rows" ]] && printf 'Set Rows %s\n' "$rows"
	fi

	return 0
}

# Converts a cell count to pixels, padding included, rounded up to an even
# number because video encoders need even dimensions.
pixels() {
	python3 -c 'import math, sys
cells, cell, pad = int(sys.argv[1]), float(sys.argv[2]), int(sys.argv[3])
px = math.ceil(cells * cell) + 2 * pad
print(px + px % 2)' "$1" "$2" "$padding"
}

append_file() {
	local file="$1"

	[[ -f "$file" ]] || return 0
	printf '\n'
	sed '$a\' "$file"
}

tape_file="$stage_dir/animation.tape"
write_tape() {
	{
	cat <<EOF
Output animation.gif

Set Shell "bash"
Set FontSize $font_size
Set LineHeight $line_height
Set Padding $padding
EOF
	size_settings
	cat <<EOF
Set Framerate $framerate
Set CursorBlink false
Set Theme "$theme"
EOF
	append_file "$options_file"
	cat <<EOF

Hide
Type "bash /work/run.sh"
Enter
Show
EOF
	append_file "$keys_file"
	cat <<EOF

Wait@$wait_timeout
Sleep $end_sleep
EOF
	} >"$tape_file"
}

write_tape

cat >"$stage_dir/run.sh" <<EOF
#!/usr/bin/env bash
cd "$container_example_dir" || exit 1
clear
unset NO_COLOR FORCE_COLOR CI
export TERM=xterm-256color COLUMNS=$columns PTERM_THEME=macchiato
exec /work/example.linux
EOF
chmod +x "$stage_dir/run.sh"

if [[ "${ANIMATE_MEASURE_ONLY:-0}" == "1" ]]; then
	print_file_group "Tape for $example_rel" "$tape_file"
	printf '%s\t%s\t%sx%s\t%s\n' "$example_rel" "-" "$columns" "$rows" "$size_source" >"$manifest_file"
	exit 0
fi

docker_args=(
	run
	--rm
	--init
	--name "$container_name"
	--volume "$REPO_ROOT:/vhs"
	--volume "$stage_dir:/work"
	--workdir "$container_example_dir"
	--env "HOME=/tmp"
)

if [[ "$(id -u)" != "0" ]]; then
	docker_args+=(--user "$(id -u):$(id -g)")
fi

# VHS sizes the browser viewport in pixels and gives up when no viewport
# yields exactly the requested rows, which happens for some row counts at
# this font size. One more row always resolves, so retry a few times with a
# taller terminal before failing.
vhs_log="$stage_dir/vhs.log"
for ((attempt = 1; ; attempt++)); do
	info "Recording $example_rel"
	rm -f "$gif_file"
	if timeout --kill-after=10s "${record_timeout}s" docker "${docker_args[@]}" "$VHS_DOCKER_IMAGE" /work/animation.tape 2>&1 | tee "$vhs_log"; then
		break
	fi

	if ((attempt < 3)) && grep -q 'could not resolve a viewport' "$vhs_log"; then
		rows=$((rows + 1))
		warn "$example_rel: VHS could not size the terminal, retrying at ${columns}x${rows}"
		write_tape
		continue
	fi

	die "VHS failed or timed out after ${record_timeout}s for $example_rel"
done
[[ -s "$gif_file" ]] || die "VHS finished but did not write $gif_file"

gif_ref="animation.gif"
if [[ "${ANIMATE_PUBLISH:-0}" == "1" ]]; then
	info "Publishing $example_rel"
	publish_output="$(docker "${docker_args[@]}" "$VHS_DOCKER_IMAGE" publish animation.gif 2>&1)" ||
		die "vhs publish failed for $example_rel: $publish_output"
	gif_ref="$(printf '%s\n' "$publish_output" | sed -nE 's#.*(https://vhs\.charm\.sh/[^ )"]+\.gif).*#\1#p' | head -n 1)"
	[[ -n "$gif_ref" ]] || die "Could not find published GIF URL for $example_rel in: $publish_output"
	success "Published $example_rel -> $gif_ref"
fi

printf '%s\t%s\t%sx%s\t%s\n' "$example_rel" "$gif_ref" "$columns" "$rows" "$size_source" >"$manifest_file"
success "Recorded $example_rel at ${columns}x${rows}"
