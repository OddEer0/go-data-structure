package treemap

import (
	"cmp"

	"github.com/OddEer0/go-data-structure/container"
	"github.com/OddEer0/go-data-structure/tree/redblacktree"
)

type (
	Iterator[T any, K any] struct {
		iterator *redblacktree.Iterator[T, K]
	}

	Map[T any, K any] interface {
		container.Container
		Copy() Map[T, K]

		Add(key T, value K)  // O(log(n))
		Remove(key T) bool   // O(log(n))
		Get(key T) (K, bool) // O(log(n))

		Values() []K
		Keys() []T

		Iterator() *Iterator[T, K]

		container.EnumWithKey[T, K, Map[T, K]]
	}

	treeMap[T any, K any] struct {
		tree redblacktree.Tree[T, K]
	}
)

func (t *treeMap[T, K]) String() string {
	return t.tree.String()
}

func (t *treeMap[T, K]) Copy() Map[T, K] {
	return &treeMap[T, K]{
		tree: t.tree.Copy(),
	}
}

func (t *treeMap[T, K]) Add(key T, value K) {
	t.tree.InsertOrUpdate(key, value)
}

func (t *treeMap[T, K]) Clear() {
	t.tree.Clear()
}

func (t *treeMap[T, K]) Get(key T) (K, bool) {
	return t.tree.Get(key)
}

func (t *treeMap[T, K]) IsEmpty() bool {
	return t.tree.IsEmpty()
}

func (t *treeMap[T, K]) Keys() []T {
	return t.tree.Keys()
}

func (t *treeMap[T, K]) Remove(key T) bool {
	return t.tree.Remove(key)
}

func (t *treeMap[T, K]) Size() int {
	return t.Size()
}

func (t *treeMap[T, K]) Values() []K {
	return t.tree.Values()
}

func New[T cmp.Ordered, K any]() Map[T, K] {
	return &treeMap[T, K]{
		tree: redblacktree.New[T, K](),
	}
}

func NewWith[T any, K any](comparator func(a, b T) int) Map[T, K] {
	return &treeMap[T, K]{
		tree: redblacktree.NewWith[T, K](comparator),
	}
}
