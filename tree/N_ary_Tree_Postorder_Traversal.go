package tree

// 590. N-ary Tree Postorder Traversal
// 一并写了算了
func postorder(root *Node) []int {
	s := make([]int, 0)

	var postTraversal func(root *Node)
	postTraversal = func(root *Node) {
		if root == nil {
			return
		}
		for _, child := range root.Children {
			postTraversal(child)
		}
		s = append(s, root.Val)
	}
	postTraversal(root)

	return s
}
