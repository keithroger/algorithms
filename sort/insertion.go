package sort

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
