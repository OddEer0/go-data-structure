package avltree

import "cmp"

type Node[T any, K any] struct {
	key         T
	value       K
	left, right *Node[T, K]
	height      int
}

type AvlTree[T any, K any] struct {
	root   *Node[T, K]
	length int
	cmp    func(T, T) int
}

type Tree[T any, K any] interface {
	Root() *Node[T, K]                 // O(1)
	Size() int                         // O(1)s
	Insert(key T, value K) *Node[T, K] // O(log(n))
	Remove(key T)                      // O(log(n))
	Update(key T, value K) bool        // O(log(n))
	GetNode(key T) (*Node[T, K], bool) // O(log(n))

	PreOrderFunc(callback func(*Node[T, K]))
	InOrderFunc(callback func(*Node[T, K]))
	PostOrderFunc(callback func(*Node[T, K]))

	ToSortedSlice() []K
	ToPreOrderNodeSlice() []K
	ToPostOrderNodeSlice() []K
}

func New[T cmp.Ordered, K any]() Tree[T, K] {
	return &AvlTree[T, K]{
		root:   nil,
		length: 0,
		cmp:    cmp.Compare[T],
	}
}

func NewWith[T any, K any](compare func(a, b T) int) Tree[T, K] {
	return &AvlTree[T, K]{
		root:   nil,
		length: 0,
		cmp:    compare,
	}
}
