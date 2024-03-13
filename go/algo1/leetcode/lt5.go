package leetcode

func isPal(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	pal := s[0:1]
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] {
				if isPal(s[i:j+1]) && len(pal) < (j-i+1) {
					pal = s[i : j+1]
				}
			}
		}
	}
	return pal
}
