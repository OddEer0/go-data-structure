package circularqueue

import (
	"fmt"
	queueInterface "github.com/OddEer0/go-data-structure/queue"
	"strings"
)

type (
	Queue[T any] interface {
		queueInterface.Queue[T]
		// Copy() Queue[T]
	}

	queue[T any] struct {
		buffer             []T
		start, end, length int
	}
)

func (q *queue[T]) Size() int {
	return q.length
}

func (q *queue[T]) Clear() {
	q.start = 0
	q.end = 0
	q.length = 0
	q.buffer = make([]T, len(q.buffer))
}

func (q *queue[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *queue[T]) String() string {
	if q.length == 0 {
		return ""
	}
	var str strings.Builder

	i := q.start
	j := 0
	for j < q.length {
		if i != q.start {
			str.WriteRune(' ')
		}
		str.WriteString(fmt.Sprintf("%v", q.buffer[i]))
		if i == len(q.buffer)-1 {
			i = 0
		} else {
			i++
		}
		j++
	}

	return str.String()
}

func (q *queue[T]) Unshift(items ...T) {
	for i := 0; q.length < len(q.buffer) && i < len(items); i++ {
		q.buffer[q.end] = items[i]
		if q.end == len(q.buffer)-1 {
			q.end = 0
		} else {
			q.end++
		}
		q.length++
	}
}

func (q *queue[T]) Shift() T {
	if q.length == 0 {
		var res T
		return res
	}
	result := q.buffer[q.start]
	if q.start == len(q.buffer)-1 {
		q.start = 0
	} else {
		q.start++
	}
	q.length--
	return result
}

func (q *queue[T]) Peek() T {
	if q.length == 0 {
		var res T
		return res
	}
	return q.buffer[q.start]
}

func New[T any](size int) Queue[T] {
	return &queue[T]{buffer: make([]T, size)}
}
