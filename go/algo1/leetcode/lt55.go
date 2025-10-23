package leetcode

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	if len(nums) == 1 {
		return nums[0] >= 0
	}

	if nums[0] == 0 {
		return false
	}

	/*
		dp[0] = nums[0]
		dp[i] = max(dp[i-1] - 1, nums[i-1]-1)
	*/
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]-1, nums[i-1]-1)
		if dp[i] < 0 {
			return false
		}
	}
	return true
}
