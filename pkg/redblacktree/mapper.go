package redblacktree

func (t *RedBlackTree[T, K]) Values() []K {
	res := make([]K, 0, t.length)

	t.InOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.value)
	})

	return res
}

func (t *RedBlackTree[T, K]) Keys() []T {
	res := make([]T, 0, t.length)

	t.InOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.key)
	})

	return res
}

func (t *RedBlackTree[T, K]) Entries() []*Entry[T, K] {
	res := make([]*Entry[T, K], 0, t.length)

	t.InOrderFunc(func(node *Node[T, K]) {
		res = append(res, &Entry[T, K]{node.key, node.value})
	})

	return res
}

func (t *RedBlackTree[T, K]) PreOrderValues() []K {
	res := make([]K, 0, t.length)

	t.PreOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.value)
	})

	return res
}

func (t *RedBlackTree[T, K]) PostOrderValues() []K {
	res := make([]K, 0, t.length)

	t.PostOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.value)
	})

	return res
}

func (t *RedBlackTree[T, K]) PreOrderKeys() []T {
	res := make([]T, 0, t.length)

	t.PreOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.key)
	})

	return res
}

func (t *RedBlackTree[T, K]) PostOrderKeys() []T {
	res := make([]T, 0, t.length)

	t.PostOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.key)
	})

	return res
}

func (t *RedBlackTree[T, K]) PreOrderEntries() []*Entry[T, K] {
	res := make([]*Entry[T, K], 0, t.length)

	t.PreOrderFunc(func(node *Node[T, K]) {
		res = append(res, &Entry[T, K]{node.key, node.value})
	})

	return res
}

func (t *RedBlackTree[T, K]) PostOrderEntries() []*Entry[T, K] {
	res := make([]*Entry[T, K], 0, t.length)

	t.PostOrderFunc(func(node *Node[T, K]) {
		res = append(res, &Entry[T, K]{node.key, node.value})
	})

	return res
}
