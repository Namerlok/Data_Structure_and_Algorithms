func canPartitionSlice(sl []int) bool {
	var sumL, sumR int
	posL, posR := 0, len(sl)-1
	for posL <= posR {
		if sumL < sumR {
			sumL += sl[posL]
			posL++
		} else {
			sumR += sl[posR]
			posR--
		}
	}
	if sumL != 0 && sumL == sumR {
		return true
	}
	return false
}

func canPartitionGrid(grid [][]int) bool {
	row := make([]int, len(grid[0]))
	line := make([]int, len(grid))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			line[i] += grid[i][j]
			row[j] += grid[i][j]
		}
	}
	return canPartitionSlice(line) || canPartitionSlice(row)
}