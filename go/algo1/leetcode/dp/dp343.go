package dp

func integerBreak(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	dp := make(map[int]int, n+1)
	dp[0] = 0
	dp[1] = 0
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		dp[i] = 0
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], j*(i-j), dp[j]*(i-j), j*dp[i-j], dp[j]*dp[i-j])
		}
	}
	return dp[n]
}
