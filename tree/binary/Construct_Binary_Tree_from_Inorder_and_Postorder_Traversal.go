package binary

// 106. Construct Binary Tree from Inorder and Postorder Traversal
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	if len(postorder) == 1 {
		return &TreeNode{Val: postorder[0]}
	}
	rootVal := postorder[len(postorder)-1]
	rootIndex := 0
	for idx, val := range inorder {
		if val == rootVal {
			rootIndex = idx
			break
		}
	}

	root := &TreeNode{Val: rootVal}

	leftPostorder := postorder[0:rootIndex]
	rightPostorder := postorder[rootIndex : len(postorder)-1]

	leftInorder := inorder[0:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)

	return root
}
