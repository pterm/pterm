#!/usr/bin/env bash

set -Eeuo pipefail

source "$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)/animation-common.sh"

manifest="${1:-}"
[[ -n "$manifest" ]] || die "Usage: render-animation-docs.sh <published-manifest.tsv>"
[[ -f "$manifest" ]] || die "Manifest not found: $manifest"

declare -A gif_urls
declare -a sections

while IFS=$'\t' read -r rel url; do
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

emit_source_block() {
	local main_file="$1"

	printf '```go\n'
	sed '$a\' "$main_file"
	printf '```\n'
}

emit_example_details() {
	local rel="$1"
	local url="$2"
	local main_file="$REPO_ROOT/_examples/$rel/main.go"

	printf '### %s\n\n' "$rel"
	printf '![Animation](%s)\n\n' "$url"
	printf '<details>\n\n<summary>SHOW SOURCE</summary>\n\n'
	emit_source_block "$main_file"
	printf '\n</details>\n\n'
}

write_example_readme() {
	local rel="$1"
	local url="$2"
	local readme="$REPO_ROOT/_examples/$rel/README.md"
	local main_file="$REPO_ROOT/_examples/$rel/main.go"

	{
		printf '# %s\n\n' "$rel"
		printf '![Animation](%s)\n\n' "$url"
		emit_source_block "$main_file"
	} >"$readme"
}

ordered_examples_for_section() {
	local section="$1"
	local rel

	if [[ -n "${gif_urls["$section/demo"]:-}" ]]; then
		printf '%s\n' "$section/demo"
	fi

	printf '%s\n' "${!gif_urls[@]}" |
		awk -F/ -v section="$section" '$1 == section && $2 != "demo" {print}' |
		sort
}

build_section_readme() {
	local section="$1"
	local section_file="$REPO_ROOT/_examples/$section/README.md"
	local rel

	: >"$section_file"
	while IFS= read -r rel; do
		[[ -n "$rel" ]] || continue
		emit_example_details "$rel" "${gif_urls[$rel]}" >>"$section_file"
		write_example_readme "$rel" "${gif_urls[$rel]}"
	done < <(ordered_examples_for_section "$section")
}

build_examples_body() {
	local output="$1"
	local section

	: >"$output"
	for section in "${sections[@]}"; do
		build_section_readme "$section"
		cat "$REPO_ROOT/_examples/$section/README.md" >>"$output"
	done
}

build_printers_table() {
	local output="$1"
	local -a printers=()
	local path feature name cell_count=0

	for path in "$REPO_ROOT/_examples"/*; do
		feature="$(basename "$path")"
		[[ "$feature" == "README.md" || "$feature" == "demo" ]] && continue
		printers+=("$feature")
	done

	IFS=$'\n' printers=($(sort <<<"${printers[*]}"))
	unset IFS

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
		[[ "$feature" == "README.md" || "$feature" == "demo" ]] && continue
		printf '<li><a href="https://github.com/pterm/pterm/tree/master/_examples/%s">%s</a></li>\n' "$feature" "$feature" >>"$output"
	done
}

update_demo_links() {
	local demo_url="${gif_urls["demo/demo"]:-}"
	[[ -n "$demo_url" ]] || die "Published URL for demo/demo is missing"

	DEMO_URL="$demo_url" REPO_ROOT="$REPO_ROOT" python3 - <<'PY'
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

index = root / "docs" / "index.html"
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

examples_body="$workdir/examples.md"
printers_table="$workdir/printers.md"
website_printers="$workdir/website-printers.html"

group "Generate markdown"
info "Rendering example README files"
build_examples_body "$examples_body"
replace_between "examples" "$REPO_ROOT/README.md" "$examples_body"
replace_between "examples" "$REPO_ROOT/_examples/README.md" "$examples_body"

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

success "Rendered markdown from VHS publish manifest"
