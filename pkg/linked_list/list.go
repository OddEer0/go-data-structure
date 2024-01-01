package linked_list

type (
	LinkedList[T any] interface {
		Push(item T)              // O(1)
		Pop() T                   // O(1)
		Unshift(item T)           // O(1)
		Shift() T                 // O(1)
		Get(index int) T          // O(n)
		Set(index int, item T)    // O(n)
		Insert(index int, item T) // O(n)
		Size() int                // O(1)
		Remove(index int)         // O(n)
		ForEach(callback func(T, int, *linkedList[T]))
		Copy() *linkedList[T]
		Map(callback func(T, int, *linkedList[T]) T) *linkedList[T]
		Some(callback func(T, int, *linkedList[T]) bool) bool
	}

	node[T any] struct {
		value T
		next  *node[T]
		prev  *node[T]
	}

	linkedList[T any] struct {
		head   *node[T]
		tail   *node[T]
		length int
	}
)

func NewLinkedList[T any]() LinkedList[T] {
	return &linkedList[T]{nil, nil, 0}
}
