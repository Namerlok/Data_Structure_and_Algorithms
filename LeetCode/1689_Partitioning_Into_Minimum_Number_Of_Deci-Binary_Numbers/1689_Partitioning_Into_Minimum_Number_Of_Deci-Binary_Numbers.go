func max(left, right int) int {
	if left > right {
		return left
	} else {
		return right
	}
}

func minPartitions(n string) int {
	var res int
	for _, el := range n {
		res = max(res, int(el)-int('0'))
	}
	return res
}