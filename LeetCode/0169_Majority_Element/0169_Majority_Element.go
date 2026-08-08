func majorityElement(nums []int) int {
	majEl, countMajEl := 0, 0
	for _, num := range nums {
		if countMajEl == 0 {
			majEl = num
		}

		if majEl == num {
			countMajEl++
		} else {
			countMajEl--
		}
	}
	return majEl
}