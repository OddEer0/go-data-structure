package rbtests

import (
	"testing"

	rbtree "github.com/OddEer0/go-data-structure/pkg/rb_tree"
	"github.com/stretchr/testify/assert"
)

func initTree() rbtree.ITree[int, int] {
	newTree := rbtree.NewRBTree[int, int]()
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

func TestNodeBaseMethods(t *testing.T) {
	t.Run("Should correct insert and balance tree", func(t *testing.T) {
		tree := rbtree.NewRBTree[int, int]()
		tree2 := rbtree.NewRBTree[int, int]()

		tree2.Insert(1000, 1000)
		tree2.Insert(900, 900)
		tree2.Insert(800, 800)
		assert.Equal(t, tree2.GetRoot().Value(), 900)
		assert.Equal(t, tree2.GetRoot().Left().Value(), 800)
		assert.Equal(t, tree2.GetRoot().Right().Value(), 1000)

		node := tree.Insert(1000, 1000)
		assert.Nil(t, node)
		assert.Equal(t, tree.GetSize(), 1)
		assert.Equal(t, tree.GetRoot().Value(), 1000)
		assert.Equal(t, tree.GetRoot().IsBlack(), true)

		node = tree.Insert(1050, 1050)
		assert.Nil(t, node)
		node = tree.Insert(950, 950)
		assert.Nil(t, node)
		assert.Equal(t, tree.GetSize(), 3)
		assert.Equal(t, tree.GetRoot().Value(), 1000)
		assert.Equal(t, []int{1000, 950, 1050}, tree.ToPreOrderNodeSlice())
		assert.Equal(t, tree.GetRoot().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().IsRed(), true)

		node = tree.Insert(700, 700)
		assert.Nil(t, node)
		assert.Equal(t, tree.GetSize(), 4)
		assert.Equal(t, tree.GetRoot().Value(), 1000)
		assert.Equal(t, []int{1000, 950, 700, 1050}, tree.ToPreOrderNodeSlice())
		assert.Equal(t, tree.GetRoot().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().IsRed(), true)

		node = tree.Insert(960, 960)
		assert.Nil(t, node)
		node = tree.Insert(1030, 1030)
		assert.Nil(t, node)
		node = tree.Insert(1100, 1100)
		assert.Nil(t, node)
		assert.Equal(t, tree.GetSize(), 7)
		assert.Equal(t, tree.GetRoot().Value(), 1000)
		assert.Equal(t, []int{1000, 950, 700, 960, 1050, 1030, 1100}, tree.ToPreOrderNodeSlice())
		assert.Equal(t, tree.GetRoot().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().IsRed(), true)

		node = tree.Insert(1150, 1150)
		assert.Nil(t, node)
		assert.Equal(t, tree.GetSize(), 8)
		assert.Equal(t, tree.GetRoot().Value(), 1000)
		assert.Equal(t, []int{1000, 950, 700, 960, 1050, 1030, 1100, 1150}, tree.ToPreOrderNodeSlice())
		assert.Equal(t, tree.GetRoot().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().Right().IsRed(), true)

		node = tree.Insert(1200, 1200)
		assert.Nil(t, node)
		assert.Equal(t, tree.GetSize(), 9)
		assert.Equal(t, tree.GetRoot().Value(), 1000)
		assert.Equal(t, []int{1000, 950, 700, 960, 1050, 1030, 1150, 1100, 1200}, tree.ToPreOrderNodeSlice())
		assert.Equal(t, tree.GetRoot().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().Left().IsRed(), true)

		node = tree.Insert(1250, 1250)
		assert.Nil(t, node)
		assert.Equal(t, tree.GetSize(), 10)
		assert.Equal(t, tree.GetRoot().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 700, 960, 1030, 1150, 1100, 1200, 1250}, tree.ToPreOrderNodeSlice())
		assert.Equal(t, tree.GetRoot().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().Right().IsRed(), true)

		node = tree.Insert(1250, 1250)
		assert.NotNil(t, node)
		assert.Equal(t, node.Value(), 1250)
		assert.Equal(t, tree.GetSize(), 10)
		assert.Equal(t, tree.GetRoot().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 700, 960, 1030, 1150, 1100, 1200, 1250}, tree.ToPreOrderNodeSlice())
		assert.Equal(t, tree.GetRoot().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().Right().IsRed(), true)

		node = tree.Insert(1220, 1220)
		assert.Nil(t, node)
		assert.Equal(t, tree.GetSize(), 11)
		assert.Equal(t, tree.GetRoot().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 700, 960, 1030, 1150, 1100, 1220, 1200, 1250}, tree.ToPreOrderNodeSlice())
		assert.Equal(t, tree.GetRoot().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().Left().IsRed(), true)

		node = tree.Insert(600, 600)
		assert.Nil(t, node)
		assert.Equal(t, tree.GetSize(), 12)
		assert.Equal(t, tree.GetRoot().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 700, 600, 960, 1030, 1150, 1100, 1220, 1200, 1250}, tree.ToPreOrderNodeSlice())
		assert.Equal(t, tree.GetRoot().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().Left().IsRed(), true)

		node = tree.Insert(650, 650)
		assert.Nil(t, node)
		assert.Equal(t, tree.GetSize(), 13)
		assert.Equal(t, tree.GetRoot().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 650, 600, 700, 960, 1030, 1150, 1100, 1220, 1200, 1250}, tree.ToPreOrderNodeSlice())
		assert.Equal(t, tree.GetRoot().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Left().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().Right().IsRed(), true)
		assert.Equal(t, tree.GetRoot().Right().Right().Left().IsRed(), true)
	})

	t.Run("Should correct get node", func(t *testing.T) {
		tc := []struct {
			key   int
			val   int
			found bool
		}{
			{1000, 1000, true},
			{700, 700, true},
			{10000, 10000, false},
		}

		tree := initTree()

		for _, c := range tc {
			node, ok := tree.GetNode(c.key)
			if ok {
				assert.Equal(t, c.val, node.Value())
			} else {
				assert.Nil(t, node)
			}
			assert.Equal(t, ok, c.found)
		}
	})
}
