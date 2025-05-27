package leetcode

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	ans := nums[0]
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < 0 {
			if ans < sum {
				ans = sum
			}
			sum = 0
		} else {
			if ans < sum {
				ans = sum
			}
		}
	}
	return ans
}
