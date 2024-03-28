package liststack

import (
	list "github.com/OddEer0/go-data-structure/list/double_list"
	stackInterface "github.com/OddEer0/go-data-structure/stack"
)

type stack[T any] struct {
	list list.List[T]
}

func (s *stack[T]) Size() int {
	return s.list.Size()
}

func (s *stack[T]) Clear() {
	s.list.Clear()
}

func (s *stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *stack[T]) String() string {
	return s.list.String()
}

func (s *stack[T]) Push(value T) {
	s.list.Push(value)
}

func (s *stack[T]) Pop() T {
	return s.list.Pop()
}

func (s *stack[T]) Peek() T {
	return s.list.Peek()
}

func New[T any]() stackInterface.Stack[T] {
	return &stack[T]{list: list.New[T]()}
}
