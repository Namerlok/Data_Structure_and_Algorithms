func numberOfSubmatrices(grid [][]byte) int {
	var result int
	n := len(grid)
	m := len(grid[0])
	cols := make([]int, m)
	colsIncludeX := make([]bool, m)

	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			if j-1 >= 0 && colsIncludeX[j-1] {
				colsIncludeX[j] = true
			}
			switch grid[i][j] {
			case 'X':
				cols[j]++
				colsIncludeX[j] = true
			case 'Y':
				cols[j]--
			case '.':
			}
			sum += cols[j]
			if colsIncludeX[j] && sum == 0 {
				result++
			}
		}
	}
	return result
}