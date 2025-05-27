package dp

func rob740(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	return max(dp[len(nums)-1], dp[len(nums)-2])
}

func deleteAndEarn(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	max_num := nums[0]
	for i := 0; i < len(nums); i++ {
		max_num = max(nums[i], max_num)
	}

	sum := make([]int, max_num+1)
	for i := 0; i < len(nums); i++ {
		sum[nums[i]] = sum[nums[i]] + nums[i]
	}
	return rob740(sum)
}
