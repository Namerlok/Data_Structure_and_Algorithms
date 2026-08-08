func containsDuplicate(nums []int) bool {
	double := make(map[int]struct{})
	for _, val := range nums {
		if _, ok := double[val]; ok {
			return true
		}
		double[val] = struct{}{}
	}
	return false
}