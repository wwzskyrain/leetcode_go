package binary

import "testing"

func TestCountPaths(t *testing.T) {
	countPaths(5, [][]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}})
}
