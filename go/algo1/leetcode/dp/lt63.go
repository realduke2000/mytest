package dp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	
	if n == 0 {
		return 0
	}
	if n == 1 {
		for j :=0;j<len(obstacleGrid[0]);j++ {
			if obstacleGrid[0][j] == 1 {
				return 0
			}
		}
		return 1
	}
	
	m := len(obstacleGrid[0])
	if m == 0 {
		return 0
	}
	if m == 1 {
		for j :=0;j<len(obstacleGrid);j++ {
			if obstacleGrid[j][0] == 1 {
				return 0
			}
		}
		return 1
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	// initialize
	dp := make([][]int, n)
	dp[0] = make([]int, m)
	dp[0][0] = 1
	// 1st row
	blocked := false
	for j := 1; j < m; j++ {
		if obstacleGrid[0][j] == 1 {
			blocked = true
			dp[0][j]=0
		}
		if blocked {
			dp[0][j] = 0
		} else {
			dp[0][j] = 1
		}
	}

    // row 2~n
	for i :=1; i < n; i++ {
		dp[i] = make([]int, m)
		if obstacleGrid[i-1][0] == 1 {
			dp[i][0] = 0
		} else {
			dp[i][0]=dp[i-1][0]
		}
		
		for j := 1; j < m; j++ {
			upmax := dp[i-1][j]
			if obstacleGrid[i-1][j] == 1 {
				upmax = 0
			}
			leftmax := dp[i][j-1]
			if obstacleGrid[i][j-1] == 1 {
				leftmax = 0
			}
			dp[i][j] = upmax + leftmax
		}
	}
	displayIntArr2(dp)
	if obstacleGrid[n-1][m-1] == 1 {
		return 0
	} else {
		return dp[n-1][m-1]
	}
}