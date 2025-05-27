package presum

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	pre := make([]int, len(nums))
	post := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = 1
		pre[i] = 1
		post[i] = 1
	}

	for i, _ := range nums {
		if i != 0 {
			pre[i] = nums[i-1] * pre[i-1]
			post[len(nums)-1-i] = post[len(nums)-i] * nums[len(nums)-i]
		}
	}
	for i, _ := range nums {
		ans[i] = pre[i] * post[i]
	}
	return ans
}
