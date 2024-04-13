package stack

import (
	"github.com/OddEer0/go-data-structure/container"
)

type Stack[T any] interface {
	container.Container
	Push(value ...T)
	Pop() T
	Peek() T
}
