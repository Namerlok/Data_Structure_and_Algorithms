func minSwaps(grid [][]int) int {
	n := len(grid)
	countZero := make([]int, n)
	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0 && grid[i][j] != 1; j-- {
			countZero[i]++
		}
	}
	countSwap := 0
	for i := 0; i < n; i++ {
		if countZero[i] < n-i-1 {
			targetPos := -1
			for j := i + 1; j < n; j++ {
				if countZero[j] >= n-i-1 {
					targetPos = j
					break
				}
			}
			if targetPos == -1 {
				return -1
			}
			countSwap += targetPos - i
			for j := targetPos; j > i; j-- {
				countZero[j-1], countZero[j] = countZero[j], countZero[j-1]
			}
		}
	}
	return countSwap
}