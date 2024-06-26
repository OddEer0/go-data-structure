package liststack

import (
	list "github.com/OddEer0/go-data-structure/list/double_list"
	stackInterface "github.com/OddEer0/go-data-structure/stack"
)

type (
	Stack[T any] interface {
		stackInterface.Stack[T]
		Copy() Stack[T]
	}

	stack[T any] struct {
		list list.List[T]
	}
)

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

func (s *stack[T]) Push(value ...T) {
	s.list.Push(value...)
}

func (s *stack[T]) Pop() T {
	return s.list.Pop()
}

func (s *stack[T]) Peek() T {
	return s.list.Peek()
}

func (s *stack[T]) Copy() Stack[T] {
	return &stack[T]{list: s.list.Copy()}
}

func New[T any]() Stack[T] {
	return &stack[T]{list: list.New[T]()}
}
