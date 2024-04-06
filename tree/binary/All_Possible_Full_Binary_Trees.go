package binary

// 894. All Possible Full Binary Trees
// 直接翻译的java版本
func allPossibleFBT(n int) []*TreeNode {
	var list []*TreeNode
	if n == 1 {
		list = append(list, &TreeNode{Val: 0})
		return list
	}
	l := 1
	for l < n-1 {
		leftRoots := allPossibleFBT(l)
		rightRoots := allPossibleFBT(n - 1 - l)
		for _, leftRoot := range leftRoots {
			for _, rightRoot := range rightRoots {
				r := &TreeNode{Val: 0, Left: leftRoot, Right: rightRoot}
				list = append(list, r)
			}
		}
		l += 2
	}
	return list
}
