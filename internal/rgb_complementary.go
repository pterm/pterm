package internal

// Complementary returns the complementary color of the given RGB values.
func Complementary(r, g, b uint8) (uint8, uint8, uint8) {
	return 255 - r, 255 - g, 255 - b
}
