package leetcode

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, n := range nums {
		mp[n] = i
	}
	for i, n := range nums {
		if j, ok := mp[target-n]; ok && j != i {
			return []int{mp[target-n], i}
		}
	}
	return nil
}
