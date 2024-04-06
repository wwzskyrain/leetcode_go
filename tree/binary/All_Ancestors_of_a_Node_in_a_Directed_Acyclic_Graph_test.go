package binary

import (
	"fmt"
	"study.go/leetcode/util"
	"testing"
)

func TestGetAncestors(t *testing.T) {
	Slice2Dimension := util.BuildIntSlice2Dimension("[[0,3],[0,4],[1,3],[2,4],[2,7],[3,5],[3,6],[3,7],[4,6]]")
	ancestors := getAncestors(8, Slice2Dimension)
	fmt.Printf("ancestors = %v \n", ancestors)
}
