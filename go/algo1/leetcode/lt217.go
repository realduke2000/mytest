package leetcode

func containsDuplicate(nums []int) bool {
	set := make(map[int]int)
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		} else {
			set[v] = 0
		}
	}
	return false
}

// lt219
func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]int)
	for i, v := range nums {
		if j, ok := set[v]; ok {
			diff := i - j
			if diff < 0 {
				diff = j - i
			}
			if diff <= k {
				return true
			}
			set[v] = j
		} else {
			set[v] = i
		}
	}
	return false
}
