package kadanes

var NINF = -999

// Kadanes returns the greatest contiguous sum from arr.
func Kadanes(arr []int) int {
	currMax, globalMax := NINF, NINF

	for _, num := range arr {
		currMax = max(currMax+num, num)
		globalMax = max(currMax, globalMax)
	}

	return globalMax
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
