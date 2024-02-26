package binary

// 105. Construct Binary Tree from Preorder and Inorder Traversal
// 这一题主要是切片的用法，前提是知道算法本身。
func buildTreeFromPreOrderAndInorder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	rootVal := preorder[0]
	rootIndex := 0
	for idx, val := range inorder {
		if val == rootVal {
			rootIndex = idx
			break
		}
	}

	root := &TreeNode{Val: rootVal}

	leftPreorder := preorder[1 : 1+rootIndex]
	rightPreOrder := preorder[rootIndex+1:]

	leftInorder := inorder[0:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	root.Left = buildTreeFromPreOrderAndInorder(leftPreorder, leftInorder)
	root.Right = buildTreeFromPreOrderAndInorder(rightPreOrder, rightInorder)

	return root
}
