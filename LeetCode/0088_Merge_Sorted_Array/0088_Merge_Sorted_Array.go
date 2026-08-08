func merge(nums1 []int, m int, nums2 []int, n int) {
	pos1, pos2, posR := m-1, n-1, len(nums1)-1
	for posR >= 0 {
		if pos2 < 0 {
			break
		}
		if pos1 < 0 || nums1[pos1] < nums2[pos2] {
			nums1[posR] = nums2[pos2]
			pos2--
		} else {
			nums1[posR] = nums1[pos1]
			pos1--
		}
		posR--
	}
}