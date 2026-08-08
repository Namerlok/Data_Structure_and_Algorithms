func productNumbers(n int) int {
	prod := 1
	for n != 0 {
		prod *= n % 10
		n /= 10
	}
	return prod
}

func smallestNumber(n int, t int) int {
	for productNumbers(n)%t != 0 {
		n++
	}
	return n
}