package stack

import "github.com/OddEer0/go-data-structure/list"

type IStack[T any] interface {
	Peek() T
	Push(value T)
	Pop() T
	Size() int
}

func NewStack[T any]() IStack[T] {
	return list.NewLinkedList[T]()
}
