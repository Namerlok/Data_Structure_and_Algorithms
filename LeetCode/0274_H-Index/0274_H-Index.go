func hIndex(citations []int) int {
	count := make([]int, len(citations)+1, len(citations)+1)
	for _, citation := range citations {
		if citation >= len(citations) {
			count[len(citations)]++
		} else {
			count[citation]++
		}
	}

	for i := len(count) - 1; i > 0; i-- {
		count[i-1] += count[i]
		if count[i] >= i {
			return i
		}
	}

	return 0
}