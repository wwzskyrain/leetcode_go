package grid

import (
	"fmt"
	"testing"
)

func TestMaxMoves(t *testing.T) {
	////maxMoves(grid [][]int)
	var grid = [][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}
	ret := maxMoves(grid)
	fmt.Printf("%d", ret)

	//[[3,2,4],[2,1,9],[1,1,7]]
	var grid2 = [][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}
	fmt.Printf("%d", maxMoves(grid2))
}
