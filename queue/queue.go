package queue

import "github.com/OddEer0/go-data-structure/container"

type Queue[T any] interface {
	container.Container
	Unshift(item T)
	Shift() T
	Peek() T
}
