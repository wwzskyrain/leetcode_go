package tree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
//429. N 叉树的层序遍历
//level traversal again. practice q-data_struct again
func levelOrder(root *Node) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	var q []*Node
	front := 0
	q = append(q, root)
	for front < len(q) {
		curLen := len(q)
		var curLevelVal []int
		for front < curLen {
			n := q[front]
			front++
			curLevelVal = append(curLevelVal, n.Val)
			for _, child := range n.Children {
				q = append(q, child)
			}
		}
		ret = append(ret, curLevelVal)
	}
	return ret
}
