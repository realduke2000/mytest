package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	p := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			p = i + 1
			break
		}
	}

	if p == -1 {
		return len(nums)
	}

	for i := p; i < len(nums); {
		j := i + 1
		for ; j < len(nums); j++ {
			if nums[i] != nums[j] {
				nums[p] = nums[j]
				p++
				i = j
				break
			}
		}
		if j == len(nums) {
			break
		}
	}
	return p
}
