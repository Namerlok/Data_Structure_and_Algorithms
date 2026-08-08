func findKthBit(n int, k int) byte {
	length := 0
	for i := 0; i < n; i++ {
		length = length*2 + 1
	}

	invert := false
	bitIsOne := false
	for i := 1; i < n; i++ {
		mid := length/2 + 1
		if k == mid {
			bitIsOne = true
			break
		}
		if k > mid {
			k = length - k + 1
			invert = !invert
		}
		length /= 2
	}

	if invert {
		bitIsOne = !bitIsOne
	}
	if bitIsOne {
		return '1'
	}
	return '0'
}