package linked_list

type (
	tCallback[T any]         func(T, int, *linkedList[T])
	tCallbackR[T any, R any] func(T, int, *linkedList[T]) R

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
		Copy() *linkedList[T]
		ForEach(callback tCallback[T])
		Map(callback tCallbackR[T, T]) *linkedList[T]
		Some(callback tCallbackR[T, bool]) bool
		Every(callback tCallbackR[T, bool]) bool
		Filter(callback tCallbackR[T, bool]) *linkedList[T]
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
