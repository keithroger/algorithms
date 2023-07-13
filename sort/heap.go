package sort

func HeapSort(arr []int) {
	n := len(arr)

	// Organize elements into a max heap
	for i := n/2 - 1; i >= 0; i-- {
		sink(arr, i, n-1)
	}

	// Move max item into sorted position
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		sink(arr, 0, i-1)
	}
}

func sink(arr []int, idx, depth int) {

	for nextIdx := (idx+1)*2 - 1; nextIdx <= depth; nextIdx = (idx+1)*2 - 1 {

		if nextIdx+1 <= depth && arr[nextIdx+1] > arr[nextIdx] {
			nextIdx++
		}

		if arr[idx] < arr[nextIdx] {
			arr[idx], arr[nextIdx] = arr[nextIdx], arr[idx]
		} else {
			break
		}

		idx = nextIdx
	}
}
