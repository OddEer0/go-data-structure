package rbtests

import (
	"testing"

	rbtree "github.com/OddEer0/go-data-structure/pkg/rb_tree"
	"github.com/stretchr/testify/assert"
)

func initTree() rbtree.ITree[int, int] {
	newTree := rbtree.NewRBTree[int, int]()
	newTree.Insert(20, 20)
	newTree.Insert(15, 15)
	newTree.Insert(10, 10)
	newTree.Insert(25, 25)
	newTree.Insert(35, 35)
	newTree.Insert(30, 30)
	newTree.Insert(40, 40)
	newTree.Insert(5, 5)

	return newTree
}

func TestNodeBaseMethods(t *testing.T) {
	t.Run("Should correct insert", func(t *testing.T) {
		tree := rbtree.NewRBTree[int, int]()

		tree.Insert(20, 20)
		assert.Equal(t, tree.GetSize(), 1)
		assert.Equal(t, tree.GetRoot().Value(), 20)

		tree.Insert(15, 15)
		tree.Insert(10, 10)
		assert.Equal(t, tree.GetSize(), 3)
		assert.Equal(t, tree.GetRoot().Value(), 15)
		assert.Equal(t, []int{15, 10, 20}, tree.ToPreOrderNodeSlice())
	})
}
