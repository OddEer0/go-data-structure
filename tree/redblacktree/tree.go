package redblacktree

import (
	"cmp"

	"github.com/OddEer0/go-data-structure/container"
)

const (
	defaultMaxDepth               = 13
	defaultMaxDepthLimit          = 500
	relativeToDepthMul            = 1.2
	black, red                    = true, false
	start, process, end  position = 0, 1, 2
)

type (
	position byte

	Entry[T any, K any] struct {
		Key   T
		Value K
	}

	Node[T any, K any] struct {
		key                 T
		value               K
		left, right, parent *Node[T, K]
		color               bool
	}

	Iterator[T any, K any] struct {
		tree *RedBlackTree[T, K]
		node *Node[T, K]
		position
	}

	PreOrderIterator[T any, K any] struct {
		tree *RedBlackTree[T, K]
		node *Node[T, K]
		position
	}

	RedBlackTree[T any, K any] struct {
		root   *Node[T, K]
		length int
		cmp    func(T, T) int
	}

	Tree[T any, K any] interface {
		container.Container
		container.EnumWithKey[T, K, Tree[T, K]]

		Root() *Node[T, K] // O(1)
		Copy() Tree[T, K]
		// Right get max node in tree
		Right() *Node[T, K]
		// Left get min node in tree
		Left() *Node[T, K] // O(n) memory: O(2 * (relativeMaxDepth+1)) min memory 13

		Insert(key T, value K) *Node[T, K] // O(log(n))
		Remove(key T) bool                 // O(log(n))
		Update(key T, value K) bool        // O(log(n))
		GetNode(key T) (*Node[T, K], bool) // O(log(n))
		Get(key T) (K, bool)               // O(log(n))

		InsertOrUpdate(key T, value K)
		RemoveMany(keys ...T)
		InsertMany(entries ...*Entry[T, K])         // O(n * log(k)) k - size
		InsertOrUpdateMany(entries ...*Entry[T, K]) // O(n * log(k)) k - size

		Iterator() *Iterator[T, K]
		PreOrderIterator() *PreOrderIterator[T, K]

		PreOrderFunc(callback func(*Node[T, K]))  // recursive
		InOrderFunc(callback func(*Node[T, K]))   // recursive
		PostOrderFunc(callback func(*Node[T, K])) // recursive

		Values() []K                      // recursive
		PreOrderValues() []K              // recursive
		PostOrderValues() []K             // recursive
		Keys() []T                        // recursive
		PreOrderKeys() []T                // recursive
		PostOrderKeys() []T               // recursive
		Entries() []*Entry[T, K]          // recursive
		PreOrderEntries() []*Entry[T, K]  // recursive
		PostOrderEntries() []*Entry[T, K] // recursive
	}
)

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

// New TODO - add with hard cmp tests
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
