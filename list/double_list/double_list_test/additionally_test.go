package double_list_test

import (
	"fmt"
	doublelist "github.com/OddEer0/go-data-structure/list/double_list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListAdditionally(t *testing.T) {
	t.Run("Should correct As method", func(t *testing.T) {
		list := initList()
		var res int
		assert.Equal(t, 9, list.Size())
		actual, err := list.As(0)
		assert.NoError(t, err)
		assert.Equal(t, 1, actual)
		actual, err = list.As(-1)
		assert.NoError(t, err)
		assert.Equal(t, 9, actual)
		actual, err = list.As(list.Size())
		assert.ErrorIs(t, doublelist.ErrOutOfRange, err)
		assert.Equal(t, res, actual)
		actual, err = list.As(-list.Size() - 1)
		assert.ErrorIs(t, doublelist.ErrOutOfRange, err)
		assert.Equal(t, res, actual)

		actual, err = list.As(2)
		assert.NoError(t, err)
		assert.Equal(t, 3, actual)
	})

	t.Run("Should correct Find", func(t *testing.T) {
		list := initList()
		var res int
		actual, found := list.Find(func(index int, item int) bool {
			if item == 8 {
				return true
			}
			return false
		})
		assert.True(t, found)
		assert.Equal(t, 8, actual)

		actual, found = list.Find(func(index int, item int) bool {
			if item == 123121 {
				return true
			}
			return false
		})
		assert.False(t, found)
		assert.Equal(t, res, actual)
	})

	t.Run("Should correct FindIndex", func(t *testing.T) {
		list := initList()
		res := -1
		actual := list.FindIndex(func(index int, item int) bool {
			if item == 8 {
				return true
			}
			return false
		})
		assert.Equal(t, 7, actual)

		actual = list.FindIndex(func(index int, item int) bool {
			if item == 123121 {
				return true
			}
			return false
		})
		assert.Equal(t, res, actual)
	})

	t.Run("Should correct Reduce", func(t *testing.T) {
		list := initList()
		iter := list.Iterator()
		sum := 0
		for iter.Next() {
			sum += iter.Value()
		}

		actual := list.Reduce(func(acc interface{}, index int, item int) interface{} {
			return acc.(int) + item
		}, 0)

		assert.Equal(t, sum, actual)
	})

	t.Run("Should correct contains", func(t *testing.T) {
		list := initList()
		assert.True(t, list.Contains(9))
		assert.True(t, list.Contains(1))
		assert.False(t, list.Contains(10))
		assert.False(t, list.Contains(0))
	})

	t.Run("Should correct search", func(t *testing.T) {
		list := initList()
		var res int
		actual, found := list.Search(9)
		assert.True(t, found)
		assert.Equal(t, 9, actual)
		actual, found = list.Search(1)
		assert.True(t, found)
		assert.Equal(t, 1, actual)
		actual, found = list.Search(10)
		assert.False(t, found)
		assert.Equal(t, res, actual)
		actual, found = list.Search(0)
		assert.False(t, found)
		assert.Equal(t, res, actual)
	})

	t.Run("Should correct indexOf", func(t *testing.T) {
		list := initList()
		list.Push(9)
		res := -1
		assert.Equal(t, 8, list.IndexOf(9))
		assert.Equal(t, 0, list.IndexOf(1))
		assert.Equal(t, res, list.IndexOf(10))
		assert.Equal(t, res, list.IndexOf(0))
	})

	t.Run("Should correct LastIndexOf", func(t *testing.T) {
		list := initList()
		list.Push(9)
		res := -1
		assert.Equal(t, 9, list.LastIndexOf(9))
		assert.Equal(t, 0, list.LastIndexOf(1))
		assert.Equal(t, res, list.LastIndexOf(10))
		assert.Equal(t, res, list.LastIndexOf(0))
	})

	t.Run("Should correct Swap", func(t *testing.T) {
		list := initList()
		var actual int
		assert.NoError(t, list.Swap(0, 8))
		actual, _ = list.Get(0)
		assert.Equal(t, 9, actual)
		actual, _ = list.Get(8)
		assert.Equal(t, 1, actual)

		assert.NoError(t, list.Swap(8, 8))
		actual, _ = list.Get(0)
		assert.Equal(t, 9, actual)
		actual, _ = list.Get(8)
		assert.Equal(t, 1, actual)

		assert.ErrorIs(t, doublelist.ErrOutOfRange, list.Swap(-1, 8))
		actual, _ = list.Get(0)
		assert.Equal(t, 9, actual)
		actual, _ = list.Get(8)
		assert.Equal(t, 1, actual)

		assert.ErrorIs(t, doublelist.ErrOutOfRange, list.Swap(9, 8))
		actual, _ = list.Get(0)
		assert.Equal(t, 9, actual)
		actual, _ = list.Get(8)
		assert.Equal(t, 1, actual)

		assert.ErrorIs(t, doublelist.ErrOutOfRange, list.Swap(0, -1))
		actual, _ = list.Get(0)
		assert.Equal(t, 9, actual)
		actual, _ = list.Get(8)
		assert.Equal(t, 1, actual)

		assert.ErrorIs(t, doublelist.ErrOutOfRange, list.Swap(0, 9))
		actual, _ = list.Get(0)
		assert.Equal(t, 9, actual)
		actual, _ = list.Get(8)
		assert.Equal(t, 1, actual)
	})

	t.Run("Should correct Reverse", func(t *testing.T) {
		list := initList()
		list.Reverse()
		i := list.Size()
		iter := list.Iterator()
		for iter.Next() {
			assert.Equal(t, i, iter.Value())
			i--
		}
	})

	t.Run("Should correct ToReverse", func(t *testing.T) {
		list := initList()
		newList := list.ToReversed()
		assert.False(t, fmt.Sprintf("%p", list) == fmt.Sprintf("%p", newList))
		assert.False(t, fmt.Sprintf("%p", list.Head()) == fmt.Sprintf("%p", newList.Head()))
		assert.False(t, fmt.Sprintf("%p", list.Tail()) == fmt.Sprintf("%p", newList.Tail()))

		i := list.Size()
		iter := newList.Iterator()
		for iter.Next() {
			assert.Equal(t, i, iter.Value())
			i--
		}

		i = 1
		iter = list.Iterator()
		for iter.Next() {
			assert.Equal(t, i, iter.Value())
			i++
		}
	})

	t.Run("Should correct Slice", func(t *testing.T) {
		list := initList()
		actual, err := list.Slice(-1, 8)
		assert.ErrorIs(t, doublelist.ErrOutOfRange, err)
		assert.Nil(t, actual)

		actual, err = list.Slice(9, 10)
		assert.ErrorIs(t, doublelist.ErrOutOfRange, err)
		assert.Nil(t, actual)

		actual, err = list.Slice(8, 10)
		assert.ErrorIs(t, doublelist.ErrOutOfRange, err)
		assert.Nil(t, actual)

		actual, err = list.Slice(6, 6)
		assert.ErrorIs(t, doublelist.ErrEndIndexLessOrEqualStart, err)
		assert.Nil(t, actual)

		actual, err = list.Slice(6, 5)
		assert.ErrorIs(t, doublelist.ErrEndIndexLessOrEqualStart, err)
		assert.Nil(t, actual)

		actual, err = list.Slice(8, 9)
		assert.NoError(t, err)
		assert.Equal(t, 1, actual.Size())
		assert.False(t, fmt.Sprintf("%p", list) == fmt.Sprintf("%p", actual))
		assert.False(t, fmt.Sprintf("%p", list.Head()) == fmt.Sprintf("%p", actual.Head()))
		assert.False(t, fmt.Sprintf("%p", list.Tail()) == fmt.Sprintf("%p", actual.Tail()))
		assert.Equal(t, "9", actual.String())

		actual, err = list.Slice(0, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, actual.Size())
		assert.False(t, fmt.Sprintf("%p", list) == fmt.Sprintf("%p", actual))
		assert.False(t, fmt.Sprintf("%p", list.Head()) == fmt.Sprintf("%p", actual.Head()))
		assert.False(t, fmt.Sprintf("%p", list.Tail()) == fmt.Sprintf("%p", actual.Tail()))
		assert.Equal(t, "1", actual.String())

		actual, err = list.Slice(0, 3)
		assert.NoError(t, err)
		assert.Equal(t, 3, actual.Size())
		assert.False(t, fmt.Sprintf("%p", list) == fmt.Sprintf("%p", actual))
		assert.False(t, fmt.Sprintf("%p", list.Head()) == fmt.Sprintf("%p", actual.Head()))
		assert.False(t, fmt.Sprintf("%p", list.Tail()) == fmt.Sprintf("%p", actual.Tail()))
		assert.Equal(t, "1 2 3", actual.String())
	})
}
