package bstests

import (
	"testing"

	binarysearchtree "github.com/OddEer0/go-data-structure/pkg/binary_search_tree"
	"github.com/stretchr/testify/assert"
)

func initTree() binarysearchtree.ITree[int, int] {
	newTree := binarysearchtree.NewBSTree[int, int]()
	newTree.Insert(10, 10)
	newTree.Insert(5, 5)
	newTree.Insert(7, 7)
	newTree.Insert(6, 6)
	newTree.Insert(8, 8)
	newTree.Insert(14, 14)
	newTree.Insert(13, 13)
	newTree.Insert(15, 15)

	return newTree
}

func TestNodeBaseMethods(t *testing.T) {
	t.Run("Should correct insert value", func(t *testing.T) {
		tree := initTree()
		sl := tree.ToSortedSlice()
		expectSl := []int{5, 6, 7, 8, 10, 13, 14, 15}
		assert.Equal(t, sl, expectSl)
		assert.Equal(t, tree.GetSize(), len(expectSl))

		ok := tree.Insert(100, 100)
		assert.Equal(t, tree.GetSize(), len(expectSl)+1)
		assert.Equal(t, ok, true)

		ok = tree.Insert(100, 100)
		assert.Equal(t, tree.GetSize(), len(expectSl)+1)
		assert.Equal(t, ok, false)
	})
}
