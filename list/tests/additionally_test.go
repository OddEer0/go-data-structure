package tests

import (
	"testing"

	llist "github.com/OddEer0/go-data-structure/list"
)

func TestAdditionallyListMethods(t *testing.T) {
	t.Run("Should IsEmpty method correct working", func(t *testing.T) {
		list := FirstMockList()
		list2 := llist.NewLinkedList[int]()
		if list2.IsEmpty() != true {
			t.Errorf("incorrect return value: %v, expected %v", list2.IsEmpty(), true)
		}
		if list.IsEmpty() != false {
			t.Errorf("incorrect return value: %v, expected %v", list.IsEmpty(), false)
		}
	})

	t.Run("Should Clean method correct working", func(t *testing.T) {
		list := FirstMockList()
		list.Clean()
		if list.Size() != 0 || list.GetHeadNode() != nil || list.GetTailNode() != nil {
			t.Errorf("incorrect clean list")
		}
	})

	t.Run("Should Has method correct working", func(t *testing.T) {
		list := FirstMockList()
		cases := []struct {
			output bool
			search int
		}{
			{true, 6},
			{true, 2},
			{true, 9},
			{false, 110},
			{false, 50},
		}
		for _, item := range cases {
			result := list.Has(item.search)
			if result != item.output {
				t.Errorf("incorrect return value: %v, expected %v", result, item.output)
			}
		}
	})

	t.Run("Should Search method correct working", func(t *testing.T) {
		list := FirstMockList()
		cases := []struct {
			search int
			output int
		}{
			{5, 5},
			{9, 9},
			{search: 100},
		}
		for _, item := range cases {
			result := list.Search(item.search)
			if result != item.output {
				t.Errorf("incorrect search list element: %v, expected %v", result, item.output)
			}
		}
	})

	t.Run("Should SearchLast method correct working", func(t *testing.T) {
		list := FirstMockList()
		cases := []struct {
			search int
			output int
		}{
			{5, 5},
			{9, 9},
			{search: 100},
		}
		for _, item := range cases {
			result := list.SearchLast(item.search)
			if result != item.output {
				t.Errorf("incorrect search list element: %v, expected %v", result, item.output)
			}
		}
	})

	t.Run("Should IndexOf method correct working", func(t *testing.T) {
		list := FirstMockList()
		list.Unshift(9)
		cases := []struct {
			search int
			output int
		}{
			{5, 5},
			{9, 0},
			{8, 8},
			{100, -1},
		}
		for _, item := range cases {
			result := list.IndexOf(item.search)
			if result != item.output {
				t.Errorf("incorrect index list element: %v, expected %v", result, item.output)
			}
		}
	})

	t.Run("Should LastIndexOf method correct working", func(t *testing.T) {
		list := FirstMockList()
		list.Unshift(9)
		cases := []struct {
			search int
			output int
		}{
			{5, 5},
			{9, 9},
			{8, 8},
			{100, -1},
		}
		for _, item := range cases {
			result := list.LastIndexOf(item.search)
			if result != item.output {
				t.Errorf("incorrect index list element: %v, expected %v", result, item.output)
			}
		}
	})

	t.Run("Should Swap method correct working", func(t *testing.T) {
		list := FirstMockList()
		list.Swap(0, 8)
		if list.Get(0) != 9 && list.Get(8) != 1 {
			t.Errorf("incorrect swap")
		}
	})

	t.Run("Should Reverse method correct working", func(t *testing.T) {
		list := FirstMockList()
		oldState := list.Copy()
		list.Reverse()
		i := oldState.Size() - 1
		list.ForEach(func(args llist.Args[int]) {
			if args.Item != oldState.Get(i) {
				t.Errorf("Should incorrect reverse list")
			}
			i--
		})
	})

	t.Run("Should ToReverse method correct working", func(t *testing.T) {
		t.Run("Should return new list", func(t *testing.T) {
			list := FirstMockList()
			newList := list.ToReversed()
			if list == newList {
				t.Errorf("return old list, expected new list")
			}
		})

		t.Run("Should correct working", func(t *testing.T) {
			list := FirstMockList()
			oldState := list.Copy()
			list.Reverse()
			i := oldState.Size() - 1
			list.ForEach(func(args llist.Args[int]) {
				if args.Item != oldState.Get(i) {
					t.Errorf("Should incorrect reverse list")
				}
				i--
			})
		})
	})

	t.Run("Should Concat method correct working", func(t *testing.T) {
		t.Run("Should return new list", func(t *testing.T) {
			list := FirstMockList()
			list2 := list.Copy()
			newList := list.Concat(list2)
			if list == newList {
				t.Errorf("return old list, expected new list")
			}
		})

		t.Run("Should correct working", func(t *testing.T) {
			list := FirstMockList()
			newAddList := list.Map(func(args llist.Args[int]) int {
				return args.Item + 9
			}).Filter(func(args llist.Args[int]) bool {
				if args.Index < 3 {
					return true
				}
				return false
			})
			newList := list.Concat(newAddList)
			i := newList.Size() - 1
			if newList.Size() != 12 || newList.Get(i) != 12 || newList.Get(i-1) != 11 || newList.Get(i-2) != 10 {
				t.Errorf("incorrect concat list")
			}
		})
	})

	t.Run("Should Slice method correct working", func(t *testing.T) {
		t.Run("Should return new list", func(t *testing.T) {
			list := FirstMockList()
			newList := list.Slice(0, 4)
			if list == newList {
				t.Errorf("return old list, expected new list")
			}
		})

		t.Run("Should correct working", func(t *testing.T) {
			list := FirstMockList()
			badSlice := list.Slice(-1, 10)
			if badSlice != nil {
				t.Errorf("incorrect error handleing, expected nil")
			}
			slice := list.Slice(1, 5)
			if slice.Size() != 5 {
				t.Errorf("incorrect slice")
			}
			items := []int{2, 3, 4, 5, 6}
			i := 0
			slice.ForEach(func(args llist.Args[int]) {
				if items[i] != slice.Get(i) {
					t.Errorf("incorrect slice list")
				}
				i++
			})
		})
	})

	t.Run("Should As method correct working", func(t *testing.T) {
		list := FirstMockList()
		cases := []struct {
			output int
			index  int
		}{
			{1, 0},
			{9, 8},
			{5, 4},
			{9, -1},
			{1, -9},
			{index: -10},
			{index: 10},
		}
		for _, item := range cases {
			if list.As(item.index) != item.output {
				t.Errorf("incorrect get As value: %v, expected %v", list.As(item.index), item.output)
			}
		}
	})
}
