package binary

import "sort"

// 2583. Kth Largest Sum in a Binary Tree
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var levelSumArray []int64
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelSum, qSize := int64(0), len(q)
		for i := 0; i < qSize; i++ {
			node := q[0]
			q = q[1:] //就用这种方式来实现队列吧，人家在力扣上都能用
			levelSum += int64(node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		levelSumArray = append(levelSumArray, levelSum)
	}
	if len(levelSumArray) < k {
		return -1
	}
	sort.Slice(levelSumArray, func(i, j int) bool {
		return levelSumArray[i] > levelSumArray[j]
	})
	return levelSumArray[k]
}
