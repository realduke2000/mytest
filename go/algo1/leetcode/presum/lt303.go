package presum

type NumArray struct {
	presum []int
}

func Constructor(nums []int) NumArray {
	var na NumArray
	na.presum = make([]int, len(nums))
	for i, n := range nums {
		if i == 0 {
			na.presum[i] = n
		} else {
			na.presum[i] = na.presum[i-1] + n
		}
	}
	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.presum[right]
	}
	return (this.presum[right] - this.presum[left]) + (this.presum[left] - this.presum[left-1])
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
