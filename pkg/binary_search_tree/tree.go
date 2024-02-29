package binarysearchtree

type Comparable interface {
	int | string
}

func ShallowGreater[T Comparable](a, b T) bool {
	return a > b
}

type Node[T Comparable, K any] struct {
	key         T
	value       K
	left, right *Node[T, K]
}

type Tree[T Comparable, K any] struct {
	root   *Node[T, K]
	length int
	cmp    func(T, T) bool
}

type ITree[T Comparable, K any] interface {
	ChangeCmpFunc(fn func(a, b T) bool)
	GetRoot() *Node[T, K]       // O(1)
	GetSize() int               // O(1)s
	Insert(key T, value K) bool // O(log(n))

	PreOrderFunc(callback func(*Node[T, K]))
	InOrderFunc(callback func(*Node[T, K]))
	PostOrderFunc(callback func(*Node[T, K]))

	ToSortedSlice() []K
}

func NewBSTree[T Comparable, K any]() ITree[T, K] {
	return &Tree[T, K]{
		root:   nil,
		length: 0,
		cmp:    ShallowGreater[T],
	}
}
