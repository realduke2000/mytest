package dp

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	ans[0] = []int{1}
	for i := 1; i < numRows; i++ {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
		ans[i][i] = 1
	}
	return ans
}
