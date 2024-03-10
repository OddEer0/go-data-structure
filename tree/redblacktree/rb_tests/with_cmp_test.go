package rbtests

import (
	"testing"

	"github.com/OddEer0/go-data-structure/tree/redblacktree"
	"github.com/stretchr/testify/assert"
)

func initCmpTree() redblacktree.Tree[int, int] {
	newTree := redblacktree.NewWith[int, int](func(a, b int) int {
		if a == b {
			return 0
		} else if a > b {
			return -1
		}
		return 1
	})

	newTree.Insert(1000, 1000)
	newTree.Insert(1050, 1050)
	newTree.Insert(950, 950)
	newTree.Insert(700, 700)
	newTree.Insert(960, 960)
	newTree.Insert(1030, 1030)
	newTree.Insert(1100, 1100)
	newTree.Insert(1150, 1150)
	newTree.Insert(1200, 1200)
	newTree.Insert(1250, 1250)

	return newTree
}

func TestWithCmp(t *testing.T) {
	t.Run("Should correct sort order by less", func(t *testing.T) {
		tree := initCmpTree()
		assert.Equal(t, 1250, tree.Left().Value())
		assert.Equal(t, 700, tree.Right().Value())
		assert.Equal(t, 1050, tree.Root().Value())
		assert.Equal(t, []int{1250, 1200, 1150, 1100, 1050, 1030, 1000, 960, 950, 700}, tree.Keys())
		assert.Equal(t, []int{1050, 1150, 1200, 1250, 1100, 1000, 1030, 950, 960, 700}, tree.PreOrderKeys())
		assert.Equal(t, []int{1250, 1200, 1100, 1150, 1030, 960, 700, 950, 1000, 1050}, tree.PostOrderKeys())

		tree.Remove(1000)
		assert.Equal(t, []int{1250, 1200, 1150, 1100, 1050, 1030, 960, 950, 700}, tree.Keys())
		assert.Equal(t, []int{1050, 1150, 1200, 1250, 1100, 960, 1030, 950, 700}, tree.PreOrderKeys())
		assert.Equal(t, []int{1250, 1200, 1100, 1150, 1030, 700, 950, 960, 1050}, tree.PostOrderKeys())
	})
}
