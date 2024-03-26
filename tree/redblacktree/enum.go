package redblacktree

func (t *RedBlackTree[T, K]) Each(callback func(key T, val K)) {
	it := t.Iterator()
	for it.Next() {
		callback(it.Key(), it.Value())
	}
}

func (t *RedBlackTree[T, K]) EachLast(callback func(key T, val K)) {
	it := t.Iterator()
	it.End()
	for it.Prev() {
		callback(it.Key(), it.Value())
	}
}

func (t *RedBlackTree[T, K]) Some(callback func(key T, val K) bool) bool {
	it := t.Iterator()
	for it.Next() {
		if callback(it.Key(), it.Value()) {
			return true
		}
	}
	return false
}

func (t *RedBlackTree[T, K]) Every(callback func(key T, val K) bool) bool {
	it := t.Iterator()
	for it.Next() {
		if !callback(it.Key(), it.Value()) {
			return false
		}
	}
	return true
}

func (t *RedBlackTree[T, K]) Map(callback func(key T, val K) (T, K)) Tree[T, K] {
	newTree := &RedBlackTree[T, K]{
		root:   nil,
		cmp:    t.cmp,
		length: 0,
	}
	it := t.Iterator()
	for it.Next() {
		newTree.Insert(callback(it.Key(), it.Value()))
	}
	return newTree
}

func (t *RedBlackTree[T, K]) Filter(callback func(key T, val K) bool) Tree[T, K] {
	newTree := &RedBlackTree[T, K]{
		root:   nil,
		cmp:    t.cmp,
		length: 0,
	}
	it := t.Iterator()
	for it.Next() {
		if callback(it.Key(), it.Value()) {
			newTree.Insert(it.Key(), it.Value())
		}
	}
	return newTree
}

func (t *RedBlackTree[T, K]) Concat(others ...Tree[T, K]) Tree[T, K] {
	newTree := t.Copy()
	for _, other := range others {
		it := other.Iterator()
		for it.Next() {
			newTree.Insert(it.Key(), it.Value())
		}
	}
	return newTree
}
