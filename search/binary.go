func BinarySearch(arr []ints, key int) {
	lo := 0
	hi := arr.length - 1

	for ;lo <= hi {
		mid := lo + (hi - lo) / 2

		if key < arr[mid] {
			hi = mid - 1
		} else if key > arr[mid] {
			lo = mid + 1
		} else {
			return mid
		}

		return -1
	}
}