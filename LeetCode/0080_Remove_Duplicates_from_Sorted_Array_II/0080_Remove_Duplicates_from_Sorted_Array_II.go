func removeDuplicates(nums []int) int {
	write := 0

	for _, num := range nums {
		if write < 2 || num != nums[write-2] {
			nums[write] = num
			write++
		}
	}

	return write
}