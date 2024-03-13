package leetcode

func calcArea(x1, x2, y1, y2 int) int {
	l := x1 - x2
	if l < 0 {
		l = l * -1
	}
	return l * min(y1, y2)
}

func maxArea(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	if len(height) == 1 {
		return height[0]
	}
	left := 0
	right := len(height) - 1
	ans := 0
	for left < right {
		area := calcArea(left, right, height[left], height[right])
		ans = max(area, ans)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}
