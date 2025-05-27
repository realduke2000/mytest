package presum

func subarraySum(nums []int, k int) int {
	ans := 0
	if len(nums) == 0 {
		return 0
	}
	mp := make(map[int]int)
	mp[0] = 1
	pre := 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if c, ok := mp[pre-k]; ok {
			ans += c
		}
		if v, ok := mp[pre]; ok {
			mp[pre] = v + 1
		} else {
			mp[pre] = 1
		}
	}
	return ans
}
