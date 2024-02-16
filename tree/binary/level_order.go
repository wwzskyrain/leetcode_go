package binary

// 145. 二叉树的后序遍历
// 这是我们第一用slice来完成队列操作，并借此实现了二叉树的层次遍历
func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	front := 0
	for front < len(q) { // q is not empty
		qLen := len(q)
		curLevel := make([]int, 0)
		for front < qLen {
			// 访问这一层
			node := q[front]
			front++
			curLevel = append(curLevel, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		// 这一层访问完了
		ret = append(ret, curLevel)
	}

	return ret
}
