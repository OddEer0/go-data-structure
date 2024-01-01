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
