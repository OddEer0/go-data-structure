package double_list_test

import (
	"fmt"
	"testing"

	doublelist "github.com/OddEer0/go-data-structure/list/double_list"
	"github.com/stretchr/testify/assert"
)

func initList() doublelist.List[int] {
	list := doublelist.New[int](1, 2, 3, 4, 5, 6, 7, 8, 9)
	return list
}

func TestBaseDoubleList(t *testing.T) {
	t.Run("Should correct Push", func(t *testing.T) {
		list := doublelist.New[int]()
		list.Push()
		assert.Equal(t, list.Size(), 0)

		list.Push(1, 2)
		list.Push(3)
		assert.Equal(t, list.Size(), 3)
	})

	t.Run("Should correct Pop", func(t *testing.T) {
		list := initList()
		assert.Equal(t, 9, list.Size())
		assert.Equal(t, 9, list.Pop())

		assert.Equal(t, 8, list.Size())
		assert.Equal(t, 8, list.Pop())

		list = doublelist.New[int]()
		var tmp int
		assert.Equal(t, 0, list.Size())
		assert.Equal(t, tmp, list.Pop())
		assert.Equal(t, 0, list.Size())
	})

	t.Run("Should correct Peek", func(t *testing.T) {
		list := initList()
		assert.Equal(t, 9, list.Size())
		assert.Equal(t, 9, list.Pop())
		assert.Equal(t, 8, list.Peek())

		assert.Equal(t, 8, list.Size())
		assert.Equal(t, 8, list.Pop())
		assert.Equal(t, 7, list.Peek())
	})

	t.Run("Should correct Unshift", func(t *testing.T) {
		list := doublelist.New[int]()
		list.Unshift()
		assert.Equal(t, list.Size(), 0)

		list.Unshift(1, 2)
		list.Unshift(3)
		assert.Equal(t, list.Size(), 3)
	})

	t.Run("Should correct Shift", func(t *testing.T) {
		list := initList()
		assert.Equal(t, 9, list.Size())
		assert.Equal(t, 1, list.Shift())

		assert.Equal(t, 8, list.Size())
		assert.Equal(t, 2, list.Shift())

		list = doublelist.New[int]()
		var tmp int
		assert.Equal(t, 0, list.Size())
		assert.Equal(t, tmp, list.Shift())
		assert.Equal(t, 0, list.Size())
	})

	t.Run("Should correct IsEmpty", func(t *testing.T) {
		list := doublelist.New[int]()
		list.Push(1)
		assert.False(t, list.IsEmpty())
		list.Pop()
		assert.True(t, list.IsEmpty())
	})

	t.Run("Should correct Head, Tail", func(t *testing.T) {
		list := initList()
		assert.Equal(t, 9, list.Tail().Value())
		assert.Equal(t, 1, list.Head().Value())

		list.Pop()
		list.Pop()
		list.Shift()
		list.Shift()

		assert.Equal(t, 7, list.Tail().Value())
		assert.Equal(t, 3, list.Head().Value())
	})

	t.Run("Should correct String", func(t *testing.T) {
		list := doublelist.New[int]()
		list.Push(1)
		list.Push(2)
		list.Push(3)
		assert.Equal(t, "1 2 3", list.String())
	})

	t.Run("Should correct Clear", func(t *testing.T) {
		list := initList()
		assert.Equal(t, 9, list.Size())
		assert.False(t, list.IsEmpty())

		list.Clear()
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
	})

	t.Run("Should correct copy", func(t *testing.T) {
		list := initList()
		copyList := list.Copy()
		assert.NotEqual(t, fmt.Sprintf("%p", list.Head()), fmt.Sprintf("%p", copyList.Head()))
		assert.Equal(t, list.Size(), copyList.Size())
		assert.Equal(t, "1 2 3 4 5 6 7 8 9", list.String())
		assert.Equal(t, "1 2 3 4 5 6 7 8 9", copyList.String())
	})

	t.Run("Should correct GetNode", func(t *testing.T) {
		list := doublelist.New[int]()
		list.Push(1)
		list.Push(2)
		list.Push(3)
		node, err := list.GetNode(1)
		assert.Nil(t, err)
		assert.Equal(t, node.Value(), 2)
		node, err = list.GetNode(0)
		assert.Nil(t, err)
		assert.Equal(t, node.Value(), 1)
		node, err = list.GetNode(2)
		assert.Nil(t, err)
		assert.Equal(t, node.Value(), 3)

		node, err = list.GetNode(3)
		assert.Nil(t, node)
		assert.Equal(t, doublelist.ErrOutOfRange, err)
		node, err = list.GetNode(-1)
		assert.Nil(t, node)
		assert.Equal(t, doublelist.ErrOutOfRange, err)
	})

	t.Run("Should correct Get", func(t *testing.T) {
		list := doublelist.New[int]()
		list.Push(1)
		list.Push(2)
		list.Push(3)
		node, err := list.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, node, 2)
		node, err = list.Get(0)
		assert.Nil(t, err)
		assert.Equal(t, node, 1)
		node, err = list.Get(2)
		assert.Nil(t, err)
		assert.Equal(t, node, 3)

		var nilRes int
		node, err = list.Get(3)
		assert.Equal(t, nilRes, node)
		assert.Equal(t, doublelist.ErrOutOfRange, err)
		node, err = list.Get(-1)
		assert.Equal(t, nilRes, node)
		assert.Equal(t, doublelist.ErrOutOfRange, err)
	})

	t.Run("Insert method testing", func(t *testing.T) {
		t.Run("Should error", func(t *testing.T) {
			list := initList()

			err := list.Insert(11, 1)
			assert.Equal(t, doublelist.ErrOutOfRange, err)
			err = list.Insert(-1, 1)
			assert.Equal(t, doublelist.ErrOutOfRange, err)
		})

		t.Run("Should insert first index", func(t *testing.T) {
			list := initList()
			currentSize := list.Size()
			cases := []struct {
				newItem int
				size    int
			}{
				{123, currentSize + 1},
				{321, currentSize + 2},
				{0, currentSize + 3},
			}
			for _, item := range cases {
				err := list.Insert(0, item.newItem)
				assert.Nil(t, err)
				assert.Equal(t, list.Size(), item.size)
				assert.Equal(t, list.Head().Value(), item.newItem)
			}
		})

		t.Run("Should insert last", func(t *testing.T) {
			list := initList()
			currentSize := list.Size()
			cases := []struct {
				newItem int
				size    int
			}{
				{123, currentSize + 1},
				{321, currentSize + 2},
				{0, currentSize + 3},
			}
			for _, item := range cases {
				lastIndex := list.Size()
				err := list.Insert(lastIndex, item.newItem)
				assert.Nil(t, err)
				assert.Equal(t, item.size, list.Size())
				assert.Equal(t, list.Tail().Value(), item.newItem)
			}
		})

		t.Run("Should insert mid", func(t *testing.T) {
			list := initList()
			currentSize := list.Size()
			cases := []struct {
				index    int
				newValue int
				size     int
				expect   string
			}{
				{4, 123, currentSize + 1, "1 2 3 4 123 5 6 7 8 9"},
				{6, 2325, currentSize + 2, "1 2 3 4 123 5 2325 6 7 8 9"},
				{1, 4325, currentSize + 3, "1 4325 2 3 4 123 5 2325 6 7 8 9"},
				{0, 5325, currentSize + 4, "5325 1 4325 2 3 4 123 5 2325 6 7 8 9"},
				{currentSize + 4, 6325, currentSize + 5, "5325 1 4325 2 3 4 123 5 2325 6 7 8 9 6325"},
			}
			for _, item := range cases {
				err := list.Insert(item.index, item.newValue)
				assert.Nil(t, err)
				assert.Equal(t, list.Size(), item.size)
				assert.Equal(t, item.expect, list.String())
			}
		})
	})

	t.Run("Should correct update", func(t *testing.T) {
		list := doublelist.New[int]()
		list.Push(1)
		list.Push(2)
		list.Push(3)
		node, err := list.GetNode(1)
		assert.Nil(t, err)
		assert.Equal(t, node.Value(), 2)
		node, err = list.GetNode(0)
		assert.Nil(t, err)
		assert.Equal(t, node.Value(), 1)
		node, err = list.GetNode(2)
		assert.Nil(t, err)
		assert.Equal(t, node.Value(), 3)

		err = list.Update(3, 1)
		assert.Equal(t, doublelist.ErrOutOfRange, err)
		err = list.Update(-1, 1)
		assert.Equal(t, doublelist.ErrOutOfRange, err)

		node, err = list.GetNode(1)
		assert.Nil(t, err)
		err = list.Update(1, node.Value()*3)
		assert.Nil(t, err)
		assert.Equal(t, node.Value(), 2*3)
		node, err = list.GetNode(0)
		assert.Nil(t, err)
		err = list.Update(0, node.Value()*3)
		assert.Nil(t, err)
		assert.Equal(t, node.Value(), 1*3)
		node, err = list.GetNode(2)
		assert.Nil(t, err)
		err = list.Update(2, node.Value()*3)
		assert.Nil(t, err)
		assert.Equal(t, node.Value(), 3*3)
	})

	t.Run("Should correct remove", func(t *testing.T) {
		list := doublelist.New[int]()
		list.Push(1)
		list.Push(2)
		list.Push(3)
		assert.Equal(t, 3, list.Size())

		err := list.Remove(-1)
		assert.Equal(t, doublelist.ErrOutOfRange, err)
		err = list.Remove(3)
		assert.Equal(t, doublelist.ErrOutOfRange, err)

		err = list.Remove(2)
		assert.Nil(t, err)
		assert.Equal(t, 2, list.Size())
		assert.Equal(t, 2, list.Tail().Value())

		err = list.Remove(0)
		assert.Nil(t, err)
		assert.Equal(t, 1, list.Size())
		assert.Equal(t, 2, list.Head().Value())

		err = list.Remove(0)
		assert.Nil(t, err)
		assert.Equal(t, 0, list.Size())
		assert.Nil(t, list.Head())
		assert.Nil(t, list.Tail())

		list.Push(1)
		list.Push(2)
		list.Push(3)
		err = list.Remove(1)
		assert.Nil(t, err)
		assert.Equal(t, 2, list.Size())
		assert.Equal(t, 3, list.Head().Next().Value())
		assert.Equal(t, 1, list.Tail().Prev().Value())

		list.Push(4)
		list.Push(5)
		list.Push(6)
		err = list.Remove(2)
		assert.Nil(t, err)
		assert.Equal(t, 4, list.Size())
		assert.Equal(t, "1 3 5 6", list.String())
	})
}
