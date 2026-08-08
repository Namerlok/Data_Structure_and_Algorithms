func maxProfit(prices []int) int {
	maxProf := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProf += prices[i] - prices[i-1]
		}
	}
	return maxProf
}