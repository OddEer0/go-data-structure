package doublelist

func (l *list[T]) Each(callback func(i int, value T)) {
	it := l.Iterator()
	for it.Next() {
		callback(it.Index(), it.Value())
	}
}

func (l *list[T]) EachLast(callback func(i int, value T)) {
	it := l.Iterator()
	it.End()
	for it.Prev() {
		callback(it.Index(), it.Value())
	}
}

func (l *list[T]) Some(callback func(i int, value T) bool) bool {
	it := l.Iterator()
	for it.Next() {
		if callback(it.Index(), it.Value()) {
			return true
		}
	}
	return false
}

func (l *list[T]) Every(callback func(i int, value T) bool) bool {
	it := l.Iterator()
	for it.Next() {
		if !callback(it.Index(), it.Value()) {
			return false
		}
	}
	return true
}

func (l *list[T]) Map(callback func(i int, value T) T) List[T] {
	newList := &list[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
	it := l.Iterator()
	for it.Next() {
		newList.Push(callback(it.Index(), it.Value()))
	}
	return newList
}

func (l *list[T]) Filter(callback func(i int, value T) bool) List[T] {
	newList := &list[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
	it := l.Iterator()
	for it.Next() {
		if callback(it.Index(), it.Value()) {
			newList.Push(it.Value())
		}
	}
	return newList
}

func (l *list[T]) Concat(other ...List[T]) List[T] {
	newList := l.Copy()
	for _, otherList := range other {
		it := otherList.Iterator()
		for it.Next() {
			newList.Push(it.Value())
		}
	}
	return newList
}
