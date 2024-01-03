package list

import "reflect"

func (l *List[T]) IsEmpty() bool {
	if l.Size() == 0 {
		return true
	}
	return false
}

func (l *List[T]) Clean() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

func (l *List[T]) Has(item T) bool {
	current := l.head
	for current != nil {
		if reflect.DeepEqual(item, current.value) {
			return true
		}
		current = current.next
	}
	return false
}

func (l *List[T]) Search(item T) T {
	current := l.head
	for current != nil {
		if reflect.DeepEqual(item, current.value) {
			return current.value
		}
		current = current.next
	}
	var res T
	return res
}

func (l *List[T]) SearchLast(item T) T {
	current := l.tail
	for current != nil {
		if reflect.DeepEqual(item, current.value) {
			return current.value
		}
		current = current.prev
	}
	var res T
	return res
}

func (l *List[T]) IndexOf(item T) int {
	current := l.head
	i := 0
	for current != nil {
		if reflect.DeepEqual(item, current.value) {
			return i
		}
		i++
		current = current.next
	}
	return -1
}

func (l *List[T]) LastIndexOf(item T) int {
	current := l.tail
	i := l.Size() - 1
	for current != nil {
		if reflect.DeepEqual(item, current.value) {
			return i
		}
		i--
		current = current.prev
	}
	return -1
}

func (l *List[T]) Swap(first int, second int) {
	if first >= l.Size() || second >= l.Size() || first < 0 || second < 0 {
		return
	}
	tmp := l.Get(first)
	l.Set(first, l.Get(second))
	l.Set(second, tmp)
}

func (l *List[T]) Reverse() {
	for i := 0; i < l.Size()/2; i++ {
		l.Swap(i, l.Size()-i-1)
	}
}

func (l *List[T]) ToReversed() IList[T] {
	newList := l.Copy()
	for i := 0; i < newList.Size()/2; i++ {
		newList.Swap(i, l.Size()-i-1)
	}
	return newList
}

func (l *List[T]) Concat(list IList[T]) IList[T] {
	newList := l.Copy()
	list.ForEach(func(args Args[T]) {
		newList.Push(args.Item)
	})
	return newList
}

func (l *List[T]) Slice(start int, end int) IList[T] {
	if start < 0 || end < 0 || start >= l.Size() || end >= l.Size() {
		return nil
	}
	newList := NewLinkedList[T]()
	for i := start; i <= end; i++ {
		newList.Push(l.Get(i))
	}
	return newList
}

func (l *List[T]) As(index int) T {
	correctIndex := index
	if correctIndex < 0 {
		correctIndex = l.Size() + correctIndex
	}
	if correctIndex < 0 || correctIndex >= l.Size() {
		var res T
		return res
	}
	return l.Get(correctIndex)
}
