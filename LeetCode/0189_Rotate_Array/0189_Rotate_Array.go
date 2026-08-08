func rotate(nums []int, k int) {
	startCyclEl := 0
	oldPos := 0
	for i := 0; i < len(nums)-1; i++ {
		pos := (oldPos + len(nums) - (k % len(nums))) % len(nums)
		if pos == startCyclEl {
			oldPos = startCyclEl + 1
			startCyclEl = oldPos
			continue
		}
		nums[oldPos], nums[pos] = nums[pos], nums[oldPos]
		oldPos = pos
	}
}