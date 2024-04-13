package treemap

func (t *treeMap[T, K]) Each(callback func(key T, val K)) {
	t.tree.Each(callback)
}

func (t *treeMap[T, K]) EachLast(callback func(key T, val K)) {
	t.tree.EachLast(callback)
}

func (t *treeMap[T, K]) Some(callback func(key T, val K) bool) bool {
	return t.tree.Some(callback)
}

func (t *treeMap[T, K]) Every(callback func(key T, val K) bool) bool {
	return t.tree.Every(callback)
}

func (t *treeMap[T, K]) Map(callback func(key T, val K) (T, K)) Map[T, K] {
	newTree := t.tree.Map(callback)
	return &treeMap[T, K]{
		tree: newTree,
	}
}

func (t *treeMap[T, K]) Filter(callback func(key T, val K) bool) Map[T, K] {
	newTree := t.tree.Filter(callback)
	return &treeMap[T, K]{
		tree: newTree,
	}
}

func (t *treeMap[T, K]) Concat(others ...Map[T, K]) Map[T, K] {
	newTree := t.Copy()
	for _, other := range others {
		it := other.Iterator()
		for it.Next() {
			newTree.Add(it.Key(), it.Value())
		}
	}
	return newTree
}
