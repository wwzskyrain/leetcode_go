package array

// 303. Range Sum Query - Immutable
type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	var preSum []int
	for i, num := range nums {
		s := num
		if i > 0 {
			s += preSum[i-1]
		}
		preSum = append(preSum, s)
	}
	return NumArray{preSum: preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > 0 {
		return this.preSum[right] - this.preSum[left-1]
	} else {
		return this.preSum[right]
	}
}
