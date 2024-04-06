package easy

// 双端队列的玩法
// 2810. Faulty Keyboard
func finalString(s string) string {
	var q []rune
	head := false
	for _, ch := range s {
		if ch != 'i' {
			if head {
				q = append([]rune{ch}, q...)
			} else {
				q = append(q, ch)
			}
		} else {
			head = !head
		}
	}
	if head {
		for i, j := 0, len(q)-1; i < j; {
			q[i], q[j] = q[j], q[i]
			i++
			j--
		}
	}
	return string(q)
}
