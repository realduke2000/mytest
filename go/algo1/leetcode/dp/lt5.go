package dp

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	if len(s) == 2 {
		if  s[0] == s[1] {
			return s
		} else {
			return s[:1]
		}
	}
	
	// initialization
	dp := make([][]int, len(s))
	for i :=0;i< len(s);i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	start := 0
	end := 0
	// optimized answers
	for L := 2; L <= len(s); L++ {
		for i := 0; i < len(s); i++ {
			j := i + L - 1
			if j >= len(s) {
				break
			}

			if s[i] == s[j] {
				if i+1 == j || dp[i+1][j-1] == 1 {
					dp[i][j] = 1
					if j - i > end - start {
						start = i
						end = j
					}
					
				}
			}
		}
	}
	return s[start:end+1]


	// // find answer
	// ans := ""
	// for i:=0;i<len(s);i++ {
	// 	for j :=0;j<len(s);j++ {
	// 		if dp[i][j] != 0 && len(ans) < j-i+1 {
	// 			ans = s[i:j+1]
	// 		}
	// 	}
	// }
	// return ans
}