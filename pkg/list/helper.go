package list

func (l *List[T]) getNode(index int) *Node[T] {
	if index >= l.length || index < 0 {
		return nil
	}

	node := l.head

	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

func (l *List[T]) getHead() *Node[T] {
	return l.head
}
