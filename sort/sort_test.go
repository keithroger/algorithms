package sort_test

import (
	"testing"

	"github.com/keithroger/algorithms/sort"
)

func TestSorts(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	tt := []struct {
		name string
		sort func([]int)
	}{
		{"SelectionSort", sort.SelectionSort},
		{"InsertionSort", sort.InsertionSort},
	}

	for _, tc := range tt {
		tc := tc // make a copy to be parallel friendly

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			arr := arr // make a copy to pass to by referance
			tc.sort(arr)

			if !isSorted(arr) {
				t.Error("Out of order")
			}
		})
	}
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if !(arr[i-1] < arr[i]) {
			return false
		}
	}

	return true
}
