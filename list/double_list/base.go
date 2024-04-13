package doublelist

import "strings"

func (l *list[T]) Head() *Node[T] {
	return l.head
}

func (l *list[T]) Tail() *Node[T] {
	return l.tail
}

func (l *list[T]) String() string {
	var str strings.Builder
	it := l.Iterator()
	isStart := true
	for it.Next() {
		if isStart {
			isStart = false
		} else {
			str.WriteRune(' ')
		}
		str.WriteString(it.Node().String())
	}
	return str.String()
}

func (l *list[T]) IsEmpty() bool {
	return l.length == 0
}

func (l *list[T]) Size() int {
	return l.length
}

func (l *list[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

func (l *list[T]) Push(items ...T) {
	for _, item := range items {
		newNode := &Node[T]{item, nil, nil}

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
}

func (l *list[T]) Peek() T {
	return l.tail.value
}

func (l *list[T]) Pop() T {
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

func (l *list[T]) Unshift(items ...T) {
	for _, item := range items {
		newNode := &Node[T]{item, nil, nil}

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
}

func (l *list[T]) Shift() T {
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

func (l *list[T]) Copy() List[T] {
	newList := &list[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
	it := l.Iterator()
	for it.Next() {
		newList.Push(it.Node().Value())
	}
	return newList
}

func (l *list[T]) Remove(index int) error {
	switch {
	case index < 0 || index >= l.length:
		return ErrOutOfRange
	case index == 0:
		l.Shift()
	case index == l.length-1:
		l.Pop()
	default:
		current := l.head.next
		for i := 1; i < index; i++ {
			current = current.next
		}
		current.prev.next = current.next
		current.next.prev = current.prev
		l.length--
	}
	return nil
}

func (l *list[T]) GetNode(index int) (*Node[T], error) {
	if index >= l.length || index < 0 {
		return nil, ErrOutOfRange
	}

	node := l.head

	for i := 0; i < index; i++ {
		node = node.next
	}

	return node, nil
}

func (l *list[T]) Get(index int) (T, error) {
	node, err := l.GetNode(index)
	if err != nil {
		var n T
		return n, err
	}

	return node.value, nil
}

func (l *list[T]) Insert(index int, item T) error {
	switch {
	case index > l.length || index < 0:
		return ErrOutOfRange
	case index == 0:
		l.Unshift(item)
	case index == l.length:
		l.Push(item)
	default:
		newNode := &Node[T]{item, nil, nil}
		current, _ := l.GetNode(index - 1)
		newNode.next = current.next
		current.next.prev = newNode
		newNode.prev = current
		current.next = newNode
		l.length++
	}

	return nil
}

func (l *list[T]) Update(index int, item T) error {
	node, err := l.GetNode(index)
	if err != nil {
		return err
	}
	node.value = item
	return nil
}
