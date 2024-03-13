package leetcode

func pow(x, e int) int {
	if e == 0 {
		return 1
	}
	power := 1
	for i := 0; i < e; i++ {
		power = power * x
	}
	return power
}

func titleToNumber(columnTitle string) int {
	var num int
	for i := len(columnTitle) - 1; i >= 0; i-- {
		num += (int)(columnTitle[i]-'A') + i*26
	}
	return num
}
