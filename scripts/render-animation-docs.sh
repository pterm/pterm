#!/usr/bin/env bash

set -Eeuo pipefail

source "$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)/animation-common.sh"

manifest="${1:-}"
[[ -n "$manifest" ]] || die "Usage: render-animation-docs.sh <manifest.tsv>"
[[ -f "$manifest" ]] || die "Manifest not found: $manifest"

declare -A gif_urls
declare -a sections

# Manifest rows: example path, GIF reference, terminal size, size source. The
# reference is "animation.gif" for a committed GIF or a vhs.charm.sh URL.
while IFS=$'\t' read -r rel url _; do
	[[ -n "$rel" && -n "$url" ]] || continue
	gif_urls["$rel"]="$url"
done <"$manifest"

section_seen=" "
while IFS= read -r section; do
	if [[ "$section_seen" != *" $section "* ]]; then
		sections+=("$section")
		section_seen+="$section "
	fi
done < <(printf '%s\n' "${!gif_urls[@]}" | awk -F/ '{print $1}' | sort)

# Prints the image source of an example's GIF as seen from base_dir, a
# directory relative to the repository root ("" for the root itself).
# Published URLs are used as they are.
gif_src() {
	local rel="$1"
	local base_dir="$2"
	local ref="${gif_urls[$rel]}"
	local target="_examples/$rel/animation.gif"

	if [[ "$ref" =~ ^https?:// ]]; then
		printf '%s' "$ref"
	elif [[ -z "$base_dir" ]]; then
		printf '%s' "$target"
	else
		printf '%s' "${target#"$base_dir/"}"
	fi
}

emit_source_block() {
	local main_file="$1"

	printf '```go\n'
	sed '$a\' "$main_file"
	printf '```\n'
}

emit_example_details() {
	local rel="$1"
	local base_dir="$2"
	local main_file="$REPO_ROOT/_examples/$rel/main.go"

	printf '### %s\n\n' "$rel"
	printf '![Animation](%s)\n\n' "$(gif_src "$rel" "$base_dir")"
	printf '<details>\n\n<summary>SHOW SOURCE</summary>\n\n'
	emit_source_block "$main_file"
	printf '\n</details>\n\n'
}

write_example_readme() {
	local rel="$1"
	local readme="$REPO_ROOT/_examples/$rel/README.md"
	local main_file="$REPO_ROOT/_examples/$rel/main.go"

	{
		printf '# %s\n\n' "$rel"
		printf '![Animation](%s)\n\n' "$(gif_src "$rel" "_examples/$rel")"
		emit_source_block "$main_file"
	} >"$readme"
}

ordered_examples_for_section() {
	local section="$1"

	if [[ -n "${gif_urls["$section/demo"]:-}" ]]; then
		printf '%s\n' "$section/demo"
	fi

	printf '%s\n' "${!gif_urls[@]}" |
		awk -F/ -v section="$section" '$1 == section && $2 != "demo" {print}' |
		sort
}

# Prints every example of a section, with GIF paths relative to base_dir.
emit_section() {
	local section="$1"
	local base_dir="$2"
	local rel

	while IFS= read -r rel; do
		[[ -n "$rel" ]] || continue
		emit_example_details "$rel" "$base_dir"
	done < <(ordered_examples_for_section "$section")
}

write_section_readmes() {
	local section rel

	for section in "${sections[@]}"; do
		emit_section "$section" "_examples/$section" >"$REPO_ROOT/_examples/$section/README.md"
		while IFS= read -r rel; do
			[[ -n "$rel" ]] || continue
			write_example_readme "$rel"
		done < <(ordered_examples_for_section "$section")
	done
}

# Writes the examples part of a README that lives in base_dir.
build_examples_body() {
	local output="$1"
	local base_dir="$2"
	local section

	: >"$output"
	for section in "${sections[@]}"; do
		emit_section "$section" "$base_dir" >>"$output"
	done
}


build_printers_table() {
	local output="$1"
	local -a printers=()
	local path feature name cell_count=0

	for path in "$REPO_ROOT/_examples"/*; do
		feature="$(basename "$path")"
		[[ -d "$path" && "$feature" != "demo" ]] || continue
		printers+=("$feature")
	done

	mapfile -t printers < <(printf '%s\n' "${printers[@]}" | sort)

	{
		printf '| Feature | Feature | Feature | Feature | Feature |\n'
		printf '| :-------: | :-------: | :-------: | :-------: | :-------: |\n'

		for feature in "${printers[@]}"; do
			if (( cell_count % 5 == 0 )); then
				printf '| '
			fi

			name="${feature^}"
			name="${name//_/ }"
			printf '%s <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/%s) |' "$name" "$feature"
			((cell_count += 1))

			if (( cell_count % 5 == 0 )); then
				printf '\n'
			fi
		done

		if (( cell_count % 5 != 0 )); then
			for ((i = cell_count % 5; i < 5; i++)); do
				printf ' |'
			done
			printf '\n'
		fi
	} >"$output"
}

build_website_printers() {
	local output="$1"
	local path feature

	: >"$output"
	for path in "$REPO_ROOT/_examples"/*; do
		feature="$(basename "$path")"
		[[ -d "$path" && "$feature" != "demo" ]] || continue
		printf '<li><a href="https://github.com/pterm/pterm/tree/master/_examples/%s">%s</a></li>\n' "$feature" "$feature" >>"$output"
	done
}

update_demo_links() {
	if [[ -z "${gif_urls["demo/demo"]:-}" ]]; then
		warn "No animation for demo/demo; leaving the README demo image alone"
		return
	fi

	DEMO_URL="$(gif_src demo/demo "")" REPO_ROOT="$REPO_ROOT" python3 - <<'PY'
import os
import re
from pathlib import Path

root = Path(os.environ["REPO_ROOT"])
demo_url = os.environ["DEMO_URL"]

readme = root / "README.md"
text = readme.read_text()
text, count = re.subn(
    r'(<a href="https://github\.com/pterm/pterm/tree/master/_examples/demo/demo"[^>]*>\s*<img src=")[^"]+(" alt="PTerm">)',
    r"\1" + demo_url + r"\2",
    text,
    count=1,
    flags=re.S,
)
if count != 1:
    raise SystemExit("could not update README demo image")
readme.write_text(text)

# The pterm.sh website is served from docs/ and cannot reach a GIF through a
# repository path, so it only follows published URLs.
if not demo_url.startswith(("http://", "https://")):
    raise SystemExit(0)

index =root / "docs" / "index.html"
text = index.read_text()
text, count = re.subn(
    r'(<img alt="PTerm demo animation"\s+src=")[^"]+("/>)',
    r"\1" + demo_url + r"\2",
    text,
    count=1,
    flags=re.S,
)
if count != 1:
    raise SystemExit("could not update docs/index.html demo image")
index.write_text(text)
PY
}

workdir="$(mktemp -d)"
trap 'rm -rf "$workdir"' EXIT

readme_examples="$workdir/readme-examples.md"
examples_readme_examples="$workdir/examples-readme-examples.md"
printers_table="$workdir/printers.md"
website_printers="$workdir/website-printers.html"

group "Generate markdown"
info "Rendering example README files"
write_section_readmes
build_examples_body "$readme_examples" ""
build_examples_body "$examples_readme_examples" "_examples"
replace_between "examples" "$REPO_ROOT/README.md" "$readme_examples"
replace_between "examples" "$REPO_ROOT/_examples/README.md" "$examples_readme_examples"

info "Rendering component lists"
build_printers_table "$printers_table"
build_website_printers "$website_printers"
replace_between "printers" "$REPO_ROOT/README.md" "$printers_table"
replace_between "printers" "$REPO_ROOT/docs/index.html" "$website_printers"

info "Updating demo image links"
update_demo_links

info "Removing committed SVG animation leftovers"
find "$REPO_ROOT/_examples" -name 'animation.svg' -delete
end_group

success "Rendered markdown from the animation manifest"
