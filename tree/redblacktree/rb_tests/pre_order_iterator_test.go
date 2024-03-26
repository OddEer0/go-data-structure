package rbtests

import (
	"testing"

	"github.com/OddEer0/go-data-structure/tree/redblacktree"
	"github.com/stretchr/testify/assert"
)

func TestPreOrderIterator(t *testing.T) {
	t.Run("Should correct iterator", func(t *testing.T) {
		tree := initTree2()
		elems := tree.PreOrderValues()
		iterator := tree.PreOrderIterator()
		for i := 0; iterator.Next(); i++ {
			assert.Equal(t, elems[i], iterator.Value())
			assert.Equal(t, elems[i], iterator.Key())
			assert.Equal(t, elems[i], iterator.Node().Value())
			assert.Equal(t, elems[i], iterator.Node().Key())
		}

		assert.Equal(t, iterator.Next(), false)

		for i := 0; iterator.Prev(); i++ {
			assert.Equal(t, elems[len(elems)-1-i], iterator.Value())
			assert.Equal(t, elems[len(elems)-1-i], iterator.Key())
			assert.Equal(t, elems[len(elems)-1-i], iterator.Node().Value())
			assert.Equal(t, elems[len(elems)-1-i], iterator.Node().Key())
		}

		assert.Equal(t, iterator.Prev(), false)

		iterator.Next()
		iterator.Next()
		iterator.Next()
		iterator.First()
		assert.Equal(t, elems[0], iterator.Value())
		iterator.Last()
		assert.Equal(t, elems[len(elems)-1], iterator.Value())

		iterator.Start()
		iterator.Next()
		assert.Equal(t, elems[0], iterator.Value())
		iterator.End()
		iterator.Prev()
		assert.Equal(t, elems[len(elems)-1], iterator.Value())

		tree.Clear()
		iterator = tree.PreOrderIterator()

		assert.False(t, iterator.Next())
		assert.False(t, iterator.Next())
		assert.False(t, iterator.Next())
		iterator.End()
		assert.False(t, iterator.Prev())
		assert.False(t, iterator.Prev())
		assert.False(t, iterator.Prev())
	})

	t.Run("Should correct PreOrderIteration 3000 elems", func(t *testing.T) {
		for v := 0; v <= 19; v++ {
			tree := redblacktree.New[int, int]()
			for i := 0; i < 3000; i++ {
				tree.Insert(i*v, i*v)
			}
			elems := tree.PreOrderValues()
			iterator := tree.PreOrderIterator()
			for i := 0; iterator.Next(); i++ {
				assert.Equal(t, elems[i], iterator.Value())
				assert.Equal(t, elems[i], iterator.Key())
				assert.Equal(t, elems[i], iterator.Node().Value())
				assert.Equal(t, elems[i], iterator.Node().Key())
			}

			assert.Equal(t, iterator.Next(), false)

			for i := 0; iterator.Prev(); i++ {
				assert.Equal(t, elems[len(elems)-1-i], iterator.Value())
				assert.Equal(t, elems[len(elems)-1-i], iterator.Key())
				assert.Equal(t, elems[len(elems)-1-i], iterator.Node().Value())
				assert.Equal(t, elems[len(elems)-1-i], iterator.Node().Key())
			}

			assert.Equal(t, iterator.Prev(), false)

			for i := 1000; i <= 2000; i++ {
				tree.Remove(i * v)
			}
			elems = tree.PreOrderValues()

			for i := 0; iterator.Next(); i++ {
				assert.Equal(t, elems[i], iterator.Value())
				assert.Equal(t, elems[i], iterator.Key())
				assert.Equal(t, elems[i], iterator.Node().Value())
				assert.Equal(t, elems[i], iterator.Node().Key())
			}

			assert.Equal(t, iterator.Next(), false)

			for i := 0; iterator.Prev(); i++ {
				assert.Equal(t, elems[len(elems)-1-i], iterator.Value())
				assert.Equal(t, elems[len(elems)-1-i], iterator.Key())
				assert.Equal(t, elems[len(elems)-1-i], iterator.Node().Value())
				assert.Equal(t, elems[len(elems)-1-i], iterator.Node().Key())
			}

			assert.Equal(t, iterator.Prev(), false)
		}
	})
}
