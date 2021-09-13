func binarySearch(arr []int, x int) bool {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + ((right - left) / 2)
		if arr[mid] == x {
			return true
		} else if x < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

