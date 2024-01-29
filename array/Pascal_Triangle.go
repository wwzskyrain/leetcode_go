package array

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans { //哈哈，这个range用的不错吧
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}
