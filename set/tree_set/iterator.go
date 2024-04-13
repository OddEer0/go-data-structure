package treeset

func (s *set[T, K]) Iterator() *Iterator[T, K] {
	return &Iterator[T, K]{
		iterator: s.tree.Iterator(),
	}
}

func (i *Iterator[T, K]) Next() bool {
	return i.iterator.Next()
}

func (i *Iterator[T, K]) Prev() bool {
	return i.iterator.Prev()
}

func (i *Iterator[T, K]) Value() T {
	return i.iterator.Key()
}

func (i *Iterator[T, K]) Start() {
	i.iterator.Start()
}

func (i *Iterator[T, K]) End() {
	i.iterator.End()
}

func (i *Iterator[T, K]) First() bool {
	i.iterator.Start()
	return i.iterator.Next()
}

func (i *Iterator[T, K]) Last() bool {
	i.iterator.End()
	return i.iterator.Prev()
}
