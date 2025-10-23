package leetcode

func reverse189(nums []int) {
	start := 0
	end := len(nums) - 1
	for start <= end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}
}
