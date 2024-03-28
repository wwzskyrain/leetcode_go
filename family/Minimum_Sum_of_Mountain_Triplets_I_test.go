package family

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumSum(t *testing.T) {
	a := assert.New(t)
	a.Equal(9, minimumSum([]int{8, 6, 1, 5, 3}))
	a.Equal(13, minimumSum([]int{5, 4, 8, 7, 10, 2}))
}
