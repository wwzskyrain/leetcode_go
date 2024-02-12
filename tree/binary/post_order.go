package binary

// 相对于昨天的pre_order，这个题目的解法毫无新意
func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var post func(node *TreeNode)

	post = func(root *TreeNode) {
		if root == nil {
			return
		}
		post(root.Left)
		post(root.Right)
		ret = append(ret, root.Val)
	}
	post(root)

	return ret
}
