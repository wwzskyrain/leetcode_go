package grid

import (
	"math/bits"
)

// 2397. Maximum Rows Covered by Columns

func maximumRows(matrix [][]int, numSelect int) int {
	m, n := len(matrix), len(matrix[0])
	mask := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mask[i] += matrix[i][j] << (n - j - 1)
		}
	}
	res, limit := 0, 1<<n
	for cur := 0; cur < limit; cur++ {
		if bits.OnesCount(uint(cur)) != numSelect {
			continue
		}
		t := 0
		for j := 0; j < m; j++ {
			if (mask[j] & cur) == mask[j] {
				t++
			}
		}
		res = max(res, t)
	}
	return res
}
