package linked_list

type (
	tCallback[T any]          func(T, int, *linkedList[T])
	tCallbackR[T, R any]      func(T, int, *linkedList[T]) R
	tCallbackReduce[T, K any] func(acc K, item T, index int, list *linkedList[T]) K

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
		Copy() *linkedList[T]     // O(n)
		ForEach(callback tCallback[T])
		ForEachRight(callback tCallback[T])
		Map(callback tCallbackR[T, T]) *linkedList[T]
		Some(callback tCallbackR[T, bool]) bool
		Every(callback tCallbackR[T, bool]) bool
		Filter(callback tCallbackR[T, bool]) *linkedList[T]
		Find(callback tCallbackR[T, bool]) T
		FindLast(callback tCallbackR[T, bool]) T
		FindIndex(callback tCallbackR[T, bool]) int
		FindIndexLast(callback tCallbackR[T, bool]) int
		Reduce(callback tCallbackReduce[T, interface{}], init interface{}) interface{}
		ReduceRight(callback tCallbackReduce[T, interface{}], init interface{}) interface{}
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
