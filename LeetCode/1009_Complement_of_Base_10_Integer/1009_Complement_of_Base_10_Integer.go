func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	var result int = 1
	for result < n {
		result *= 2
	}
	return ((result - 1) ^ n) % n
}