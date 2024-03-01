package medium

import "testing"

func TestValidPartition(t *testing.T) {
	println(validPartition([]int{4, 4, 4, 5, 6}))
	println(validPartition([]int{4, 4, 5, 6}))
	println(validPartition([]int{1, 1, 1, 2}))
}
