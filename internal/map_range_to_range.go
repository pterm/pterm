package internal

func MapRangeToRange(fromMin, fromMax, toMin, toMax, current float32) int {
	if fromMax-fromMin == 0 {
		return 0
	}
	return int(toMin + ((toMax-toMin)/(fromMax-fromMin))*(current-fromMin))
}
