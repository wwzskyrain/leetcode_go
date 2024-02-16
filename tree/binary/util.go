package binary

import (
	"strconv"
	"strings"
)

func BuildBinaryTree(input string) *TreeNode {
	input = strings.Replace(input, "[", "", 1)
	input = strings.Replace(input, "]", "", 1)
	if len(strings.Trim(input, " ")) == 0 {
		return nil
	}
	split := strings.Split(input, ",")
	if len(split) == 0 {
		return nil
	}
	//	队列的使用
	q := make([]*TreeNode, 0)
	front := -1

	var root TreeNode
	for i := 0; i < len(split); {
		if i == 0 {
			v, _ := strconv.Atoi(split[i])
			i++
			root = TreeNode{Val: v}
			q = append(q, &root)
			front++
		} else {
			frontNode := q[front]
			front++
			if i < len(split) {
				s := split[i]
				i++
				if s != "null" {
					intVal, _ := strconv.Atoi(s)
					frontNode.Left = &TreeNode{Val: intVal}
					q = append(q, frontNode.Left)
				}
			}
			if i < len(split) {
				s := split[i]
				i++
				if s != "null" {
					intVal, _ := strconv.Atoi(s)
					frontNode.Right = &TreeNode{Val: intVal}
					q = append(q, frontNode.Right)
				}
			}
		}
	}

	return &root
}
