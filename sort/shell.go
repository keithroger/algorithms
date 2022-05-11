package sort

func Shell(arr []int) {
	h := 1
	for h < len(arr)/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < len(arr); i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j] = arr[j-h]
			}
		}
		h /= 3
	}
}
