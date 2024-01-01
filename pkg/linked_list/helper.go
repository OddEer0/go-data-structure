package linked_list

func (l *linkedList[T]) getNode(index int) *node[T] {
	if index >= l.length || index < 0 {
		return nil
	}

	node := l.head

	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}
