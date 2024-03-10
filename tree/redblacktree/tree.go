package redblacktree

import "cmp"

const (
	black, red                    = true, false
	start, proccess, end position = 0, 1, 2
)

type position byte

type Entry[T any, K any] struct {
	Key   T
	Value K
}

type Node[T any, K any] struct {
	key                 T
	value               K
	left, right, parent *Node[T, K]
	color               bool
}

type Iterator[T any, K any] struct {
	tree *RedBlackTree[T, K]
	node *Node[T, K]
	position
}

type RedBlackTree[T any, K any] struct {
	root   *Node[T, K]
	length int
	cmp    func(T, T) int
}

type Tree[T any, K any] interface {
	Root() *Node[T, K] // O(1)
	Size() int         // O(1)
	Clear()            // O(1)
	IsEmpty() bool     // O(1)

	Insert(key T, value K) *Node[T, K] // O(log(n))
	Remove(key T) bool                 // O(log(n))
	Update(key T, value K) bool        // O(log(n))
	GetNode(key T) (*Node[T, K], bool) // O(log(n))
	Get(key T) (K, bool)               // O(log(n))

	InsertOrUpdate(key T, value K)              // O(log(n))
	InsertMany(entries ...*Entry[T, K])         // O(n * log(k)) k - size
	InsertOrUpdateMany(entries ...*Entry[T, K]) // O(n * log(k)) k - size
	RemoveMany(keys ...T)                       // O(n * log(k)) k - size

	PreOrderFunc(callback func(*Node[T, K]))  // recursive
	InOrderFunc(callback func(*Node[T, K]))   // recursive
	PostOrderFunc(callback func(*Node[T, K])) // recursive

	Iterator() *Iterator[T, K]

	Values() []K                      // recursive
	PreOrderValues() []K              // recursive
	PostOrderValues() []K             // recursive
	Keys() []T                        // recursive
	PreOrderKeys() []T                // recursive
	PostOrderKeys() []T               // recursive
	Entries() []*Entry[T, K]          // recursive
	PreOrderEntries() []*Entry[T, K]  // recursive
	PostOrderEntries() []*Entry[T, K] // recursive

	Right() *Node[T, K]
	Left() *Node[T, K]
}

var instance interface{} = nil

func nilNode[T any, K any]() *Node[T, K] {
	if instance == nil {
		instance = &Node[T, K]{
			left:   nil,
			right:  nil,
			parent: nil,
			color:  black,
		}
	}
	return instance.(*Node[T, K])
}

func NewNode[T any, K any](key T, value K) *Node[T, K] {
	return &Node[T, K]{
		key,
		value,
		nilNode[T, K](),
		nilNode[T, K](),
		nilNode[T, K](),
		red,
	}
}

func New[T cmp.Ordered, K any]() Tree[T, K] {
	return &RedBlackTree[T, K]{
		root:   nil,
		length: 0,
		cmp:    cmp.Compare[T],
	}
}

func NewWith[T any, K any](comparator func(a, b T) int) Tree[T, K] {
	return &RedBlackTree[T, K]{
		root:   nil,
		length: 0,
		cmp:    comparator,
	}
}
