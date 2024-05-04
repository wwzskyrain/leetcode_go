package tree

import (
	"fmt"
	"testing"
)
import "study.go/leetcode/util"

func TestGetCoprimes(t *testing.T) {
	//	nums = [2,3,3,2], edges = [[0,1],[1,2],[1,3]]

	edges := util.BuildIntSlice2Dimension("[[0,1],[1,2],[1,3]]")
	coprimes := getCoprimes([]int{2, 3, 3, 2}, edges)
	fmt.Printf("coprimes=%v", coprimes)

}
