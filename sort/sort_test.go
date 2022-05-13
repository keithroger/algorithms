package sort_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/keithroger/algorithms/sort"
)

var (
	arrs = []struct {
		name string
		nums []int
	}{
		{"Ascending", []int{1, 2, 3, 4}},
		{"Descending", []int{4, 3, 2, 1}},
		{"Empty", []int{}},
		{"OddCount", []int{0, 1, 2}},
		{"Negatives", []int{-5, -2, -1, -3}},
		{"Repeats", []int{2, 2, 1, 2, 3, 4}},
		{"Uniform", []int{3, 3, 3, 3, 3, 3}},
		{"Random1", []int{-10, 10, -20, 3, 4, 2}},
		{"Random2", []int{-30, 2, 3, -40, 5, 100, -200}},
		{"Random3", []int{20, 2, 4, -2, -3, -5, -1, -5}},
	}
)

func TestSorts(t *testing.T) {
	sorts := []struct {
		name string
		sort func([]int)
	}{
		{"SelectionSort", sort.Selection},
		{"InsertionSort", sort.Insertion},
		{"MergeSort", sort.Merge},
		{"MergeSort2", sort.Merge2},
		{"ShellSort", sort.Shell},
		// {"QuickSort", sort.Quick},
	}

	rand.Seed(1984)

	for _, s := range sorts {
		for _, arr := range arrs {
			s := s // make a copy to be parallel friendly
			nums := make([]int, len(arr.nums))
			copy(nums, arr.nums)

			name := fmt.Sprintf("%s_%s", s.name, arr.name)
			t.Run(name, func(t *testing.T) {
				// t.Parallel()

				fmt.Println(name)
				s.sort(nums)

				if !isSorted(nums) {
					t.Errorf("%s out of order\nGot: %v", name, nums)
				}
			})
		}
	}
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if !(arr[i-1] <= arr[i]) {
			return false
		}
	}

	return true
}
