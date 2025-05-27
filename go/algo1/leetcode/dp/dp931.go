package dp

func minFallingPathSum(matrix [][]int) int {
    if len(matrix) <1 ||len(matrix) > 100 {
		return -1
	}
	if len(matrix) == 1 {
		return matrix[0][0]
	}
	dp := make([][]int, len(matrix))
	dp[0] = make([]int, len(matrix))
	for i :=0;i<len(matrix); i++ {
		dp[0][i] = matrix[0][i]
	}

	for i := 1; i < len(matrix);i++ {
		dp[i] = make([]int, len(matrix))
		dp[i][0] = min(dp[i-1][0], dp[i-1][1]) + matrix[i][0]
		for j := 1; j < len(matrix);j++ {
			_min := min(dp[i-1][j], dp[i-1][j-1])
			if j+1 < len(matrix) {
				_min = min(_min, dp[i-1][j+1])
			}
			dp[i][j] = _min + matrix[i][j]
		}
	}
	ans := dp[len(matrix)-1][0]
	for i := 1; i < len(matrix);i++ {
		ans = min(ans, dp[len(matrix)-1][i])
	}
	return ans
}