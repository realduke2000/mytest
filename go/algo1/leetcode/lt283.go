package leetcode

func moveZeroes(nums []int) {
	left := 0
	for left < len(nums) && nums[left] != 0 {
		left++
	}

	right := left + 1
	for right < len(nums) && nums[right] == 0 {
		right++
	}

	for left < right && left < len(nums) && right < len(nums) {
		tmp := nums[left]
		nums[left] = nums[right]
		nums[right] = tmp

		for left < len(nums) && nums[left] != 0 {
			left++
		}
		right = left + 1
		for right < len(nums) && nums[right] == 0 {
			right++
		}
	}

}
