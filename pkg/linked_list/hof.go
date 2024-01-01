package linked_list

func (l *linkedList[T]) ForEach(callback func(T, int, *linkedList[T])) {
	current := l.head
	for i := 0; i < l.length; i++ {
		callback(current.value, i, l)
		current = current.next
	}
}

func (l *linkedList[T]) Map(callback func(T, int, *linkedList[T]) T) *linkedList[T] {
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

func (l *linkedList[T]) Some(callback func(T, int, *linkedList[T]) bool) bool {
	current := l.head
	for i := 0; i < l.length; i++ {
		if callback(current.value, i, l) {
			return true
		}
		current = current.next
	}
	return false
}
