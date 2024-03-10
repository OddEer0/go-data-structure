package tests

import (
	"testing"

	llist "github.com/OddEer0/go-data-structure/list"
)

func TestLinkedListHofMethods(t *testing.T) {
	t.Run("Should ForEach method correct working", func(t *testing.T) {
		list := FirstMockList()
		i := 0
		list.ForEach(func(args llist.Args[int]) {
			if args.Item != list.Get(i) {
				t.Errorf("fail incorrect item")
			}
			if args.Index != i {
				t.Errorf("fail incorrect index")
			}
			if list != args.List {
				t.Errorf("callback list not equal list")
			}
			i++
		})
	})

	t.Run("Should ForEachRight method correct working", func(t *testing.T) {
		list := FirstMockList()
		i := list.Size() - 1
		list.ForEachRight(func(args llist.Args[int]) {
			if args.Item != list.Get(i) {
				t.Errorf("fail incorrect item")
			}
			if args.Index != i {
				t.Errorf("fail incorrect index")
			}
			if list != args.List {
				t.Errorf("callback list not equal list")
			}
			i--
		})
	})

	t.Run("Map method testing", func(t *testing.T) {
		t.Run("Should return new list", func(t *testing.T) {
			list := FirstMockList()
			newList := list.Map(func(args llist.Args[int]) int {
				return args.Item
			})
			if list == newList {
				t.Errorf("fail return old list")
			}
		})

		t.Run("Should correct callback arguments", func(t *testing.T) {
			list := FirstMockList()
			i := 0
			list.Map(func(args llist.Args[int]) int {
				if args.Item != list.Get(i) {
					t.Errorf("incorrect item: %v, expected %v", args.Item, list.Get(i))
				}
				if args.Index != i {
					t.Errorf("fail incorrect index: %v, expected %v", args.Index, i)
				}
				if list != args.List {
					t.Errorf("callback list not equal list")
				}
				i++
				return args.Item
			})
		})

		t.Run("Should correct data mapping", func(t *testing.T) {
			list := FirstMockList()
			newList := list.Map(func(args llist.Args[int]) int {
				return args.Item + 1
			})
			list.ForEach(func(args llist.Args[int]) {
				if args.Item+1 != newList.Get(args.Index) {
					t.Errorf("incorrect data mapping: %v, expected %v", newList.Get(args.Index), args.Item+1)
				}
			})
		})
	})

	t.Run("Some method testing", func(t *testing.T) {
		t.Run("Should correct callback arguments", func(t *testing.T) {
			list := FirstMockList()
			i := 0
			list.Some(func(args llist.Args[int]) bool {
				if args.Item != list.Get(i) {
					t.Errorf("incorrect item: %v, expected %v", args.Item, list.Get(i))
				}
				if args.Index != i {
					t.Errorf("incorrect index: %v, expected %v", args.Index, i)
				}
				if list != args.List {
					t.Errorf("callback list not equal list")
				}
				i++
				return false
			})
		})

		t.Run("Should correct working", func(t *testing.T) {
			list := FirstMockList()
			cases := []struct {
				output bool
				search int
			}{
				{true, 9},
				{true, 5},
				{true, 2},
				{false, 1213},
				{false, 321},
			}

			for _, item := range cases {
				result := list.Some(func(args llist.Args[int]) bool {
					if args.Item == item.search {
						return true
					}
					return false
				})
				if result != item.output {
					t.Errorf("incorrect return value: %v, expected: %v", result, item.output)
				}
			}
		})
	})

	t.Run("Every method testing", func(t *testing.T) {
		t.Run("Should correct callback arguments", func(t *testing.T) {
			list := FirstMockList()
			i := 0
			list.Every(func(args llist.Args[int]) bool {
				if args.Item != list.Get(i) {
					t.Errorf("incorrect item: %v, expected %v", args.Item, list.Get(i))
				}
				if args.Index != i {
					t.Errorf("incorrect index: %v, expected %v", args.Index, i)
				}
				if list != args.List {
					t.Errorf("callback list not equal list")
				}
				i++
				return true
			})
		})

		t.Run("Should correct working", func(t *testing.T) {
			list := FirstMockList()
			cases := []struct {
				output bool
				ld     int
			}{
				{true, 10},
				{true, 12},
				{true, 32},
				{false, 5},
				{false, 4},
			}
			for _, item := range cases {
				result := list.Every(func(args llist.Args[int]) bool {
					if args.Item < item.ld {
						return true
					}
					return false
				})
				if result != item.output {
					t.Errorf("incorrect return value: %v, expected %v", result, item.output)
				}
			}
		})
	})

	t.Run("Filter method testing", func(t *testing.T) {
		t.Run("Should return new list", func(t *testing.T) {
			list := FirstMockList()
			newList := list.Filter(func(args llist.Args[int]) bool {
				return true
			})
			if list == newList {
				t.Errorf("fail return old list")
			}
		})

		t.Run("Should correct callback arguments", func(t *testing.T) {
			list := FirstMockList()
			i := 0
			list.Filter(func(args llist.Args[int]) bool {
				if args.Item != list.Get(i) {
					t.Errorf("incorrect item: %v, expected %v", args.Item, list.Get(i))
				}
				if args.Index != i {
					t.Errorf("incorrect index: %v, expected %v", args.Index, i)
				}
				if list != args.List {
					t.Errorf("callback list not equal list")
				}
				i++
				return true
			})
		})

		t.Run("Should correct working", func(t *testing.T) {
			list := FirstMockList()
			cases := []struct {
				size int
				ld   int
				gt   int
			}{
				{9, 10, 0},
				{8, 10, 1},
				{4, 10, 5},
				{0, 100, 10},
				{2, 3, 0},
			}
			for _, item := range cases {
				newList := list.Filter(func(args llist.Args[int]) bool {
					if args.Item < item.ld && args.Item > item.gt {
						return true
					}
					return false
				})
				if item.size != newList.Size() {
					t.Errorf("incorrect filterng. size: %v, expected %v", newList.Size(), item.size)
				}
			}
		})
	})

	t.Run("Find method testing", func(t *testing.T) {
		t.Run("Should correct callback arguments", func(t *testing.T) {
			list := FirstMockList()
			i := 0
			list.Find(func(args llist.Args[int]) bool {
				if args.Item != list.Get(i) {
					t.Errorf("incorrect item: %v, expected %v", args.Item, list.Get(i))
				}
				if args.Index != i {
					t.Errorf("incorrect index: %v, expected %v", args.Index, i)
				}
				if list != args.List {
					t.Errorf("callback list not equal list")
				}
				i++
				return true
			})
		})

		t.Run("Should correct working", func(t *testing.T) {
			list := FirstMockList()
			cases := []struct {
				output int
				search int
			}{
				{5, 5},
				{1, 1},
				{9, 9},
				{search: 100},
			}
			for _, item := range cases {
				result := list.Find(func(args llist.Args[int]) bool {
					if args.Item == item.search {
						return true
					}
					return false
				})
				if result != item.output {
					t.Errorf("incorrect return value: %v, expect %v", result, item.output)
				}
			}
		})
	})

	t.Run("FindIndex method testing", func(t *testing.T) {
		t.Run("Should correct callback arguments", func(t *testing.T) {
			list := FirstMockList()
			i := 0
			list.FindIndex(func(args llist.Args[int]) bool {
				if args.Item != list.Get(i) {
					t.Errorf("incorrect item: %v, expected %v", args.Item, list.Get(i))
				}
				if args.Index != i {
					t.Errorf("incorrect index: %v, expected %v", args.Index, i)
				}
				if list != args.List {
					t.Errorf("callback list not equal list")
				}
				i++
				return true
			})
		})

		t.Run("Should correct working", func(t *testing.T) {
			list := FirstMockList()
			list.Push(9)
			list.Push(5)
			cases := []struct {
				output int
				search int
			}{
				{4, 5},
				{0, 1},
				{8, 9},
				{-1, 100},
				{-1, 11},
			}
			for _, item := range cases {
				result := list.FindIndex(func(args llist.Args[int]) bool {
					if args.Item == item.search {
						return true
					}
					return false
				})
				if result != item.output {
					t.Errorf("incorrect return value: %v, expect %v", result, item.output)
				}
			}
		})
	})

	t.Run("FindLast method testing", func(t *testing.T) {
		t.Run("Should correct callback arguments", func(t *testing.T) {
			list := FirstMockList()
			i := list.Size() - 1
			list.FindLast(func(args llist.Args[int]) bool {
				if args.Item != list.Get(i) {
					t.Errorf("incorrect item: %v, expected %v", args.Item, list.Get(i))
				}
				if args.Index != i {
					t.Errorf("incorrect index: %v, expected %v", args.Index, i)
				}
				if list != args.List {
					t.Errorf("callback list not equal list")
				}
				i--
				return true
			})
		})

		t.Run("Should correct working", func(t *testing.T) {
			list := FirstMockList()
			cases := []struct {
				output int
				search int
			}{
				{5, 5},
				{1, 1},
				{9, 9},
				{search: 100},
			}
			for _, item := range cases {
				result := list.FindLast(func(args llist.Args[int]) bool {
					if args.Item == item.search {
						return true
					}
					return false
				})
				if result != item.output {
					t.Errorf("incorrect return value: %v, expect %v", result, item.output)
				}
			}
		})
	})

	t.Run("FindLastIndex method testing", func(t *testing.T) {
		t.Run("Should correct callback arguments", func(t *testing.T) {
			list := FirstMockList()
			i := list.Size() - 1
			list.FindIndexLast(func(args llist.Args[int]) bool {
				if args.Item != list.Get(i) {
					t.Errorf("incorrect item: %v, expected %v", args.Item, list.Get(i))
				}
				if args.Index != i {
					t.Errorf("incorrect index: %v, expected %v", args.Index, i)
				}
				if list != args.List {
					t.Errorf("callback list not equal list")
				}
				i--
				return true
			})
		})

		t.Run("Should correct working", func(t *testing.T) {
			list := FirstMockList()
			list.Push(9)
			list.Push(5)
			cases := []struct {
				output int
				search int
			}{
				{10, 5},
				{0, 1},
				{9, 9},
				{-1, 100},
			}
			for _, item := range cases {
				result := list.FindIndexLast(func(args llist.Args[int]) bool {
					if args.Item == item.search {
						return true
					}
					return false
				})
				if result != item.output {
					t.Errorf("incorrect return value: %v, expect %v", result, item.output)
				}
			}
		})
	})
	// TODO - reduce tests
	t.Run("Reduce method testing", func(t *testing.T) {
		t.Run("Should correct callback arguments", func(t *testing.T) {
			list := FirstMockList()
			i := 0
			list.Reduce(func(args llist.AccArgs[int, interface{}]) interface{} {
				if args.Acc != 1 {
					t.Errorf("incorrect accumulate value: %v, expected: %v", args.Acc, 1)
				}
				if args.Item != list.Get(i) {
					t.Errorf("incorrect item: %v, expected %v", args.Item, list.Get(i))
				}
				if args.Index != i {
					t.Errorf("incorrect index: %v, expected %v", args.Index, i)
				}
				if list != args.List {
					t.Errorf("callback list not equal list")
				}
				i++
				return args.Acc
			}, 1)
		})

		t.Run("Should correct working", func(t *testing.T) {
			list := FirstMockList()
			cases := []struct {
				maxIteration int
				output       int
			}{
				{9, 45},
				{5, 15},
				{6, 21},
			}
			for _, item := range cases {
				result := list.Reduce(func(args llist.AccArgs[int, interface{}]) interface{} {
					if args.Index < item.maxIteration {
						return args.Acc.(int) + args.Item
					}
					return args.Acc
				}, 0)
				if result != item.output {
					t.Errorf("incorrect result: %v, expected %v", result, item.output)
				}
			}
		})
	})
	// TODO - reduceRight tests
	t.Run("ReduceRight method testing", func(t *testing.T) {
		t.Run("Should correct callback arguments", func(t *testing.T) {
			list := FirstMockList()
			i := list.Size() - 1
			list.ReduceRight(func(args llist.AccArgs[int, interface{}]) interface{} {
				if args.Acc != 1 {
					t.Errorf("incorrect accumulate value: %v, expected: %v", args.Acc, 1)
				}
				if args.Item != list.Get(i) {
					t.Errorf("incorrect item: %v, expected %v", args.Item, list.Get(i))
				}
				if args.Index != i {
					t.Errorf("incorrect index: %v, expected %v", args.Index, i)
				}
				if list != args.List {
					t.Errorf("callback list not equal list")
				}
				i--
				return args.Acc
			}, 1)
		})

		t.Run("Should correct working", func(t *testing.T) {
			list := FirstMockList()
			cases := []struct {
				maxIteration int
				output       int
			}{
				{9, 45},
				{5, 35},
				{6, 39},
			}
			for _, item := range cases {
				result := list.ReduceRight(func(args llist.AccArgs[int, interface{}]) interface{} {
					if args.List.Size()-1-args.Index < item.maxIteration {
						return args.Acc.(int) + args.Item
					}
					return args.Acc
				}, 0)
				if result != item.output {
					t.Errorf("incorrect result: %v, expected %v", result, item.output)
				}
			}
		})
	})
}
