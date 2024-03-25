package medium

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	println(coinChange([]int{1, 2, 5}, 11))
	println(coinChange([]int{2}, 3))
	println(coinChange([]int{1}, 0))
}
