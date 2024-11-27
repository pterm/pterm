package putils

import (
	"strconv"
	"strings"

	"github.com/pterm/pterm"
)

// RGBFromHEX converts a HEX and returns a new RGB.
// If the hex code is not valid it returns pterm.ErrHexCodeIsInvalid.
func RGBFromHEX(hex string) (pterm.RGB, error) {
	hex = strings.ToLower(hex)
	hex = strings.ReplaceAll(hex, "#", "")
	hex = strings.ReplaceAll(hex, "0x", "")

	if len(hex) == 3 {
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}
	if len(hex) != 6 {
		return pterm.RGB{}, pterm.ErrHexCodeIsInvalid
	}

	i64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return pterm.RGB{}, err
	}
	c := int(i64)

	// #nosec G115
	return pterm.RGB{
		R: uint8(c >> 16),             //nolint:gosec
		G: uint8((c & 0x00FF00) >> 8), //nolint:gosec
		B: uint8(c & 0x0000FF),        //nolint:gosec
	}, nil
}
