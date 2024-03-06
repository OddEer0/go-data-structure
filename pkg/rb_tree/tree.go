package rbtree

const (
	black = true
	red   = false
)

func ShallowGreater[T Comparable](a, b T) bool {
	return a > b
}

type Comparable interface {
	int | string
}

type Node[T Comparable, K any] struct {
	key                 T
	value               K
	left, right, parent *Node[T, K]
	color               bool
}

type Tree[T Comparable, K any] struct {
	root   *Node[T, K]
	length int
	cmp    func(T, T) bool
}

type ITree[T Comparable, K any] interface {
	ChangeCmpFunc(fn func(a, b T) bool)
	GetRoot() *Node[T, K]              // O(1)
	GetSize() int                      // O(1)s
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

func NilNode[T Comparable, K any]() *Node[T, K] {
	return &Node[T, K]{
		left:   nil,
		right:  nil,
		parent: nil,
		color:  black,
	}
}

var instance interface{} = nil

func nilNode[T Comparable, K any]() *Node[T, K] {
	if instance == nil {
		instance = NilNode[T, K]()
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

func NewRBTree[T Comparable, K any]() ITree[T, K] {
	return &Tree[T, K]{
		root:   nil,
		length: 0,
		cmp:    ShallowGreater[T],
	}
}
