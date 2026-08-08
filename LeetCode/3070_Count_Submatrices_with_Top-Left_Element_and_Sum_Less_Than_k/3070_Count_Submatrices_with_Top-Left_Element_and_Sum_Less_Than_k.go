func countSubmatrices(grid [][]int, k int) int {
	var result int
	n := len(grid)
	m := len(grid[0])
	cols := make([]int, m)

	for i := 0; i < n; i++ {
		row := 0
		for j := 0; j < m; j++ {
			cols[j] += grid[i][j]
			row += cols[j]
			if k >= row {
				result++
			}
		}
	}

	return result
}