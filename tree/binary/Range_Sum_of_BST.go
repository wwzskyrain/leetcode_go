package binary

// 938. Range Sum of BST
func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val >= low && root.Val <= high {
			sum += root.Val
		}
		inorder(root.Left)
		inorder(root.Right)
	}
	inorder(root)
	return sum
}
