package binary

// 107. 二叉树的层序遍历 II
// 在第一题的基础上，把结果slice来一个对调就OK
func levelOrderBottom(root *TreeNode) [][]int {
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

	i, j := 0, len(ret)-1
	for i < j {
		temp := ret[i]
		ret[i] = ret[j]
		ret[j] = temp
		i++
		j--
	}
	return ret
}
