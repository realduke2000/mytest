package dp

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// dp[i] max profie until day i
	// dp[0] = 0
	// dp[1] = max(dp[0], prices[1]-prices[0])
	// dp[2] = max(dp[1], prices[2]-prices[1])
	// dp[i] = max(dp[i-1], prices[i]-prices[i-1])
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		dp[i] = max(dp[i-1], prices[i]-prices[i-1]+dp[i-1])
	}
	return dp[len(dp)-1]
}
