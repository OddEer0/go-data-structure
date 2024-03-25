package bstests

import (
	"testing"

	binarysearchtree "github.com/OddEer0/go-data-structure/tree/binarysearchtree"
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

	t.Run("Should correct get element", func(t *testing.T) {
		tc := []struct {
			key   int
			val   int
			found bool
		}{
			{10, 10, true},
			{14, 14, true},
			{1000, 1000, false},
		}

		tree := initTree()

		for _, c := range tc {
			node, ok := tree.GetNodeByKey(c.key)
			if ok {
				assert.Equal(t, c.val, node.Value())
			} else {
				assert.Nil(t, node)
			}
			assert.Equal(t, ok, c.found)
		}
	})

	t.Run("Should correct remove element", func(t *testing.T) {
		tc := []struct {
			deleteKey int
			sl        []int
		}{
			// {10, []int{5, 6, 7, 8, 13, 14, 15}},
			{100, []int{5, 6, 7, 8, 10, 13, 14, 15}},
			{14, []int{5, 6, 7, 8, 10, 13, 15}},
			{8, []int{5, 6, 7, 10, 13, 15}},
			{10, []int{5, 6, 7, 13, 15}},
		}

		tree := initTree()

		for _, c := range tc {
			tree.Remove(c.deleteKey)
			sl := tree.ToSortedSlice()
			assert.Equal(t, c.sl, sl)
		}
	})
}
