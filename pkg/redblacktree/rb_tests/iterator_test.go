package rbtests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	t.Run("Should correct iterator", func(t *testing.T) {
		tree := initTree2()
		elems := tree.Values()
		iterator := tree.Iterator()
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
		iterator = tree.Iterator()

		assert.False(t, iterator.Next())
		assert.False(t, iterator.Next())
		assert.False(t, iterator.Next())
		iterator.End()
		assert.False(t, iterator.Prev())
		assert.False(t, iterator.Prev())
		assert.False(t, iterator.Prev())
	})
}
