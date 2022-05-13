package sort

import (
	"fmt"
	"math/rand"
)

func Quick(arr []int) {
	fmt.Println(arr)
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	quick(arr)
	fmt.Printf("Ans: %v\n", arr)
	fmt.Println()
}

func quick(arr []int) {
	fmt.Println(arr)
	if len(arr) <= 1 {
		return
	}

	idx := partition(arr)
	quick(arr[:idx])
	quick(arr[idx+1:])
}

func partition(arr []int) int {
	pivot := arr[0]
	leftIdx, rightIdx := 1, len(arr)-1

	for leftIdx < rightIdx {
		for arr[leftIdx] < pivot && leftIdx != len(arr)-1 {
			leftIdx++
		}

		for pivot < arr[rightIdx] && rightIdx != 0 {
			rightIdx--
		}

		if leftIdx <= rightIdx {
			arr[leftIdx], arr[rightIdx] = arr[rightIdx], arr[leftIdx]
			leftIdx++
			rightIdx--
			fmt.Printf("Swap: %v\n", arr)
		}
	}

	arr[0], arr[leftIdx] = arr[leftIdx], arr[0]
	fmt.Printf("Pivot:%d\tL:%d\tR:%d\n", pivot, leftIdx, rightIdx)
	fmt.Println(arr)

	return leftIdx
}
