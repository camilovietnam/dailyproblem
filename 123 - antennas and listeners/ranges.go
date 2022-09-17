package ranges

func MinRange(listeners []int, antennas []int) int {
	max := -1
	min := -1

	for _, listener := range listeners {
		min = -1
		for _, antenna := range antennas {
			if abs(listener-antenna) < min || min == -1 {
				min = abs(listener - antenna)
			}
		}

		if min > max || max == -1 {
			max = min
		}
	}

	return max
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
