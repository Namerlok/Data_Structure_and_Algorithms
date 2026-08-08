func jump(nums []int) int {
	curEnd := 0
	farthest := 0
	minJump := 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == curEnd {
			minJump++
			curEnd = farthest
		}
	}
	return minJump
}