package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return make([][]int, 0)
	}

	sort.Ints(nums)
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return make([][]int, 0)
	}

	ans := make([][]int, 0)
	k := 0
	left := k + 1
	right := len(nums) - 1

	for k < len(nums) && nums[k] <= 0 {
		if left >= right {
			k++
			left = k + 1
			right = len(nums) - 1
			continue
		}
		if nums[k]+nums[left]+nums[right] > 0 {
			right--
		} else if nums[k]+nums[left]+nums[right] < 0 {
			left++
		} else {
			dup := false
			for _, a1 := range ans {
				if a1[0] == nums[k] && a1[1] == nums[left] && a1[2] == nums[right] {
					dup = true
					break
				}
			}
			if !dup {
				ans = append(ans, []int{nums[k], nums[left], nums[right]})
			}
			left++
			right--
		}
	}
	return ans
}
