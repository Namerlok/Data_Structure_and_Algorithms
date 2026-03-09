package main

func min(left, right int) int {
	if left < right {
		return left
	} else {
		return right
	}
}

func numberOfStableArrays(zero int, one int, limit int) int {
	dp := make([][][2]int, zero+1)
	for i := 0; i <= zero; i++ {
		dp[i] = make([][2]int, one+1)
	}
	for i := 1; i <= min(zero, limit); i++ {
		dp[i][0][0] = 1
	}
	for j := 1; j <= min(one, limit); j++ {
		dp[0][j][1] = 1
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			dp[i][j][0] = dp[i-1][j][1] + dp[i-1][j][0]
			if i > limit {
				dp[i][j][0] -= dp[i-1-limit][j][1]
			}
			dp[i][j][0] = (dp[i][j][0]%int(1e9+7) + int(1e9+7)) % int(1e9+7)

			dp[i][j][1] = dp[i][j-1][0] + dp[i][j-1][1]
			if j > limit {
				dp[i][j][1] -= dp[i][j-1-limit][0]
			}
			dp[i][j][1] = (dp[i][j][1]%int(1e9+7) + int(1e9+7)) % int(1e9+7)
		}
	}

	return (dp[zero][one][0] + dp[zero][one][1]) % int(1e9+7)
}
