package educ

func BinSearch(data []int, value int) int {
	start := 0
	end := len(data) - 1

	for start <= end {
		mid := int((start + end) / 2)
		if data[mid] == value {
			return mid
		} else {
			if data[mid] > value {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}
