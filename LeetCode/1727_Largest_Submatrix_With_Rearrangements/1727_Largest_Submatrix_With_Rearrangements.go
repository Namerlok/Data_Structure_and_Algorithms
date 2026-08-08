type Square struct {
	next   *Square
	col    int
	height int
}

func largestSubmatrix(matrix [][]int) int {
	counterOneInRow := make([][]int, 2)
	counterOneInRow[0] = make([]int, len(matrix[0]))
	counterOneInRow[1] = make([]int, len(matrix[0]))
	maxSquare := 0
	var square *Square
	lenSquare := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				counterOneInRow[i%2][j] = counterOneInRow[(i+1)%2][j] + 1
			} else {
				counterOneInRow[i%2][j] = 0
			}
			if counterOneInRow[i%2][j] == 1 {
				square = &Square{
					next:   square,
					col:    j,
					height: 1,
				}
				lenSquare++
			}
		}
		cur := square
		var prev *Square
		for cur != nil {
			if counterOneInRow[i%2][cur.col] == 0 {
				if prev == nil {
					square = cur.next
				} else {
					prev.next = cur.next
				}
				lenSquare--
			} else {
				cur.height = counterOneInRow[i%2][cur.col]
				prev = cur
			}
			cur = cur.next
		}
		cur = square
		for k := 0; cur != nil; k++ {
			if cur.height*(lenSquare-k) > maxSquare {
				maxSquare = cur.height * (lenSquare - k)
			}
			cur = cur.next
		}
	}
	return maxSquare
}