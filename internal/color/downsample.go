package color

// cubeSteps holds the color intensities of the xterm 6x6x6 color cube
// (palette indices 16-231).
var cubeSteps = [6]int{0, 95, 135, 175, 215, 255}

// basicPalette holds the xterm default values of the 16 base colors, indexed
// 0-15 in SGR order (black, red, green, yellow, blue, magenta, cyan, white,
// then their bright variants).
var basicPalette = [16][3]int{
	{0, 0, 0},       // black
	{205, 0, 0},     // red
	{0, 205, 0},     // green
	{205, 205, 0},   // yellow
	{0, 0, 238},     // blue
	{205, 0, 205},   // magenta
	{0, 205, 205},   // cyan
	{229, 229, 229}, // white
	{127, 127, 127}, // bright black
	{255, 0, 0},     // bright red
	{0, 255, 0},     // bright green
	{255, 255, 0},   // bright yellow
	{92, 92, 255},   // bright blue
	{255, 0, 255},   // bright magenta
	{0, 255, 255},   // bright cyan
	{255, 255, 255}, // bright white
}

// RGBTo256 returns the index of the xterm 256-color palette entry closest to
// the given 24-bit color. It considers both the 6x6x6 color cube (indices
// 16-231) and the grayscale ramp (indices 232-255) and picks the closer one,
// the same way tmux does.
func RGBTo256(r, g, b uint8) uint8 {
	ri, gi, bi := cubeIndex(r), cubeIndex(g), cubeIndex(b)
	cubeColor := [3]int{cubeSteps[ri], cubeSteps[gi], cubeSteps[bi]}
	cube := 16 + 36*ri + 6*gi + bi

	if cubeColor == [3]int{int(r), int(g), int(b)} {
		return uint8(cube) //nolint:gosec // 16 + 36*5 + 6*5 + 5 = 231 always fits.
	}

	// The grayscale ramp holds the grays 8, 18, ..., 238 at indices 232-255.
	grayIndex := max(min((int(r)+int(g)+int(b))/3-3, 235), 0) / 10
	gray := 8 + 10*grayIndex

	if colorDistance([3]int{gray, gray, gray}, r, g, b) < colorDistance(cubeColor, r, g, b) {
		return uint8(232 + grayIndex) //nolint:gosec // 232 + 235/10 = 255 always fits.
	}

	return uint8(cube) //nolint:gosec // 16 + 36*5 + 6*5 + 5 = 231 always fits.
}

// RGBToBasic returns the SGR foreground code (30-37 or 90-97) of the base
// ANSI color closest to the given 24-bit color. Add 10 for the matching
// background code.
func RGBToBasic(r, g, b uint8) uint8 {
	best, bestDistance := 0, colorDistance(basicPalette[0], r, g, b)

	for i := 1; i < len(basicPalette); i++ {
		if d := colorDistance(basicPalette[i], r, g, b); d < bestDistance {
			best, bestDistance = i, d
		}
	}

	if best < 8 {
		return uint8(30 + best)
	}

	return uint8(90 + best - 8)
}

// cubeIndex returns the index of the cube step closest to the given
// intensity.
func cubeIndex(v uint8) int {
	if v < 48 {
		return 0
	}

	if v < 114 {
		return 1
	}

	return int(v-35) / 40
}

// colorDistance returns the squared euclidean distance between a palette
// color and a 24-bit color.
func colorDistance(palette [3]int, r, g, b uint8) int {
	dr := palette[0] - int(r)
	dg := palette[1] - int(g)
	db := palette[2] - int(b)

	return dr*dr + dg*dg + db*db
}
