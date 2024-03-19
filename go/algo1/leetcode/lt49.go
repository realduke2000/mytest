package leetcode

import "encoding/hex"

func hashAnagrams49(s string) string {
	hash := make([]byte, 256)
	for i, _ := range hash {
		hash[i] = 0
	}
	for _, r := range s {
		hash[r]++
	}
	return hex.EncodeToString(hash)
}

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string, 0)
	for _, s := range strs {
		k := hashAnagrams49(s)
		if _, ok := hash[k]; ok {
			hash[k] = append(hash[k], s)
		} else {
			hash[k] = []string{s}
		}
	}
	var ans [][]string
	for _, v := range hash {
		ans = append(ans, v)
	}
	return ans
}
