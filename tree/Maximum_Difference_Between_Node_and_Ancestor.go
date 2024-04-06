package tree

func maxAncestorDiff(root *TreeNode) int {

	var dfs func(root *TreeNode, max, min int) int
	dfs = func(root *TreeNode, max, min int) int {
		if root == nil {
			return 0
		}
		diff := MathMax(MathAbs(root.Val-max), MathAbs(root.Val-min))
		max = MathMax(max, root.Val)
		min = MathMin(min, root.Val)
		diff = MathMax(diff, dfs(root.Left, max, min))
		diff = MathMax(diff, dfs(root.Right, max, min))
		return diff
	}
	return dfs(root, root.Val, root.Val)
}

func MathMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MathMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MathAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
