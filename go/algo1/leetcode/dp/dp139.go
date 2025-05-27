package dp

import "strings"

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1);
	dp[0]=true
	for i := 1; i <= len(s); i++ {
		for j :=0;j<len(wordDict);j++ {
			w := wordDict[j]
			if strings.HasSuffix(s[:i], w) && dp[i-len(w)] {
				dp[i] = true
				break
			} 
		}
	}
	return dp[len(dp)-1]
}
