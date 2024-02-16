package binary

import "testing"

func TestLevelOrder(t *testing.T) {
	//root := BuildBinaryTree("[3,9,20,null,null,15,7]")
	//levelOrder(root)

	//root := BuildBinaryTree("[1]")
	//levelOrder(root)

	root := BuildBinaryTree("[ ]")
	levelOrder(root)
}
