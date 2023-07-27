package dynamicsize

// DynamicSize finds the smallest window size greater than goal.
// func DynamicSize(arr []int, goal int) int {
// minWindowSize := 999
// currSum := 0
//
// start := 0
// for end := 0; end < len(arr); end++ {
// currSum += arr[end]
//
// for ; start < end && currSum >= goal; start++ {
// currSum -= arr[start-1]
//
// windowSize := end - start + 1
// fmt.Printf("start: %v end: %v sum: %v windowSize:%v\n", start, end, currSum, windowSize)
// if windowSize < minWindowSize {
// minWindowSize = windowSize
// }
//
// }
// }
//
// return minWindowSize
// }

// DynamicSize finds the smallest window size greater than goal.
func DynamicSize(arr []int, goal int) int {
	minWindowSize := 999
	currSum := 0        // current sum
	left, right := 0, 0 // points to left and right positions in the array

	for left < len(arr) {
		if currSum < goal && right < len(arr) {
			currSum += arr[right]
			right++
		} else if currSum >= goal {
			currSum -= arr[left]
			left++

			if windowSize := right - left + 1; windowSize < minWindowSize {
				minWindowSize = windowSize
			}
		} else {
			// sum < goal and left has not made it to the left
			// When this happens we do not need to process the rest.
			break
		}

	}

	return minWindowSize

}
