package redblacktree

func (t *RedBlackTree[T, K]) Iterator() *Iterator[T, K] {
	return &Iterator[T, K]{t, nil, start}
}

func (i *Iterator[T, K]) Next() bool {
	switch i.position {
	case end:
		return false
	case start:
		if i.tree.length == 0 {
			return false
		}
		i.node = i.tree.Left()
		i.position = proccess
		return true
	default:
		if i.node.right.NotNilNode() {
			i.node = i.node.right
			for i.node.left.NotNilNode() {
				i.node = i.node.left
			}
			return true
		}
		for i.node.parent.NotNilNode() {
			node := i.node
			i.node = i.node.Parent()
			if node == i.node.left {
				return true
			}
		}

		i.node = nil
		i.position = end
		return false
	}
}

func (i *Iterator[T, K]) Prev() bool {
	switch i.position {
	case start:
		return false
	case end:
		if i.tree.length == 0 {
			return false
		}
		i.node = i.tree.Right()
		i.position = proccess
		return true
	default:
		if i.node.left.NotNilNode() {
			i.node = i.node.left
			for i.node.right.NotNilNode() {
				i.node = i.node.right
			}
			return true
		}
		for i.node.parent.NotNilNode() {
			node := i.node
			i.node = i.node.Parent()
			if node == i.node.right {
				return true
			}
		}

		i.node = nil
		i.position = start
		return false
	}
}

func (i *Iterator[T, K]) Value() K {
	return i.node.value
}

func (i *Iterator[T, K]) Key() T {
	return i.node.key
}

func (i *Iterator[T, K]) Node() *Node[T, K] {
	return i.node
}

func (i *Iterator[T, K]) Start() {
	i.node = nil
	i.position = start
}

func (i *Iterator[T, K]) End() {
	i.node = nil
	i.position = end
}

func (i *Iterator[T, K]) First() bool {
	i.Start()
	return i.Next()
}

func (i *Iterator[T, K]) Last() bool {
	i.End()
	return i.Prev()
}
