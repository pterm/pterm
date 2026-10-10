#!/usr/bin/env python3
"""Measure captured terminal output so the animation terminal can fit it.

Usage: measure-output.py <columns> <file>

Prints "<widest> <rows>": the display width of the widest line, and how many
terminal rows the output fills at the given column count (long lines wrap).
Escape sequences are stripped, and wide characters such as CJK and most
emoji count as two cells.
"""

import re
import sys
import unicodedata

# CSI (colors, cursor movement), OSC (hyperlinks, titles) and the remaining
# two-byte escapes. None of them take up cells on screen.
ESCAPES = re.compile(
    r"\x1b\[[0-?]*[ -/]*[@-~]"
    r"|\x1b\][^\x07\x1b]*(?:\x07|\x1b\\)"
    r"|\x1b[@-Z\\-_]"
)


def char_width(ch):
    if unicodedata.combining(ch):
        return 0

    category = unicodedata.category(ch)
    # Format and control characters (zero width joiner, variation selectors
    # and the like) do not advance the cursor.
    if category in ("Cf", "Cc", "Mn", "Me"):
        return 0

    if unicodedata.east_asian_width(ch) in ("W", "F"):
        return 2

    return 1


def line_width(line):
    width = 0
    for ch in line:
        if ch == "\t":
            width += 8 - width % 8
        else:
            width += char_width(ch)

    return width


def visible_lines(text):
    text = ESCAPES.sub("", text).replace("\r\n", "\n")
    lines = text.split("\n")

    # The final newline ends the last line; it does not start a new one.
    if lines and lines[-1] == "":
        lines.pop()

    return lines


def main():
    if len(sys.argv) != 3:
        sys.exit("usage: measure-output.py <columns> <file>")

    columns = int(sys.argv[1])
    if columns <= 0:
        sys.exit("columns must be positive")

    # newline="" keeps carriage returns; universal newlines would turn them
    # into line breaks.
    with open(sys.argv[2], encoding="utf-8", errors="replace", newline="") as handle:
        lines = visible_lines(handle.read())

    widest = 0
    rows = 0
    for line in lines:
        # A carriage return moves back to the start of the line without
        # clearing it, as spinners do when they redraw in place. The line is
        # as wide as the longest text written onto it.
        width = max(line_width(part.rstrip()) for part in line.split("\r"))
        widest = max(widest, width)
        rows += max(1, -(-width // columns))

    print(widest, rows)


if __name__ == "__main__":
    main()
