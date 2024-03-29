package treeset

import (
	"cmp"
	"fmt"
	"github.com/OddEer0/go-data-structure/container"
	"github.com/OddEer0/go-data-structure/tree/redblacktree"
	"strings"
)

type (
	Iterator[T any, K any] struct {
		iterator *redblacktree.Iterator[T, bool]
	}

	Set[T any, K any] interface {
		container.Container
		Copy() Set[T, K]
		container.EnumWithIndex[T, Set[T, K]]

		Add(items T) bool
		Remove(items T) bool
		Contains(items ...T) bool

		Iterator() *Iterator[T, K]

		Intersection(another Set[T, K]) Set[T, K]
		Union(another Set[T, K]) Set[T, K]
		Difference(another Set[T, K]) Set[T, K]
		SymmetricDifference(another Set[T, K]) Set[T, K]
		IsDisjoinFrom(another Set[T, K]) bool
		IsSubsetOf(another Set[T, K]) bool
		IsSupersetOf(another Set[T, K]) bool

		Values() []T
	}

	set[T any, K any] struct {
		cmp  func(a, b T) int
		tree redblacktree.Tree[T, bool]
	}
)

func (s *set[T, K]) Size() int {
	return s.tree.Size()
}

func (s *set[T, K]) Clear() {
	s.tree.Clear()
}

func (s *set[T, K]) IsEmpty() bool {
	return s.tree.IsEmpty()
}

func (s *set[T, K]) String() string {
	iterator := s.tree.Iterator()
	var str strings.Builder
	isFirst := true
	for iterator.Next() {
		if isFirst {
			isFirst = false
		} else {
			str.WriteRune(' ')
		}
		str.WriteString(fmt.Sprintf("%v", iterator.Key()))
	}
	return str.String()
}

func (s *set[T, K]) Copy() Set[T, K] {
	return &set[T, K]{tree: s.tree.Copy(), cmp: s.cmp}
}

func (s *set[T, K]) Add(item T) bool {
	if s.tree.Insert(item, false) == nil {
		return true
	}
	return false
}

func (s *set[T, K]) Remove(item T) bool {
	return s.tree.Remove(item)
}

func (s *set[T, K]) Contains(items ...T) bool {
	for _, item := range items {
		_, has := s.tree.Get(item)
		if !has {
			return false
		}
	}
	return true
}

func (s *set[T, K]) Intersection(another Set[T, K]) Set[T, K] {
	newSet := &set[T, K]{
		tree: redblacktree.NewWith[T, bool](s.cmp),
		cmp:  s.cmp,
	}

	iterator := another.Iterator()
	for iterator.Next() {
		if s.Contains(iterator.Value()) {
			newSet.Add(iterator.Value())
		}
	}

	return newSet
}

func (s *set[T, K]) Union(another Set[T, K]) Set[T, K] {
	newSet := s.Copy()

	iterator := another.Iterator()
	for iterator.Next() {
		newSet.Add(iterator.Value())
	}

	return newSet
}

func (s *set[T, K]) Difference(another Set[T, K]) Set[T, K] {
	newSet := s.Copy()

	iterator := another.Iterator()
	for iterator.Next() {
		newSet.Remove(iterator.Value())
	}

	return newSet
}

func (s *set[T, K]) SymmetricDifference(another Set[T, K]) Set[T, K] {
	newSet := s.Copy()

	iterator := another.Iterator()
	for iterator.Next() {
		if s.Contains(iterator.Value()) {
			newSet.Remove(iterator.Value())
		} else {
			newSet.Add(iterator.Value())
		}
	}

	return newSet
}

func (s *set[T, K]) IsDisjoinFrom(another Set[T, K]) bool {
	iterator := another.Iterator()
	for iterator.Next() {
		if s.Contains(iterator.Value()) {
			return false
		}
	}

	return true
}

func (s *set[T, K]) IsSubsetOf(another Set[T, K]) bool {
	iterator := s.Iterator()
	for iterator.Next() {
		if !another.Contains(iterator.Value()) {
			return false
		}
	}

	return true
}

func (s *set[T, K]) IsSupersetOf(another Set[T, K]) bool {
	iterator := another.Iterator()
	for iterator.Next() {
		if !s.Contains(iterator.Value()) {
			return false
		}
	}

	return true
}

func (s *set[T, K]) Values() []T {
	return s.tree.Keys()
}

func New[T cmp.Ordered](init ...T) Set[T, bool] {
	return &set[T, bool]{tree: redblacktree.New[T, bool](), cmp: cmp.Compare[T]}
}

func NewWith[T any](comparator func(a, b T) int, init ...T) Set[T, bool] {
	return &set[T, bool]{tree: redblacktree.NewWith[T, bool](comparator), cmp: comparator}
}
