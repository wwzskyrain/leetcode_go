package binary

// 103. 二叉树的锯齿形层序遍历
// 还是二叉树，一会从左到右，一会从右到左。还是队列使用
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var nodeLevels [][]*TreeNode
	level := 0
	nodeLevels = append(nodeLevels, []*TreeNode{root})
	for level < len(nodeLevels) {
		// 前一层产生next层
		curLevel := nodeLevels[level]
		var nextLevel []*TreeNode
		for _, node := range curLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		level++
		if len(nextLevel) == 0 {
			continue
		}
		nodeLevels = append(nodeLevels, nextLevel)
	}

	var ret [][]int
	for l, nodeLevel := range nodeLevels {
		var nodeValLevel []int
		for _, node := range nodeLevel {
			nodeValLevel = append(nodeValLevel, node.Val)
		}
		if l%2 == 1 {
			reverseIntSlice(nodeValLevel)
		}
		ret = append(ret, nodeValLevel)
	}
	return ret
}

func reverseIntSlice(s []int) []int {
	i, j := 0, len(s)-1
	var temp int
	for i < j {
		temp = s[i]
		s[i] = s[j]
		s[j] = temp
		i++
		j--
	}
	return s
}
