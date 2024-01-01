package linked_list

func (l *linkedList[T]) ForEach(callback tCallback[T]) {
	current := l.head
	for i := 0; i < l.length; i++ {
		callback(current.value, i, l)
		current = current.next
	}
}

func (l *linkedList[T]) Map(callback tCallbackR[T, T]) *linkedList[T] {
	newList := l.Copy()
	current := l.head
	newListCurrent := newList.head
	for i := 0; i < l.length; i++ {
		newListCurrent.value = callback(current.value, i, l)
		current = current.next
		newListCurrent = newListCurrent.next
	}
	return newList
}

func (l *linkedList[T]) Some(callback tCallbackR[T, bool]) bool {
	current := l.head
	for i := 0; i < l.length; i++ {
		if callback(current.value, i, l) {
			return true
		}
		current = current.next
	}
	return false
}

func (l *linkedList[T]) Every(callback tCallbackR[T, bool]) bool {
	current := l.head
	for i := 0; i < l.length; i++ {
		if !callback(current.value, i, l) {
			return false
		}
		current = current.next
	}
	return true
}

func (l *linkedList[T]) Filter(callback tCallbackR[T, bool]) *linkedList[T] {
	newList := &linkedList[T]{}

	l.ForEach(func(item T, index int, list *linkedList[T]) {
		if callback(item, index, list) {
			newList.Push(item)
		}
	})

	return newList
}

func (l *linkedList[T]) Find(callback tCallbackR[T, bool]) T {
	current := l.head
	for i := 0; i < l.length; i++ {
		if callback(current.value, i, l) {
			return current.value
		}
		current = current.next
	}
	var res T
	return res
}

func (l *linkedList[T]) FindLast(callback tCallbackR[T, bool]) T {
	current := l.tail
	for i := l.length - 1; i >= 0; i-- {
		if callback(current.value, i, l) {
			return current.value
		}
		current = current.prev
	}
	var res T
	return res
}

func (l *linkedList[T]) FindIndex(callback tCallbackR[T, bool]) int {
	current := l.head
	for i := 0; i < l.length; i++ {
		if callback(current.value, i, l) {
			return i
		}
		current = current.next
	}
	return -1
}

func (l *linkedList[T]) FindIndexLast(callback tCallbackR[T, bool]) int {
	current := l.tail
	for i := l.length - 1; i >= 0; i-- {
		if callback(current.value, i, l) {
			return i
		}
		current = current.prev
	}
	return -1
}
