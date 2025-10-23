package dp

func maximalSquareNaive(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	maxSquare := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			size := len(matrix) - i
			if size > len(matrix[i])-j {
				size = len(matrix[i]) - j
			}

			if matrix[i][j] == 48 { // ascii 48 - '0'
				continue
			}

			for s := 1; s <= size; s++ {
				isSquare := true
				for ii := 0; ii < s; ii++ {
					for jj := 0; jj < s; jj++ {
						if matrix[i+ii][j+jj] == 48 {
							isSquare = false
							break
						}
					}
					if !isSquare {
						break
					}
				}
				if isSquare {
					if maxSquare < s*s {
						maxSquare = s * s
					}
				} else {
					break
				}
			}
		}
	}
	return maxSquare
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	// init
	dp := make([][]int, len(matrix))
	dp[0] = make([]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
		}
	}
	for i := 1; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		}
	}

	// create dp
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
			} else {
				dp[i][j] = 0
			}
		}
	}
	maxSqual := 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			maxSqual = max(maxSqual, dp[i][j]*dp[i][j])
		}
	}
	return maxSqual
}
