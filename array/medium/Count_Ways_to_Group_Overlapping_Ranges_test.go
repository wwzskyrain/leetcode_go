package medium

import (
	"fmt"
	"testing"
)

func Test_countWays(t *testing.T) {

	fmt.Println(countWays([][]int{{0, 0}, {8, 9}, {12, 13}, {1, 3}}))
	fmt.Println(countWays([][]int{{6, 10}, {5, 15}}))
}
