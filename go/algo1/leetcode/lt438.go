package leetcode

import "encoding/hex"

func hashAnagrams438(s string) string {
	hash := make([]byte, 256)
	for i, _ := range hash {
		hash[i] = 0
	}
	for _, r := range s {
		hash[r] = hash[r] + 1
	}
	return hex.EncodeToString(hash)
}

func findAnagrams(s string, p string) []int {
	if s == "" || p == "" || len(s) < len(p) {
		return []int{}
	}
	if s == p {
		return []int{0}
	}
	if len(s) == len(p) && s[0] != p[0] {
		return []int{}
	}

	ans := []int{}
	hash := hashAnagrams438(p)
	for i := 0; i <= len(s)-len(p); i++ {
		_h := hashAnagrams438(s[i : i+len(p)])
		if _h == hash {
			ans = append(ans, i)
		}
	}
	return ans
}
