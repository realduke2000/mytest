package dp

func maxProduct(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	_max := make([]int, len(nums))
	_max[0] = nums[0]
	_min := make([]int, len(nums))
	_min[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		_max[i] = max(nums[i], _max[i-1]*nums[i], _min[i-1]*nums[i])
		_min[i] = min(nums[i], _max[i-1]*nums[i], _min[i-1]*nums[i])
		dp[i] = max(dp[i-1], _max[i])

	}
	return dp[len(nums)-1]
}
