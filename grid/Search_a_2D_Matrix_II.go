package grid

// 240. Search a 2D Matrix II
func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	column := len(matrix[0])
	i, j := 0, column-1
	for i < row && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}
