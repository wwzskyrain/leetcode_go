package medium

// 1535. Find the Winner of an Array Game
func getWinner(arr []int, k int) int {
	m := arr[0]
	count := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > m {
			m = arr[i]
			count = 1
		} else {
			count += 1
		}
		if count == k {
			return m
		}
	}
	return m
}
