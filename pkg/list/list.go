package list

type (
	Args[T any] struct {
		Item  T
		Index int
		List  IList[T]
	}
	AccArgs[T any, K any] struct {
		Acc   K
		Item  T
		Index int
		List  IList[T]
	}
	tCallback[T any]          func(Args[T])
	tCallbackR[T, R any]      func(Args[T]) R
	tCallbackReduce[T, K any] func(args AccArgs[T, K]) K

	IList[T any] interface {
		GetHeadNode() *Node[T]
		GetTailNode() *Node[T]
		Push(item T)              // O(1)
		Pop() T                   // O(1)
		Unshift(item T)           // O(1)
		Shift() T                 // O(1)
		Get(index int) T          // O(n)
		Set(index int, item T)    // O(n)
		Insert(index int, item T) // O(n)
		Size() int                // O(1)
		Remove(index int)         // O(n)
		Copy() IList[T]           // O(n)
		ForEach(callback tCallback[T])
		ForEachRight(callback tCallback[T])
		Map(callback tCallbackR[T, T]) IList[T]
		Some(callback tCallbackR[T, bool]) bool
		Every(callback tCallbackR[T, bool]) bool
		Filter(callback tCallbackR[T, bool]) IList[T]
		Find(callback tCallbackR[T, bool]) T
		FindLast(callback tCallbackR[T, bool]) T
		FindIndex(callback tCallbackR[T, bool]) int
		FindIndexLast(callback tCallbackR[T, bool]) int
		Reduce(callback tCallbackReduce[T, interface{}], init interface{}) interface{}
		ReduceRight(callback tCallbackReduce[T, interface{}], init interface{}) interface{}
		// IsEmpty
		// Clear
		// Has
		// Search
		// SearchLast
		// IndexOf
		// LastIndexOf
		// Reverse
		// ToReversed
		// Concat
		// Sort
		// ToSorted
		// Slice
		// Splice
		// ToSpliced
		// As
	}

	Node[T any] struct {
		value T
		next  *Node[T]
		prev  *Node[T]
	}

	List[T any] struct {
		head   *Node[T]
		tail   *Node[T]
		length int
	}
)

func NewLinkedList[T any]() IList[T] {
	return &List[T]{nil, nil, 0}
}
