package leetcode

import "fmt"

func maxProfitExhausitive(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit = max(prices[j]-prices[i], profit)
		}
	}
	return profit
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// [1,4,10]
	// [7,1,5,3,6,4]
	// max profie if sell on dp[i] day
	// dp[0] = 0
	// dp[1] = prices[1]-prices[0]
	// dp[i] = max(dp[i-1]+prices[i]-prices[i-1], prices[i]-prices[i-1])
	dp := make([]int, len(prices))
	dp[0] = 0
	dp[1] = prices[1] - prices[0]
	profit := max(dp[0], dp[1])
	for i := 2; i < len(prices); i++ {
		dp[i] = max(dp[i-1]+prices[i]-prices[i-1], prices[i]-prices[i-1])
		profit = max(profit, dp[i])
	}
	fmt.Printf("%v\n", dp)
	return profit
}
