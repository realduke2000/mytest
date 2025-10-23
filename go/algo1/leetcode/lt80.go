package leetcode

func removeDuplicates80(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	// find 1st 3rd same numbers
	p := -1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				if j-i > 1 {
					p = j
					break
				}
			} else {
				break
			}
		}
		if p != -1 {
			break
		}
	}
	if p == -1 {
		return len(nums)
	}

	/*
		1	2
	*/
	cnt := 2
	pivot := nums[p]
	for i := p + 1; i < len(nums); i++ {
		if nums[i] != pivot {
			nums[p] = nums[i]
			p++
			cnt = 1
			pivot = nums[i]
		} else {
			if cnt < 2 {
				nums[p] = pivot
				p++
			}
			cnt++
		}
	}
	return p
}

func removeDuplicatesStack(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	datap := 2
	top := 2

	for datap < len(nums) {
		if nums[datap] == nums[top-1] {
			if nums[datap] == nums[top-2] {
				datap++
				continue
			} else {
				nums[top] = nums[datap]
				top++
				datap++
			}
		} else {
			nums[top] = nums[datap]
			datap++
			top++
		}
	}
	return top
}
