package medium

// 331. Verify Preorder Serialization of a Binary Tree
func isValidSerialization(preorder string) bool {
	n := len(preorder)
	slots := 1
	for i := 0; i < n; i++ {
		if slots == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			slots--
			i++
		} else {
			//	读第一个数字
			for i < n && preorder[i] != ',' {
				i++
			}
			slots++ //slots = slots - 1 + 2
		}
	}
	return slots == 0
}
