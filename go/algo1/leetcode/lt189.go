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

func rotate(nums []int, k int) {
	if k <= 0 || k%len(nums) == 0 || nums == nil || len(nums) <= 1 {
		return
	}
	k = k % len(nums)
	reverse189(nums)
	reverse189(nums[:k])
	reverse189(nums[k:])
}
