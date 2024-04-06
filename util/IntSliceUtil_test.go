package util

import (
	"fmt"
	"testing"
)

func TestBuildIntSlice2Dimension(t *testing.T) {
	Slice2D := BuildIntSlice2Dimension("[[0,3],[0,4],[1,3],[2,4],[2,7],[3,5],[3,6],[3,7],[4,6]]")
	for _, s := range Slice2D {
		fmt.Printf("s = %v\n", s)
	}
}
