package leetcode

import "sort"

type lt56Interval [][]int

func (p lt56Interval) Len() int {
	return len(p)
}
func (p lt56Interval) Less(i int, j int) bool {
	return p[i][0] < p[j][0]
}
func (p lt56Interval) Swap(i int, j int) {
	tmps := p[i][0]
	tmpe := p[i][1]
	p[i][0] = p[j][0]
	p[i][1] = p[j][1]
	p[j][0] = tmps
	p[j][1] = tmpe
}

func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return nil
	}
	if len(intervals) == 1 {
		return intervals
	}
	sort.Sort(lt56Interval(intervals))
	ans := [][]int{
		{intervals[0][0], intervals[0][1]},
	}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= ans[len(ans)-1][1] && intervals[i][0] >= ans[len(ans)-1][0] {
			ans[len(ans)-1][1] = max(intervals[i][1], ans[len(ans)-1][1])
		} else {
			ans = append(ans, []int{intervals[i][0], intervals[i][1]})
		}
	}
	return ans
}
