package leetcode

func romanMap(r rune) int {
	switch r {
	case 'I': return 1
	case 'V': return 5
	case 'X': return 10
	case 'L': return 50
	case 'C': return 100
	case 'D': return 500
	case 'M': return 1000
	default: return 0
	}
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	rMap := map[byte]int {'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	num := 0
	for i := range s {
		if v, ok := rMap[s[i]]; !ok {
			return -1
		} else {
			if i < len(s) -1 && v < rMap[s[i+1]] {
				num -= v
			} else {
				num += v
			}
		}
	}
	return num
}
