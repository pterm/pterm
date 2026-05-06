#!/usr/bin/env bash

set -Eeuo pipefail

source "$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)/animation-common.sh"

example_dir="${1:-}"
manifest_dir="${2:-}"

[[ -n "$example_dir" ]] || die "Usage: record-example.sh <example-dir> <manifest-dir>"
[[ -n "$manifest_dir" ]] || die "Usage: record-example.sh <example-dir> <manifest-dir>"
[[ -f "$example_dir/main.go" ]] || die "Example has no main.go: $example_dir"

example_dir="$(cd -- "$example_dir" && pwd)"
example_rel="${example_dir#"$REPO_ROOT/_examples/"}"
key="$(example_key "$example_rel")"

mkdir -p "$manifest_dir" "$REPO_ROOT/testdata/raw"

tape_file="$example_dir/pterm-animation.tape"
gif_file="$example_dir/animation.gif"
binary_file="$example_dir/.pterm-vhs-example"
runner_file="$example_dir/.pterm-vhs-run"
golden_file="$REPO_ROOT/testdata/raw/${key}.txt"
manifest_file="$manifest_dir/${key}.tsv"
golden_rel="$(realpath --relative-to="$example_dir" "$golden_file")"
container_workdir="/vhs/_examples/$example_rel"
container_tape="$container_workdir/$(basename "$tape_file")"
container_runner="$container_workdir/$(basename "$runner_file")"
vhs_image="${VHS_DOCKER_IMAGE:-ghcr.io/charmbracelet/vhs:latest}"

cleanup() {
	rm -f "$tape_file" "$binary_file" "$runner_file"
}
trap cleanup EXIT

run_command="bash $container_runner"
theme="${VHS_THEME:-catppuccin-macchiato}"

write_keys_tape() {
	local keys_file="$example_dir/keys.tape"

	if [[ ! -f "$keys_file" ]]; then
		return
	fi

	printf '\n'
	info "Adding keys.tape for $example_rel" >&2
	sed '$a\' "$keys_file"
}

write_options_tape() {
	local options_file="$example_dir/options.tape"

	if [[ ! -f "$options_file" ]]; then
		return
	fi

	printf '\n'
	info "Adding options.tape for $example_rel" >&2
	sed '$a\' "$options_file"
}

if [[ "$example_rel" == interactive_* && ! -f "$example_dir/ci.go" && ! -f "$example_dir/keys.tape" ]]; then
	warn "No ci.go automation or keys.tape found for $example_rel"
fi

{
	cat <<EOF
Output animation.gif
Output $golden_rel

Set Width 1200
Set Height 600
Set Framerate 60
Set CursorBlink false
Set Theme "$theme"
EOF
	write_options_tape
	cat <<EOF

Hide
Type "$run_command"
Enter
Show
EOF
	write_keys_tape
	cat <<EOF
Wait@${VHS_RUN_TIMEOUT:-300s}
Sleep 5
EOF
} >"$tape_file"

info "Building $example_rel"
(cd "$example_dir" && go build -o "$binary_file" .)

cat >"$runner_file" <<'EOF'
#!/usr/bin/env bash
set -Eeuo pipefail

cd "$(dirname "$0")"
clear
export CI="${CI:-true}"
export TERM="${TERM:-xterm-256color}"
./.pterm-vhs-example
EOF
chmod +x "$runner_file"

docker_args=(
	run
	--rm
	--volume "$REPO_ROOT:/vhs"
	--workdir "$container_workdir"
	--env "HOME=/tmp"
)

if [[ "$(id -u)" != "0" ]]; then
	docker_args+=(--user "$(id -u):$(id -g)")
fi

info "Recording $example_rel"
docker "${docker_args[@]}" "$vhs_image" "$container_tape"

if [[ "${ANIMATE_PUBLISH:-1}" != "1" ]]; then
	warn "ANIMATE_PUBLISH is disabled; keeping local animation.gif for $example_rel"
	printf '%s\t%s\n' "$example_rel" "animation.gif" >"$manifest_file"
	exit 0
fi

info "Publishing $example_rel"
publish_output="$(docker "${docker_args[@]}" "$vhs_image" publish animation.gif 2>&1)"
gif_url="$(printf '%s\n' "$publish_output" | sed -nE 's#.*(https://vhs\.charm\.sh/[^ )"]+\.gif).*#\1#p' | head -n 1)"

if [[ -z "$gif_url" ]]; then
	printf '%s\n' "$publish_output"
	die "Could not find published GIF URL for $example_rel"
fi

printf '%s\t%s\n' "$example_rel" "$gif_url" >"$manifest_file"
success "Published $example_rel -> $gif_url"
