package leetcode

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 0
	p := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			k++
			p = i
			break
		}
	}

	if p == -1 {
		return len(nums) - k
	}

	for i := p + 1; i < len(nums); i++ {
		if nums[i] != val {
			nums[p] = nums[i]
			p++
		} else {
			k++
		}
	}
	return len(nums) - k
}
