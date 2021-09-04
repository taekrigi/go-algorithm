// https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(a []int, b []int) float64 {
	aIndex := 0
	bIndex := 0
	totalSize := len(a) + len(b)
	
	result := make([]int, totalSize)

	for i := 0; i < totalSize; i++ {
			min, _aIndex, _bIndex := getMin(a, b, aIndex, bIndex)
			aIndex = _aIndex
			bIndex = _bIndex
			result[i] = min
	}

	return getMedian(result)
}

func getMin(a, b []int, aIndex, bIndex int) (int, int, int) {
	if len(a) <= aIndex {
			return b[bIndex], aIndex, bIndex + 1
	}
	
	if len(b) <= bIndex {
			return a[aIndex], aIndex + 1, bIndex
	}
	
	if a[aIndex] <= b[bIndex] {
			return a[aIndex], aIndex + 1, bIndex
	}
	
	return b[bIndex], aIndex, bIndex + 1
}

func getMedian(result []int) float64 {
	totalSize := len(result)
	middle := totalSize / 2
	
	if totalSize % 2 == 0 {
			return float64(result[middle - 1] + result[middle]) / 2
	}
	
	return float64(result[middle])
}
