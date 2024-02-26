package binary

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTreeFromPreOrderAndInorder(preorder, inorder)
	fmt.Printf("root = %v", root)
}
