package binary

import "testing"

func TestZigzagLevelOrder(t *testing.T) {
	root := BuildBinaryTree("[3,9,20,null,null,15,7]")
	zigzagLevelOrder(root)
}
