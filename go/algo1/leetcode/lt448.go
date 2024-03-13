package leetcode

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		for x > len(nums)+1 {
			x = x - len(nums) - 1
		}
		nums[x-1] += len(nums) + 1
	}
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] < len(nums)+1 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
