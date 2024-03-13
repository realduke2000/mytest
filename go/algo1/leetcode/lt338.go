package leetcode

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = count1(i)
	}
	return ans
}

func count1(n int) int {
	i := 1
	cnt := 0
	for i <= n {
		if n&i == i {
			cnt++
		}
		i = i << 1
	}
	return cnt
}
