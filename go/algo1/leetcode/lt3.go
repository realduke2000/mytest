package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	left := 0
	right := 1
	maxLen := 0
	bitmap := make(map[uint8]int)
	bitmap[s[left]] = left

	for right < len(s) && left < right {
		if dupIndex, ok := bitmap[s[right]]; ok {
			for i := left; i <= dupIndex; i++ {
				delete(bitmap, s[i])
			}
			maxLen = max(maxLen, right-left)
			left = dupIndex + 1

			if left >= right {
				bitmap[s[left]] = left
				right = left + 1
			}
		} else {
			bitmap[s[right]] = right
			right++
		}
	}
	return max(maxLen, right-left)
}
