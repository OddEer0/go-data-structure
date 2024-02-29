package binarysearchtree

func (t *Tree[T, K]) ToSortedSlice() []K {
	res := make([]K, 0, t.length)

	t.InOrderFunc(func(node *Node[T, K]) {
		res = append(res, node.value)
	})

	return res
}
