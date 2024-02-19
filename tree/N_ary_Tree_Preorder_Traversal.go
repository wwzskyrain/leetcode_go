package tree

// 589. N-ary Tree Preorder Traversal
// 和二叉树的递归遍历，没什么不同。又复习了一遍而已
func preorder(root *Node) []int {

	s := make([]int, 0)

	var preTraversal func(root *Node)
	preTraversal = func(root *Node) {
		if root == nil {
			return
		}
		s = append(s, root.Val)
		for _, child := range root.Children {
			preTraversal(child)
		}
	}
	preTraversal(root)

	return s
}
