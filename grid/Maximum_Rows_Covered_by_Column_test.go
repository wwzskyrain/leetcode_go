package grid

import (
	"fmt"
	"testing"
)

// 2397. Maximum Rows Covered by Columns
func TestMaximumRows(t *testing.T) {
	matrix := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}
	numSelect := 1
	res := maximumRows(matrix, numSelect)
	fmt.Println(res)
}
