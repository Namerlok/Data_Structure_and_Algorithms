func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		b, a = a+b, b
	}
	return b
}