package redblacktree

const (
	black = true
	red   = false
)

func ShallowGreater[T Comparable](a, b T) bool {
	return a > b
}

func ShallowLess[T Comparable](a, b T) bool {
	return a < b
}

type Comparable interface {
	int | string
}

type Entry[T Comparable, K any] struct {
	key   T
	value K
}

type Node[T Comparable, K any] struct {
	key                 T
	value               K
	left, right, parent *Node[T, K]
	color               bool
}

type RedBlackTree[T Comparable, K any] struct {
	root   *Node[T, K]
	length int
	cmp    func(T, T) bool
}

type Tree[T Comparable, K any] interface {
	Root() *Node[T, K] // O(1)
	Size() int         // O(1)s
	Clear()
	IsEmpty() bool

	Insert(key T, value K) *Node[T, K] // O(log(n))
	Remove(key T) bool                 // O(log(n))
	Update(key T, value K) bool        // O(log(n))
	GetNode(key T) (*Node[T, K], bool) // O(log(n))
	Get(key T) (K, bool)               // use GetNode
	InsertOrUpdate(key T, value K)
	InsertMany(entries ...Entry[T, K])
	InsertOrUpdateMany(entries ...Entry[T, K])
	RemoveMany(keys ...T)

	PreOrderFunc(callback func(*Node[T, K]))
	InOrderFunc(callback func(*Node[T, K]))
	PostOrderFunc(callback func(*Node[T, K]))

	// Iterator

	Values() []K
	PreOrderValues() []K
	PostOrderValues() []K
	Keys() []T
	PreOrderKeys() []T
	PostOrderKeys() []T
	Entries() []*Entry[T, K]
	PreOrderEntries() []*Entry[T, K]
	PostOrderEntries() []*Entry[T, K]
}

var instance interface{} = nil

func nilNode[T Comparable, K any]() *Node[T, K] {
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

func NewNode[T Comparable, K any](key T, value K) *Node[T, K] {
	return &Node[T, K]{
		key,
		value,
		nilNode[T, K](),
		nilNode[T, K](),
		nilNode[T, K](),
		red,
	}
}

func New[T Comparable, K any]() Tree[T, K] {
	return &RedBlackTree[T, K]{
		root:   nil,
		length: 0,
		cmp:    ShallowLess[T],
	}
}

func NewWith[T Comparable, K any](fn func(a, b T) bool) Tree[T, K] {
	return &RedBlackTree[T, K]{
		root:   nil,
		length: 0,
		cmp:    ShallowLess[T],
	}
}
