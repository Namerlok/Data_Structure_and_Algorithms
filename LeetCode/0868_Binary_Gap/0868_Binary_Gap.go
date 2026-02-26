package main

func MaxInt(left, right int) int {
	if left < right {
		return right
	} else {
		return left
	}
}

func binaryGap(n int) int {
	maxWay, way := 0, 0
	for n&1 == 0 && n > 0 {
		n >>= 1
	}
	for n > 0 {
		if n&1 == 1 && way != 0 {
			maxWay = MaxInt(maxWay, way)
			way = 0
		}
		way++
		n >>= 1
	}
	return maxWay
}
