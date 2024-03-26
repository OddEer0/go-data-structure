package doublelist

func (l *list[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{
		list:     l,
		index:    0,
		current:  nil,
		position: start,
	}
}

func (i *Iterator[T]) Next() bool {
	switch i.position {
	case end:
		return false
	case start:
		if i.list.length == 0 {
			return false
		}
		i.current = i.list.head
		i.position = process
		i.index++
		return true
	default:
		i.current = i.current.next
		if i.current == nil {
			i.position = end
			return false
		}
		i.index++
		return true
	}
}

func (i *Iterator[T]) Prev() bool {
	switch i.position {
	case start:
		return false
	case end:
		if i.list.length == 0 {
			return false
		}
		i.current = i.list.tail
		i.position = process
		i.index--
		return true
	default:
		i.current = i.current.prev
		if i.current == nil {
			i.position = start
			return false
		}
		i.index--
		return true
	}
}

func (i *Iterator[T]) Value() T {
	return i.current.value
}

func (i *Iterator[T]) Index() int {
	return i.index
}

func (i *Iterator[T]) Node() *Node[T] {
	return i.current
}

func (i *Iterator[T]) Start() {
	i.current = nil
	i.index = 0
	i.position = start
}

func (i *Iterator[T]) End() {
	i.current = nil
	i.index = i.list.Size() - 1
	i.position = end
}

func (i *Iterator[T]) First() bool {
	i.Start()
	return i.Next()
}

func (i *Iterator[T]) Last() bool {
	i.End()
	return i.Prev()
}
