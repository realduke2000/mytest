package dp

import "math"

func numSquares(n int) int {
	if n <= 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		minTmp := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minTmp = min(dp[i-j*j]+1, minTmp)
		}
		dp[i] = min(dp[i-1]+1, minTmp)
	}
	return dp[n]
}
