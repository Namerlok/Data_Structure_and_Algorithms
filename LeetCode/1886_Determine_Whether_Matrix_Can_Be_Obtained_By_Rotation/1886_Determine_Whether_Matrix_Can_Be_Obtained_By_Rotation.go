func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	var is0Rotate bool = true
	var is90Rotate bool = true
	var is180Rotate bool = true
	var is270Rotate bool = true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				is0Rotate = false
			}
			if mat[i][j] != target[j][n-1-i] {
				is90Rotate = false
			}
			if mat[i][j] != target[n-1-i][n-1-j] {
				is180Rotate = false
			}
			if mat[i][j] != target[n-1-j][i] {
				is270Rotate = false
			}
			if !is0Rotate && !is90Rotate && !is180Rotate && !is270Rotate {
				return false
			}
		}
	}
	return true
}