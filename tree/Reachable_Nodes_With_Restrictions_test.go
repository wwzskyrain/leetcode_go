package tree

import "testing"

func TestReachableNodes(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}
	restricted := []int{4, 5}
	reachableNodes(7, edges, restricted)
}
