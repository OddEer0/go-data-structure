package avltests

import (
	"testing"

	avltree "github.com/OddEer0/go-data-structure/pkg/avl_tree"
	"github.com/stretchr/testify/assert"
)

func initTree() avltree.ITree[int, int] {
	newTree := avltree.NewAVLTree[int, int]()
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
	t.Run("Should correct insert and balance", func(t *testing.T) {
		tree := avltree.NewAVLTree[int, int]()

		tree.Insert(20, 20)
		assert.Equal(t, tree.GetSize(), 1)
		assert.Equal(t, tree.GetRoot().Value(), 20)

		tree.Insert(15, 15)
		tree.Insert(10, 10)
		assert.Equal(t, tree.GetSize(), 3)
		assert.Equal(t, tree.GetRoot().Value(), 15)
		assert.Equal(t, []int{15, 10, 20}, tree.ToPreOrderNodeSlice())

		tree.Insert(25, 25)
		tree.Insert(35, 35)
		assert.Equal(t, tree.GetSize(), 5)
		assert.Equal(t, tree.GetRoot().Value(), 15)
		assert.Equal(t, []int{10, 20, 35, 25, 15}, tree.ToPostOrderNodeSlice())
		assert.Equal(t, []int{15, 10, 25, 20, 35}, tree.ToPreOrderNodeSlice())

		tree.Insert(30, 30)
		assert.Equal(t, tree.GetSize(), 6)
		assert.Equal(t, tree.GetRoot().Value(), 25)
		assert.Equal(t, []int{10, 20, 15, 30, 35, 25}, tree.ToPostOrderNodeSlice())
		assert.Equal(t, []int{25, 15, 10, 20, 35, 30}, tree.ToPreOrderNodeSlice())

		tree.Insert(40, 40)
		assert.Equal(t, tree.GetSize(), 7)
		assert.Equal(t, tree.GetRoot().Value(), 25)
		assert.Equal(t, []int{10, 20, 15, 30, 40, 35, 25}, tree.ToPostOrderNodeSlice())
		assert.Equal(t, []int{25, 15, 10, 20, 35, 30, 40}, tree.ToPreOrderNodeSlice())
	})

	t.Run("Should correct remove element", func(t *testing.T) {
		tree := initTree()

		tree.Remove(30)
		tree.Remove(40)
		assert.Equal(t, tree.GetSize(), 6)
		assert.Equal(t, tree.GetRoot().Value(), 15)
		assert.Equal(t, []int{5, 10, 20, 35, 25, 15}, tree.ToPostOrderNodeSlice())
		assert.Equal(t, []int{15, 10, 5, 25, 20, 35}, tree.ToPreOrderNodeSlice())
	})
}
