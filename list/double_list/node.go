package doublelist

import "fmt"

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.value)
}

func (n *Node[T]) Value() T {
	return n.value
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}
