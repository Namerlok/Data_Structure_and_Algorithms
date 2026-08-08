func constructProductMatrix(grid [][]int) [][]int {
	mod := 12345
	n, m := len(grid), len(grid[0])

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			res[i][m-1] = 1
		} else {
			res[i][m-1] = (res[i+1][0] * (grid[i+1][0] % mod)) % mod
		}
		for j := m - 2; j >= 0; j-- {
			res[i][j] = (res[i][j+1] * (grid[i][j+1] % mod)) % mod
		}
	}
	pefixM := 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res[i][j] = (pefixM * res[i][j]) % mod
			pefixM = (pefixM * (grid[i][j] % mod)) % mod
		}
	}
	return res
}