func findDifferentBinaryString(nums []string) string {
	numSet := make(map[int]struct{}, len(nums))
	var len int = len(nums[0])
	for _, numStr := range nums {
		if numInt, err := strconv.Atoi(numStr); err == nil {
			numSet[numInt] = struct{}{}
		}
	}
	var res int
	for it := 0; it < 2*len; it++ {
		if _, ok := numSet[res]; !ok {
			break
		}
		if it%2 == 1 {
			res++
		} else {
			res *= 10
		}
	}
	resStr := make([]byte, len)
	for it := len - 1; it >= 0; it-- {
		resStr[it] = byte(res%10) + '0'
		res /= 10
	}
	return string(resStr)
}