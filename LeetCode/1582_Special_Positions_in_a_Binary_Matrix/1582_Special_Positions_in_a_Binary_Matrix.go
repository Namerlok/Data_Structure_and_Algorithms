func numSpecial(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	countOneInCol := make([]int, n)
	countOneInRow := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				countOneInCol[j]++
				countOneInRow[i]++
			}
		}
	}
	var result int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && countOneInCol[j] == 1 && countOneInRow[i] == 1 {
				result++
			}
		}
	}
	return result
}