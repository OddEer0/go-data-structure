package double_list_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListEnum(t *testing.T) {
	t.Run("Should correct Each", func(t *testing.T) {
		list := initList()
		i := 0
		list.Each(func(index int, val int) {
			value, _ := list.Get(i)
			assert.Equal(t, value, val)
			assert.Equal(t, i, index)
			i++
		})
	})

	t.Run("Should correct EachLast", func(t *testing.T) {
		list := initList()
		i := list.Size() - 1
		list.EachLast(func(index int, val int) {
			value, _ := list.Get(i)
			assert.Equal(t, value, val)
			assert.Equal(t, i, index)
			i--
		})
	})

	t.Run("Should correct Some", func(t *testing.T) {
		list := initList()
		has := list.Some(func(index int, val int) bool {
			if val == 7 {
				return true
			}
			return false
		})
		assert.True(t, has)
		has = list.Some(func(index int, val int) bool {
			if val == 132132 {
				return true
			}
			return false
		})
		assert.False(t, has)
	})

	t.Run("Should correct Every", func(t *testing.T) {
		list := initList()
		check := list.Every(func(index int, val int) bool {
			if val < 10 {
				return true
			}
			return false
		})
		assert.True(t, check)
		check = list.Every(func(index int, val int) bool {
			if val < 9 {
				return true
			}
			return false
		})
		assert.False(t, check)
	})

	t.Run("Should correct Map", func(t *testing.T) {
		list := initList()
		newList := list.Map(func(index int, val int) int {
			return val * 2
		})
		assert.Equal(t, list.Size(), newList.Size())
		assert.False(t, fmt.Sprintf("%p", list) == fmt.Sprintf("%p", newList))
		assert.False(t, fmt.Sprintf("%p", list.Head()) == fmt.Sprintf("%p", newList.Head()))
		assert.False(t, fmt.Sprintf("%p", list.Tail()) == fmt.Sprintf("%p", newList.Tail()))
	})

	t.Run("Should correct Filter", func(t *testing.T) {
		list := initList()
		newList := list.Filter(func(index int, val int) bool {
			if val < 9 {
				return true
			}
			return false
		})
		assert.Equal(t, list.Size()-1, newList.Size())
		assert.False(t, fmt.Sprintf("%p", list) == fmt.Sprintf("%p", newList))
		assert.False(t, fmt.Sprintf("%p", list.Head()) == fmt.Sprintf("%p", newList.Head()))
		assert.False(t, fmt.Sprintf("%p", list.Tail()) == fmt.Sprintf("%p", newList.Tail()))
	})

	t.Run("Should correct Concat", func(t *testing.T) {
		list := initList()
		list2 := list.Map(func(index int, val int) int {
			return val + list.Size()
		})
		list3 := list2.Map(func(index int, val int) int {
			return val + list2.Size()
		})
		assert.Equal(t, list.Size(), list2.Size())
		assert.Equal(t, list.Size(), list3.Size())

		concatList := list.Concat(list2, list3)
		assert.Equal(t, 27, concatList.Size())
		i := 1
		concatList.Each(func(index int, val int) {
			assert.Equal(t, i-1, index)
			assert.Equal(t, i, val)
			i++
		})
	})
}
