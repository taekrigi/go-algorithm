func binarySearch(arr []int, x int) bool {
	arr = []int {1,2,3,4,5,6,7,8,9,10}
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

