package arraystack

import (
	"fmt"
	stackInterface "github.com/OddEer0/go-data-structure/stack"
	"strings"
)

type (
	Stack[T any] interface {
		stackInterface.Stack[T]
		Cap() int
		Clip()
		ClearWithSize(size int)
		Copy() Stack[T]
	}

	stack[T any] struct {
		arr []T
	}
)

func (s *stack[T]) Size() int {
	return len(s.arr)
}

func (s *stack[T]) Cap() int {
	return cap(s.arr)
}

func (s *stack[T]) Clear() {
	s.arr = []T{}
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *stack[T]) String() string {
	var str strings.Builder

	for i := range s.arr {
		if i != 0 {
			str.WriteRune(' ')
		}
		str.WriteString(fmt.Sprintf("%v", s.arr[i]))
	}

	return str.String()
}

func (s *stack[T]) Push(value ...T) {
	if len(value) == 0 {
		return
	}
	s.arr = append(s.arr, value...)
}

func (s *stack[T]) Pop() T {
	if len(s.arr) > 0 {
		val := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		return val
	}
	var res T
	return res
}

func (s *stack[T]) Peek() T {
	return s.arr[len(s.arr)-1]
}

func (s *stack[T]) Clip() {
	s.arr = s.arr[:len(s.arr):len(s.arr)]
}

func (s *stack[T]) ClearWithSize(size int) {
	s.arr = make([]T, 0, size)
}

func (s *stack[T]) Copy() Stack[T] {
	sl := make([]T, 0, cap(s.arr))
	for i := range s.arr {
		sl = append(sl, s.arr[i])
	}
	return &stack[T]{arr: sl}
}

func New[T any]() Stack[T] {
	return &stack[T]{arr: []T{}}
}

func NewWithSize[T any](size int) Stack[T] {
	return &stack[T]{arr: make([]T, 0, size)}
}
