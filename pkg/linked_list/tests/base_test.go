package tests

import (
	"github.com/OddEer0/go-data-structure/pkg/linked_list"
	"testing"
)

func FirstMockList() linked_list.LinkedList[int] {
	list := linked_list.NewLinkedList[int]()

	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	list.Push(5)
	list.Push(6)
	list.Push(7)
	list.Push(8)
	list.Push(9)
	return list
}

func TestBaseLinkedList(t *testing.T) {
	t.Run("Should init linked list", func(t *testing.T) {
		list := linked_list.NewLinkedList[int]()
		if list == nil {
			t.Errorf("fail init list")
		}
	})

	t.Run("Should push and get items", func(t *testing.T) {
		list := linked_list.NewLinkedList[int]()
		list.Push(3)
		list.Push(4)
		list.Unshift(2)
		list.Unshift(1)
		cases := []struct {
			result int
			real   int
		}{
			{1, list.Get(0)},
			{2, list.Get(1)},
			{3, list.Get(2)},
			{4, list.Get(3)},
		}
		for _, item := range cases {
			if item.real != item.result {
				t.Errorf("fail push, unshift, get")
			}
		}
	})

	t.Run("Should shift and unshift items", func(t *testing.T) {
		list := FirstMockList()
		cases := []struct {
			returnResult     int
			realReturnResult int
			resultList       int
			realResultList   int
		}{
			{9, list.Pop(), 8, list.Get(7)},
			{8, list.Pop(), 7, list.Get(6)},
			{7, list.Pop(), 6, list.Get(5)},
			{1, list.Shift(), 2, list.Get(0)},
			{2, list.Shift(), 3, list.Get(0)},
			{3, list.Shift(), 4, list.Get(0)},
		}
		for _, item := range cases {
			if item.resultList != item.realResultList && item.returnResult != item.realReturnResult {
				t.Errorf("fail shift unshift method")
			}
		}
	})

	t.Run("Should get list size", func(t *testing.T) {
		list := FirstMockList()
		if list.Size() != 9 {
			t.Errorf("incorrect get list size")
		}
		list.Push(10)
		if list.Size() != 10 {
			t.Errorf("incorrect get list size")
		}
		list.Pop()
		if list.Size() != 9 {
			t.Errorf("incorrect get list size")
		}
	})

	t.Run("Should remove list item", func(t *testing.T) {
		list := FirstMockList()
		cases := []struct {
			removeIndex      int
			currentItemIndex int
			size             int
		}{
			{4, 6, 8},
			{6, 9, 7},
			{1, 3, 6},
		}
		for _, item := range cases {
			list.Remove(item.removeIndex)
			if list.Size() != item.size && list.Get(item.removeIndex) != item.currentItemIndex {
				t.Errorf("fail remove method")
			}
		}
	})

	t.Run("Should set new value", func(t *testing.T) {
		list := FirstMockList()
		cases := []struct {
			index    int
			newValue int
		}{
			{4, 7},
			{2, 7},
			{7, 123},
		}
		for _, item := range cases {
			list.Set(item.index, item.newValue)
			if item.newValue != list.Get(item.index) {
				t.Errorf("fail set method")
			}
		}
	})

	t.Run("Insert method testing", func(t *testing.T) {
		t.Run("Should insert first index", func(t *testing.T) {
			list := FirstMockList()
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
				prev := list.Get(0)
				list.Insert(0, item.newItem)
				if list.Size() != item.size {
					t.Errorf("fail incorrect list size")
				}
				if list.Get(0) != item.newItem && list.Get(1) != prev {
					t.Errorf("fail incorrect new item insert")
				}
			}
		})

		t.Run("Should insert last", func(t *testing.T) {
			list := FirstMockList()
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
				prev := list.Get(lastIndex - 1)
				list.Insert(lastIndex, item.newItem)
				if list.Size() != item.size {
					t.Errorf("fail incorrect list size")
				}
				if list.Get(lastIndex) != item.newItem && list.Get(lastIndex-1) != prev {
					t.Errorf("fail incorrect new item insert")
				}
			}
		})

		t.Run("Should insert mid", func(t *testing.T) {
			list := FirstMockList()
			currentSize := list.Size()
			cases := []struct {
				index    int
				newValue int
				size     int
			}{
				{4, 123, currentSize + 1},
				{6, 2325, currentSize + 2},
				{1, 4325, currentSize + 3},
			}
			for _, item := range cases {
				prevItemIndex := list.Get(item.index)
				list.Insert(item.index, item.newValue)
				if list.Size() != item.size {
					t.Errorf("fail incorrect list size")
				}
				if list.Get(item.index) != item.newValue && list.Get(item.index+1) != prevItemIndex {
					t.Errorf("fail incorrect new item insert")
				}
			}
		})
	})
}
