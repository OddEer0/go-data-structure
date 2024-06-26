package rbtests

import (
	"fmt"
	"testing"

	"github.com/OddEer0/go-data-structure/tree/redblacktree"
	"github.com/stretchr/testify/assert"
)

func initTree() redblacktree.Tree[int, int] {
	newTree := redblacktree.New[int, int]()
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

func initTree2() redblacktree.Tree[int, int] {
	tree := initTree()

	tree.Insert(1220, 1220)
	tree.Insert(600, 600)
	tree.Insert(650, 650)

	return tree
}

func TestBaseMethods(t *testing.T) {
	t.Run("Should correct insert and balance tree", func(t *testing.T) {
		tree := redblacktree.New[int, int]()
		tree2 := redblacktree.New[int, int]()

		tree2.Insert(1000, 1000)
		tree2.Insert(900, 900)
		tree2.Insert(800, 800)
		assert.Equal(t, tree2.Root().Value(), 900)
		assert.Equal(t, tree2.Root().Left().Value(), 800)
		assert.Equal(t, tree2.Root().Right().Value(), 1000)

		node := tree.Insert(1000, 1000)
		assert.Nil(t, node)
		assert.Equal(t, tree.Size(), 1)
		assert.Equal(t, tree.Root().Value(), 1000)
		assert.Equal(t, tree.Root().IsBlack(), true)

		node = tree.Insert(1050, 1050)
		assert.Nil(t, node)
		node = tree.Insert(950, 950)
		assert.Nil(t, node)
		assert.Equal(t, tree.Size(), 3)
		assert.Equal(t, tree.Root().Value(), 1000)
		assert.Equal(t, []int{1000, 950, 1050}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Right().IsRed(), true)

		node = tree.Insert(700, 700)
		assert.Nil(t, node)
		assert.Equal(t, tree.Size(), 4)
		assert.Equal(t, tree.Root().Value(), 1000)
		assert.Equal(t, []int{1000, 950, 700, 1050}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().IsRed(), true)

		node = tree.Insert(960, 960)
		assert.Nil(t, node)
		node = tree.Insert(1030, 1030)
		assert.Nil(t, node)
		node = tree.Insert(1100, 1100)
		assert.Nil(t, node)
		assert.Equal(t, tree.Size(), 7)
		assert.Equal(t, tree.Root().Value(), 1000)
		assert.Equal(t, []int{1000, 950, 700, 960, 1050, 1030, 1100}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Right().IsRed(), true)

		node = tree.Insert(1150, 1150)
		assert.Nil(t, node)
		assert.Equal(t, tree.Size(), 8)
		assert.Equal(t, tree.Root().Value(), 1000)
		assert.Equal(t, []int{1000, 950, 700, 960, 1050, 1030, 1100, 1150}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Right().IsRed(), true)

		node = tree.Insert(1200, 1200)
		assert.Nil(t, node)
		assert.Equal(t, tree.Size(), 9)
		assert.Equal(t, tree.Root().Value(), 1000)
		assert.Equal(t, []int{1000, 950, 700, 960, 1050, 1030, 1150, 1100, 1200}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Right().Left().IsRed(), true)

		node = tree.Insert(1250, 1250)
		assert.Nil(t, node)
		assert.Equal(t, tree.Size(), 10)
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 700, 960, 1030, 1150, 1100, 1200, 1250}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Right().IsRed(), true)

		node = tree.Insert(1250, 1250)
		assert.NotNil(t, node)
		assert.Equal(t, node.Value(), 1250)
		assert.Equal(t, tree.Size(), 10)
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 700, 960, 1030, 1150, 1100, 1200, 1250}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Right().IsRed(), true)

		node = tree.Insert(1220, 1220)
		assert.Nil(t, node)
		assert.Equal(t, tree.Size(), 11)
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 700, 960, 1030, 1150, 1100, 1220, 1200, 1250}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Right().Left().IsRed(), true)

		node = tree.Insert(600, 600)
		assert.Nil(t, node)
		assert.Equal(t, tree.Size(), 12)
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 700, 600, 960, 1030, 1150, 1100, 1220, 1200, 1250}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Right().Left().IsRed(), true)

		node = tree.Insert(650, 650)
		assert.Nil(t, node)
		assert.Equal(t, tree.Size(), 13)
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 650, 600, 700, 960, 1030, 1150, 1100, 1220, 1200, 1250}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Right().Left().IsRed(), true)
	})

	t.Run("Should correct left right working", func(t *testing.T) {
		tree := initTree()
		assert.Equal(t, 700, tree.Left().Value())
		assert.Equal(t, 1250, tree.Right().Value())

		tree = redblacktree.New[int, int]()
		assert.Nil(t, tree.Left())
		assert.Nil(t, tree.Right())
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

	t.Run("Should correct get", func(t *testing.T) {
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
			value, ok := tree.Get(c.key)
			if ok {
				assert.Equal(t, c.val, value)
			}
			assert.Equal(t, ok, c.found)
		}
	})

	t.Run("Should correct update", func(t *testing.T) {
		tc := []struct {
			key       int
			val       int
			isSuccess bool
		}{
			{1050, 2000, true},
			{960, 2500, true},
			{1200, 2200, true},
			{10000, 2000, false},
		}

		tree := initTree()

		for _, c := range tc {
			ok := tree.Update(c.key, c.val)
			if ok {
				val, _ := tree.Get(c.key)
				assert.Equal(t, val, c.val)
			}
			assert.Equal(t, c.isSuccess, ok)
		}
	})

	t.Run("Should correct remove and balance element", func(t *testing.T) {
		tree := initTree2()
		tree.Remove(1250)
		tree.Remove(1200)
		assert.Equal(t, tree.Size(), 11)
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 650, 600, 700, 960, 1030, 1150, 1100, 1220}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)

		tree.Remove(1220)
		assert.Equal(t, tree.Size(), 10)
		assert.Equal(t, tree.Root().Value(), 1000)
		assert.Equal(t, []int{1000, 950, 650, 600, 700, 960, 1050, 1030, 1150, 1100}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().Right().IsRed(), true)
		assert.Equal(t, tree.Root().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Left().IsRed(), true)

		tree = initTree2()
		tree.Remove(960)
		assert.Equal(t, tree.Size(), 12)
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 650, 600, 950, 700, 1030, 1150, 1100, 1220, 1200, 1250}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Right().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Right().Right().IsRed(), true)

		tree.Remove(600)
		assert.Equal(t, tree.Size(), 11)
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 700, 650, 950, 1030, 1150, 1100, 1220, 1200, 1250}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Right().Right().IsRed(), true)

		tree.Remove(1030)
		assert.Equal(t, tree.Size(), 10)
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, []int{1050, 700, 650, 1000, 950, 1150, 1100, 1220, 1200, 1250}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Right().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Right().Right().Right().IsRed(), true)

		tree = initTree2()
		tree.Remove(1150)
		assert.Equal(t, tree.Size(), 12)
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 650, 600, 700, 960, 1030, 1200, 1100, 1220, 1250}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Right().IsRed(), true)

		tree.Remove(2000)
		assert.Equal(t, tree.Size(), 12)
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, []int{1050, 1000, 950, 650, 600, 700, 960, 1030, 1200, 1100, 1220, 1250}, tree.PreOrderKeys())
		assert.Equal(t, tree.Root().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Left().Left().Left().Left().IsRed(), true)
		assert.Equal(t, tree.Root().Left().Left().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Left().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().IsBlack(), true)
		assert.Equal(t, tree.Root().Right().Right().Right().IsRed(), true)
	})

	t.Run("Getter and Setter", func(t *testing.T) {
		tree := initTree2()
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, tree.Size(), 13)
	})

	t.Run("Should clear and is empty correct working", func(t *testing.T) {
		tree := initTree2()
		assert.Equal(t, tree.Root().Value(), 1050)
		assert.Equal(t, tree.Size(), 13)
		assert.Equal(t, tree.IsEmpty(), false)

		tree.Clear()
		assert.Equal(t, tree.IsEmpty(), true)
		assert.Equal(t, tree.Size(), 0)
		assert.Nil(t, tree.Root())
	})

	t.Run("Should correctly copy tree", func(t *testing.T) {
		tree := initTree2()
		copyTree := tree.Copy()
		assert.False(t, fmt.Sprintf("%p", tree.Root()) == fmt.Sprintf("%p", copyTree.Root()))
		assert.Equal(t, tree.Size(), copyTree.Size())

		assert.Equal(t, tree.Values(), copyTree.Values())
		assert.Equal(t, tree.PreOrderValues(), copyTree.PreOrderValues())
		assert.Equal(t, tree.PostOrderValues(), copyTree.PostOrderValues())

		tree2 := redblacktree.New[int, int]()
		copyTree2 := tree2.Copy()
		assert.False(t, fmt.Sprintf("%p", tree2) == fmt.Sprintf("%p", copyTree2))
		assert.Equal(t, tree2.Size(), copyTree2.Size())

		tree3 := redblacktree.New[int, int]()
		for i := 1; i <= 3000; i++ {
			tree3.Insert(i, i)
		}
		copyTree3 := tree3.Copy()
		assert.Equal(t, tree3.Values(), copyTree3.Values())
		assert.Equal(t, tree3.PreOrderValues(), copyTree3.PreOrderValues())
		assert.Equal(t, tree3.PostOrderValues(), copyTree3.PostOrderValues())
	})

	t.Run("Should correct String method", func(t *testing.T) {
		tree := initTree2()
		expectedString := `RedBlackTree
│           ┌── Key: 1250, Value: 1250
│       ┌── Key: 1220, Value: 1220
│       │   └── Key: 1200, Value: 1200
│   ┌── Key: 1150, Value: 1150
│   │   └── Key: 1100, Value: 1100
└── Key: 1050, Value: 1050
    │   ┌── Key: 1030, Value: 1030
    └── Key: 1000, Value: 1000
        │   ┌── Key: 960, Value: 960
        └── Key: 950, Value: 950
            │   ┌── Key: 700, Value: 700
            └── Key: 650, Value: 650
                └── Key: 600, Value: 600` + "\n"
		assert.Equal(t, expectedString, tree.String())
	})
}
