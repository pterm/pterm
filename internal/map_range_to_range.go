package internal

func MapRangeToRange(fromMin, fromMax, toMin, toMax, current float32) int {
	return int(toMin + ((toMax-toMin)/(fromMax-fromMin))*(current-fromMin))
}
