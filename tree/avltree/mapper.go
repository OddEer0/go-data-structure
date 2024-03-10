package avltree

func (t *AvlTree[T, K]) ToSortedSlice() []K {
	res := make([]K, 0, t.length)

	t.InOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.value)
	})

	return res
}

func (t *AvlTree[T, K]) ToPreOrderNodeSlice() []K {
	res := make([]K, 0, t.length)

	t.PreOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.value)
	})

	return res
}

func (t *AvlTree[T, K]) ToPostOrderNodeSlice() []K {
	res := make([]K, 0, t.length)

	t.PostOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.value)
	})

	return res
}
