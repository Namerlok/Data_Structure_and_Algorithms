func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	right := 1
	for i := len(nums) - 2; i >= 0; i-- {
		right *= nums[i+1]
		answer[i] *= right
	}
	return answer
}