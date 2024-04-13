package doublelist

import (
	"errors"

	"github.com/OddEer0/go-data-structure/container"
)

type position byte

const (
	start, process, end position = 0, 1, 2
)

var (
	ErrOutOfRange               = errors.New("out of range")
	ErrEndIndexLessOrEqualStart = errors.New("end index less or equal start")
)

type (
	Iterator[T any] struct {
		*list[T]
		index   int
		current *Node[T]
		position
	}

	Node[T any] struct {
		value T
		next  *Node[T]
		prev  *Node[T]
	}

	List[T any] interface {
		container.Container
		Copy() List[T]
		Head() *Node[T]
		Tail() *Node[T]

		Iterator() *Iterator[T]

		container.EnumWithIndex[T, List[T]]

		Push(value ...T)
		Peek() T
		Pop() T
		Unshift(value ...T)
		Shift() T

		Insert(index int, item T) error
		Remove(index int) error
		Get(index int) (T, error)
		GetNode(index int) (*Node[T], error)
		Update(index int, item T) error

		As(index int) (T, error)
		Find(func(index int, item T) bool) (T, bool)
		FindIndex(func(index int, item T) bool) int
		Reduce(callback func(acc interface{}, index int, item T) interface{}, init interface{}) interface{}
		Contains(item T) bool
		Search(item T) (T, bool)
		IndexOf(item T) int
		LastIndexOf(item T) int
		Swap(first int, second int) error
		Reverse()
		ToReversed() List[T]
		Slice(start int, end int) (List[T], error)
	}

	list[T any] struct {
		head   *Node[T]
		tail   *Node[T]
		length int
	}
)

func New[T any](init ...T) List[T] {
	l := &list[T]{nil, nil, 0}
	l.Push(init...)
	return l
}
