package leetcode

func isPalindrome125(s string) bool {
	validated := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] <= '9' && s[i] >= '0' {
			validated = append(validated, s[i])
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			validated = append(validated, s[i]-'A'+'a')
		}
	}

	n := len(validated)
	if n <= 1 {
		return true
	}
	if n == 2 {
		return validated[0] == validated[1]
	}

	for i := 0; i < n/2; i++ {
		if validated[i] != validated[n-1-i] {
			return false
		}
	}
	return true
}


func charDelta(b byte) byte {
	if b >= 'a' && b <= 'z' || b <= '9' && b >= '0' {
		return b
	}
	if b >= 'A' && b <= 'Z' {
		return b-'A'+'a'
	}
	return 0
}

func isPalindrome2Pointers(s string) bool {
	if len(s) <=1 {
		return true
	}
	if len(s) == 2 {
		if charDelta(s[0]) == 0 || charDelta(s[1]) == 0 {
			return true 
		} else {
			return charDelta(s[0]) == charDelta(s[1])
		}
	}	
	left := 0
	right := len(s)-1
	for left < right {
		if charDelta(s[left]) == 0 {
			left++
			continue
		}
		if charDelta(s[right]) == 0 {
			right--
			continue
		}
		if charDelta(s[left]) != charDelta(s[right]) {
			return  false
		}
		left++
		right--
	}
	return true
}