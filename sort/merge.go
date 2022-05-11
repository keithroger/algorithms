package sort

func Merge(arr []int) {
	msort(arr, 0, len(arr)-1)
}

func msort(arr []int, lo, hi int) {
	if hi <= lo {
		return
	}

	mid := lo + (hi-lo)/2
	msort(arr, lo, mid)
	msort(arr, mid+1, hi)
	merge(arr, lo, mid, hi)
}

func merge(arr []int, lo, mid, hi int) {
	i, j := lo, mid+1
	aux := arr

	for k := lo; k <= hi; k++ {
		if i > mid {
			arr[k] = aux[j]
			j++
		} else if j > hi {
			arr[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			arr[k] = aux[j]
			j++
		} else {
			arr[k] = aux[i]
			i++
		}
	}
}

func Merge2(arr []int) {
	mid := len(arr) / 2

	if len(arr) <= 1 {
		return
	}

	Merge2(arr[:mid])
	Merge2(arr[mid:])
	merge2(arr)
}

func merge2(arr []int) {
	mid, hi := len(arr)/2, len(arr)-1
	leftIdx, rightIdx := 0, mid
	aux := make([]int, len(arr))
	copy(aux, arr)

	for i := range arr {
		if leftIdx > mid-1 {
			arr[i] = aux[rightIdx]
			rightIdx++
		} else if rightIdx > hi {
			arr[i] = aux[leftIdx]
			leftIdx++
		} else if aux[leftIdx] < aux[rightIdx] {
			arr[i] = aux[leftIdx]
			leftIdx++
		} else {
			arr[i] = aux[rightIdx]
			rightIdx++
		}
	}

}
