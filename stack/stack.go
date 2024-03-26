package stack

import list "github.com/OddEer0/go-data-structure/list/double_list"

type IStack[T any] interface {
	Peek() T
	Push(value T)
	Pop() T
	Size() int
}

func NewStack[T any]() IStack[T] {
	return list.New[T]()
}
