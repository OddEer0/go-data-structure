package linked_list

func (l *linkedList[T]) Push(item T) {
	newNode := &node[T]{item, nil, nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}

	l.length++
}

func (l *linkedList[T]) Unshift(item T) {
	newNode := &node[T]{item, nil, nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
	}

	l.length++
}

func (l *linkedList[T]) Pop() T {
	if l.head == nil {
		var res T
		return res
	}

	result := l.tail.value
	prev := l.tail.prev

	if prev != nil {
		prev.next = nil
		l.tail = prev
	} else {
		l.head = nil
		l.tail = nil
	}

	l.length--

	return result
}

func (l *linkedList[T]) Shift() T {
	if l.head == nil {
		var res T
		return res
	}

	result := l.head.value
	next := l.head.next

	if next != nil {
		next.prev = nil
		l.head = next
	} else {
		l.head = nil
		l.tail = nil
	}

	l.length--

	return result
}

func (l *linkedList[T]) Size() int {
	return l.length
}

func (l *linkedList[T]) Get(index int) T {
	if result := l.getNode(index); result != nil {
		return result.value
	}
	var res T
	return res
}

func (l *linkedList[T]) Remove(index int) {
	if index < 0 || index >= l.length {
		return
	}

	if index == 0 {
		l.Shift()
	} else if index == l.length-1 {
		l.Pop()
	} else {
		current := l.head.next
		for i := 1; i < index; i++ {
			current = current.next
		}
		current.prev.next = current.next
		current.next.prev = current.prev
		l.length--
	}
}

func (l *linkedList[T]) Copy() *linkedList[T] {
	newList := &linkedList[T]{}
	l.ForEach(func(item T, _ int, _ *linkedList[T]) {
		newList.Push(item)
	})
	return newList
}

func (l *linkedList[T]) Insert(index int, item T) {
	if index > l.length || index < 0 {
		return
	}

	if index == 0 {
		l.Unshift(item)
	} else if index == l.length {
		l.Push(item)
	} else {
		newNode := &node[T]{item, nil, nil}
		current := l.getNode(index - 1)
		newNode.next = current.next
		current.next.prev = newNode
		newNode.prev = current
		current.next = newNode
		l.length++
	}
}

func (l *linkedList[T]) Set(index int, item T) {
	l.getNode(index).value = item
}
