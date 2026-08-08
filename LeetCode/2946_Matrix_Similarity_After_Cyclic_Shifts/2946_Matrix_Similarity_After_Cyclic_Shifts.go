func pow(base, exponent int) int {
	var res int = 1
	for i := 0; i < exponent; i++ {
		res *= base
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func areSimilar(mat [][]int, k int) bool {
	m, n := len(mat), len(mat[0])
	if k%n == 0 {
		return true
	}
	for i := 0; i < m; i++ {
		offset := pow(-1, i&1) * k
		for j := 0; j < n; j++ {
			if mat[i][j] != mat[i][abs(j+n+offset)%n] {
				return false
			}
		}
	}
	return true
}