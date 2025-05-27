package dp

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 || m ==1 || n == 1 {
		return 1
	}
	dp := make([][]int, n)

	// initialize
	for i :=0; i < n; i++ {
		dp[i] = make([]int, m)
		if i == 0 {
			dp[0][0] = 0
			for j := 1; j < m; j++ {
				dp[0][j] = 1
			}
		} else {
			dp[i][0] = 1
			for j := 1; j < m; j++ {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[n-1][m-1]
}