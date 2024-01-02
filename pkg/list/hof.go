package list

func (l *List[T]) ForEach(callback tCallback[T]) {
	current := l.head
	args := Args[T]{current.value, 0, l}
	for ; args.Index < l.length; args.Index++ {
		callback(args)
		current = current.next
		if current != nil {
			args.Item = current.value
		}
	}
}

func (l *List[T]) ForEachRight(callback tCallback[T]) {
	current := l.tail
	args := Args[T]{current.value, l.length - 1, l}
	for ; args.Index >= 0; args.Index-- {
		callback(args)
		current = current.prev
		if current != nil {
			args.Item = current.value
		}
	}
}

func (l *List[T]) Map(callback tCallbackR[T, T]) IList[T] {
	newList := l.Copy()
	current := l.head
	newListCurrent := newList.GetHeadNode()
	args := Args[T]{current.value, 0, l}
	for ; args.Index < l.length; args.Index++ {
		newListCurrent.value = callback(args)
		current = current.next
		if current != nil {
			args.Item = current.value
		}
		newListCurrent = newListCurrent.next
	}
	return newList
}

func (l *List[T]) Some(callback tCallbackR[T, bool]) bool {
	current := l.head
	args := Args[T]{current.value, 0, l}
	for ; args.Index < l.length; args.Index++ {
		if callback(args) {
			return true
		}
		current = current.next
		if current != nil {
			args.Item = current.value
		}
	}
	return false
}

func (l *List[T]) Every(callback tCallbackR[T, bool]) bool {
	current := l.head
	args := Args[T]{current.value, 0, l}
	for ; args.Index < l.length; args.Index++ {
		if !callback(args) {
			return false
		}
		current = current.next
		if current != nil {
			args.Item = current.value
		}
	}
	return true
}

func (l *List[T]) Filter(callback tCallbackR[T, bool]) IList[T] {
	newList := NewLinkedList[T]()

	l.ForEach(func(args Args[T]) {
		if callback(args) {
			newList.Push(args.Item)
		}
	})

	return newList
}

func (l *List[T]) Find(callback tCallbackR[T, bool]) T {
	current := l.head
	args := Args[T]{current.value, 0, l}
	for ; args.Index < l.length; args.Index++ {
		if callback(args) {
			return current.value
		}
		current = current.next
		if current != nil {
			args.Item = current.value
		}
	}
	var res T
	return res
}

func (l *List[T]) FindLast(callback tCallbackR[T, bool]) T {
	current := l.tail
	args := Args[T]{current.value, l.length - 1, l}
	for ; args.Index >= 0; args.Index-- {
		if callback(args) {
			return current.value
		}
		current = current.prev
		if current != nil {
			args.Item = current.value
		}
	}
	var res T
	return res
}

func (l *List[T]) FindIndex(callback tCallbackR[T, bool]) int {
	current := l.head
	args := Args[T]{current.value, 0, l}
	for ; args.Index < l.length; args.Index++ {
		if callback(args) {
			return args.Index
		}
		current = current.next
		if current != nil {
			args.Item = current.value
		}
	}
	return -1
}

func (l *List[T]) FindIndexLast(callback tCallbackR[T, bool]) int {
	current := l.tail
	args := Args[T]{current.value, l.length - 1, l}
	for ; args.Index >= 0; args.Index-- {
		if callback(args) {
			return args.Index
		}
		current = current.prev
		if current != nil {
			args.Item = current.value
		}
	}
	return -1
}

func (l *List[T]) Reduce(callback tCallbackReduce[T, interface{}], init interface{}) interface{} {
	acc := init
	l.ForEach(func(args Args[T]) {
		acc = callback(AccArgs[T, interface{}]{acc, args.Item, args.Index, args.List})
	})
	return acc
}

func (l *List[T]) ReduceRight(callback tCallbackReduce[T, interface{}], init interface{}) interface{} {
	acc := init
	l.ForEachRight(func(args Args[T]) {
		acc = callback(AccArgs[T, interface{}]{acc, args.Item, args.Index, args.List})
	})
	return acc
}
