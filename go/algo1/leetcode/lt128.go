package leetcode

func longestConsecutive(nums []int) int {
	bm := make(map[int]bool)
	_min := 0
	_max := 0
	for _, n := range nums {
		if n > _max {
			_max = n
		}
		if n < _min {
			_min = n
		}
		bm[n] = true
	}
	ans := 0
	cnt := 0

	for k, _ := range bm {
		if _, ok := bm[k-1]; ok {
			continue
		}
		for i := k; i <= _max; i++ {
			if _, ok := bm[i]; ok {
				cnt++
			} else {
				break
			}
		}
		ans = max(ans, cnt)
		cnt = 0
	}

	return max(ans, cnt)
}
