func min(left, right int) int {
	if left < right {
		return left
	} else {
		return right
	}
}

func minOperations(s string) int {
	var resZero, resOne int
	for pos := 0; pos < len(s); pos++ {
		if int(s[pos]-'0') == pos%2 {
			resOne++
		} else {
			resZero++
		}
	}
	return min(resOne, resZero)
}