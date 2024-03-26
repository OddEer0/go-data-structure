package doublelist

import "reflect"

func (l *list[T]) As(index int) T {
	correctIndex := index
	if correctIndex < 0 {
		correctIndex = l.Size() + correctIndex
	}
	if correctIndex < 0 || correctIndex >= l.Size() {
		var res T
		return res
	}
	res, _ := l.Get(correctIndex)
	return res
}

func (l *list[T]) Find(callback func(index int, item T) bool) (T, bool) {
	it := l.Iterator()
	for it.Next() {
		if callback(it.Index(), it.Value()) {
			return it.Value(), true
		}
	}
	var res T
	return res, false
}

func (l *list[T]) FindIndex(callback func(index int, item T) bool) int {
	it := l.Iterator()
	for it.Next() {
		if callback(it.Index(), it.Value()) {
			return it.Index()
		}
	}
	return -1
}

func (l *list[T]) Reduce(callback func(index int, item T) interface{}, init interface{}) interface{} {
	it := l.Iterator()
	for it.Next() {
		init = callback(it.Index(), it.Value())
	}

	return init
}
func (l *list[T]) ReduceRight(callback func(index int, item T) interface{}, init interface{}) interface{} {
	it := l.Iterator()
	it.End()
	for it.Prev() {
		init = callback(it.Index(), it.Value())
	}

	return init
}

func (l *list[T]) Contains(item T) bool {
	it := l.Iterator()
	for it.Next() {
		if reflect.DeepEqual(item, it.Value()) {
			return true
		}
	}
	return false
}

func (l *list[T]) Search(item T) T {
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

func (l *list[T]) SearchLast(item T) T {
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

func (l *list[T]) IndexOf(item T) int {
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

func (l *list[T]) LastIndexOf(item T) int {
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

func (l *list[T]) Swap(first int, second int) error {
	tmp, err := l.Get(first)
	if err != nil {
		return err
	}
	sec, err := l.Get(second)
	if err != nil {
		return err
	}
	err = l.Update(first, sec)
	if err != nil {
		return err
	}
	err = l.Update(second, tmp)
	if err != nil {
		return err
	}
	return nil
}

func (l *list[T]) Reverse() {
	for i := 0; i < l.Size()/2; i++ {
		_ = l.Swap(i, l.Size()-i-1)
	}
}

func (l *list[T]) ToReversed() List[T] {
	newList := l.Copy()
	for i := 0; i < newList.Size()/2; i++ {
		_ = newList.Swap(i, l.Size()-i-1)
	}
	return newList
}

func (l *list[T]) Slice(start int, end int) (List[T], error) {
	if start < 0 || end < 0 || start >= l.Size() || end >= l.Size() {
		return nil, ErrOutOfRange
	}
	newList := New[T]()
	for i := start; i < end; i++ {
		elem, _ := l.Get(i)
		newList.Push(elem)
	}
	return newList, nil
}
