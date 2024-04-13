package double_list_test

import (
	"testing"

	doublelist "github.com/OddEer0/go-data-structure/list/double_list"
	"github.com/stretchr/testify/assert"
)

func TestListIterator(t *testing.T) {
	t.Run("Should correct iterator", func(t *testing.T) {
		list := initList()
		iterator := list.Iterator()

		i := 0
		for iterator.Next() {
			val, _ := list.Get(i)
			assert.Equal(t, val, iterator.Value())
			assert.Equal(t, i, iterator.Index())
			i++
		}

		assert.False(t, iterator.Next())

		i--
		for iterator.Prev() {
			val, _ := list.Get(i)
			assert.Equal(t, val, iterator.Value())
			assert.Equal(t, i, iterator.Index())
			i--
		}

		assert.False(t, iterator.Prev())

		iterator.Next()
		iterator.Next()
		iterator.Next()
		assert.Equal(t, iterator.Node().Value(), 3)
		assert.Equal(t, iterator.Index(), 2)

		iterator.End()
		assert.False(t, iterator.Next())

		iterator.Prev()
		iterator.Prev()
		iterator.Prev()
		assert.Equal(t, iterator.Node().Value(), 7)
		assert.Equal(t, iterator.Index(), 6)

		assert.True(t, iterator.Last())
		assert.Equal(t, iterator.Index(), list.Size()-1)
		assert.Equal(t, iterator.Value(), 9)

		assert.True(t, iterator.First())
		assert.Equal(t, iterator.Index(), 0)
		assert.Equal(t, iterator.Value(), 1)
	})

	t.Run("Should correct empty list iterator", func(t *testing.T) {
		list := doublelist.New[int]()
		iterator := list.Iterator()
		assert.False(t, iterator.Next())
		iterator.End()
		assert.False(t, iterator.Prev())
	})
}
