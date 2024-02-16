package binary

import (
	"fmt"
	"testing"
)

func TestBuildBinaryTree(t *testing.T) {
	tree := BuildBinaryTree("[3,9,20,null,null,15,7]")
	fmt.Printf("%v", tree)
}
