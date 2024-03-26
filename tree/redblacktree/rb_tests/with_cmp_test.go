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

	t.Run("Should correct sort order by less 4000 elements", func(t *testing.T) {
		tree := redblacktree.NewWith[int, int](func(a, b int) int {
			if a == b {
				return 0
			} else if a > b {
				return -1
			}
			return 1
		})

		for i := 1; i <= 4000; i++ {
			tree.Insert(i, i)
		}
		i := 4000
		values := tree.Values()
		for _, val := range values {
			assert.Equal(t, i, val)
			i--
		}

		for j := 1001; j <= 1500; j++ {
			tree.Remove(j)
		}

		for j := 3001; j <= 3500; j++ {
			tree.Remove(j)
		}

		i = 4000
		values = tree.Values()
		for _, val := range values {
			assert.Equal(t, i, val)
			if val == 3501 {
				i = 3000
			} else if val == 1501 {
				i = 1000
			} else {
				i--
			}
		}
	})

	t.Run("Should correct work with hard cmp functions", func(t *testing.T) {
		tree := redblacktree.NewWith[int, int](func(a, b int) int {
			if a == b {
				return 0
			}
			if a < b {
				if a < 50 && a > 0 || b < 50 && b > 0 {
					return 1
				}
				return -1
			}
			if a < 50 && a > 0 || b < 50 && b > 0 {
				return -1
			}
			return 1
		})

		for i := 1; i <= 400; i++ {
			tree.Insert(i, i)
		}

		iterator := tree.Iterator()

		for i := 50; iterator.Next(); i++ {
			if i <= 400 {
				assert.Equal(t, iterator.Value(), i)
			} else {
				assert.Equal(t, iterator.Value(), 50-(i-400))
			}
		}

		for i := 35; i <= 75; i++ {
			tree.Remove(i)
		}

		for i := 76; iterator.Next(); i++ {
			if i <= 400 {
				assert.Equal(t, iterator.Value(), i)
			} else {
				assert.Equal(t, iterator.Value(), 50-(i-400))
			}
		}
	})
}
