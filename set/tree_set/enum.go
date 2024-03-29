package treeset

import "github.com/OddEer0/go-data-structure/tree/redblacktree"

func (s *set[T, K]) Each(callback func(index int, value T)) {
	iterator := s.Iterator()
	i := 0
	for iterator.Next() {
		callback(i, iterator.Value())
		i++
	}
}

func (s *set[T, K]) EachLast(callback func(index int, value T)) {
	iterator := s.Iterator()
	iterator.End()
	i := s.Size() - 1
	for iterator.Prev() {
		callback(i, iterator.Value())
		i--
	}
}

func (s *set[T, K]) Some(callback func(index int, value T) bool) bool {
	iterator := s.Iterator()
	i := 0
	for iterator.Next() {
		if callback(i, iterator.Value()) {
			return true
		}
		i++
	}

	return false
}

func (s *set[T, K]) Every(callback func(index int, value T) bool) bool {
	iterator := s.Iterator()
	i := 0
	for iterator.Next() {
		if !callback(i, iterator.Value()) {
			return false
		}
		i++
	}

	return true
}

func (s *set[T, K]) Map(callback func(index int, value T) T) Set[T, K] {
	tree := redblacktree.NewWith[T, bool](s.cmp)
	iterator := s.Iterator()
	i := 0
	for iterator.Next() {
		tree.Insert(callback(i, iterator.Value()), false)
		i++
	}
	return &set[T, K]{tree: tree, cmp: s.cmp}
}

func (s *set[T, K]) Filter(callback func(index int, value T) bool) Set[T, K] {
	tree := redblacktree.NewWith[T, bool](s.cmp)
	iterator := s.Iterator()
	i := 0
	for iterator.Next() {
		if callback(i, iterator.Value()) {
			tree.Insert(iterator.Value(), false)
		}
		i++
	}
	return &set[T, K]{tree: tree, cmp: s.cmp}
}

func (s *set[T, K]) Concat(others ...Set[T, K]) Set[T, K] {
	tree := redblacktree.NewWith[T, bool](s.cmp)
	for _, otherSet := range others {
		iterator := otherSet.Iterator()
		for iterator.Next() {
			tree.Insert(iterator.Value(), false)
		}
	}
	return &set[T, K]{tree: tree, cmp: s.cmp}
}
