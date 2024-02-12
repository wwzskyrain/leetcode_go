package binary

// 这个学到的东西挺多的。
// 1. 如何定义多递归面的全局变量——
// 2. 如何定义匿名函数
// 3. 如何来引用其他目录下的struct呢？这个问题仍然没有被解决?
func preorderTraversal(root *TreeNode) []int {
	var list []int
	var preOrder func(node *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		list = append(list, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)
	return list
}
