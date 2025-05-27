package dp


func mymin(arr []int) int {
	r := arr[0]
	for i := 0; i < len(arr);i++ {
		if arr[i] < r {
			r = arr[i]
		}
	}
	return r
}

/*
2
3 4
6 5 7
4 1 8 3
*/

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	dp := make([][]int, len(triangle))
	dp[0] = []int{triangle[0][0]}

	for i := 1; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
		dp[i][0] = dp[i-1][0]+triangle[i][0]
		for j:=1;j<len(triangle[i]);j++ {
			if len(dp[i-1]) > j {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			} else {
				dp[i][j] = dp[i-1][j-1]+triangle[i][j]
			}
			
		}
	}
	r := mymin(dp[len(dp)-1])
	return r
}
