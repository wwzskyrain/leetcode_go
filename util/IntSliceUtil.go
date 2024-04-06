package util

import (
	"strconv"
	"strings"
)

func BuildIntSlice() {

}

// BuildIntSlice2Dimension 入参demo [[0,3],[0,4],[1,3],[2,4],[2,7],[3,5],[3,6],[3,7],[4,6]]
func BuildIntSlice2Dimension(s string) [][]int {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "[[", "", 1)
	s = strings.Replace(s, "]]", "", 1)
	pieces := strings.Split(s, "],[")
	var Slice2Dimension [][]int
	for _, str := range pieces {
		p := strings.Split(str, ",")
		var Slice []int
		for _, str2 := range p {
			i, _ := strconv.Atoi(str2)
			Slice = append(Slice, i)
		}
		Slice2Dimension = append(Slice2Dimension, Slice)
	}
	return Slice2Dimension
}
