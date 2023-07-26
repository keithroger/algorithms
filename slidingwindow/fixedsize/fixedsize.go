package fixedsize

import "fmt"

// fixedsize returns the max sum of k contiguous values from arr.
func fixedSize(arr []int, size int) int {
	// Get sum of the first k sized section
	maxSum := sum(arr, 0, size+1)
	currSum := maxSum

	// Slide the window over the array
	for start := 1; start < len(arr)-size+1; start++ {
		fmt.Println(start)
		// Remove the first value.
		currSum -= arr[start-1]

		// Add the next value.
		currSum += arr[start+size-1]

		fmt.Printf("currSum: %v\n", currSum)

		// Update max sum if current is greater.
		if currSum > maxSum {
			maxSum = currSum
		}

	}

	return maxSum
}

func sum(arr []int, start, end int) int {
	sum := 0
	for ; start < end; start++ {
		sum += arr[start]
	}

	return start
}
