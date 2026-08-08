func pow(num, pow int) int {
	result := 1
	for i := 0; i < pow; i++ {
		result *= num
	}
	return result
}

func getHappyString(n int, k int) string {
	blockSize := 3 * pow(2, n-1)
	if blockSize < k {
		return ""
	}
	result := make([]rune, n)
	blockSize /= 3
	if blockSize >= k {
		result[0] = 'a'
	} else if blockSize*2 >= k {
		result[0] = 'b'
		k -= blockSize
	} else {
		result[0] = 'c'
		k -= blockSize * 2
	}

	for pos := 1; pos < n; pos++ {
		blockSize /= 2
		useFirstOption := blockSize >= k
		if !useFirstOption {
			k -= blockSize
		}

		isFirstOption := true
		for _, choice := range []rune{'a', 'b', 'c'} {
			if result[pos-1] != choice {
				if useFirstOption == isFirstOption {
					result[pos] = choice
					break
				} else {
					isFirstOption = false
				}
			}
		}
		fmt.Println(isFirstOption)
	}
	return string(result)
}