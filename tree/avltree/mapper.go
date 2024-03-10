package avltree

func (t *Tree[T, K]) ToSortedSlice() []K {
	res := make([]K, 0, t.length)

	t.InOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.value)
	})

	return res
}

func (t *Tree[T, K]) ToPreOrderNodeSlice() []K {
	res := make([]K, 0, t.length)

	t.PreOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.value)
	})

	return res
}

func (t *Tree[T, K]) ToPostOrderNodeSlice() []K {
	res := make([]K, 0, t.length)

	t.PostOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.value)
	})

	return res
}
