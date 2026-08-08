func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	for i := y; i < y+k; i++ {
		j, k := x, x+k-1
		for j < k {
			grid[j][i], grid[k][i] = grid[k][i], grid[j][i]
			j++
			k--
		}
	}
	return grid
}