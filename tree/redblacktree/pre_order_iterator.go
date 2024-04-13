package redblacktree

func (t *RedBlackTree[T, K]) PreOrderIterator() *PreOrderIterator[T, K] {
	return &PreOrderIterator[T, K]{t, nil, start}
}

func (i *PreOrderIterator[T, K]) Next() bool {
	switch i.position {
	case end:
		return false
	case start:
		if i.tree.length == 0 {
			return false
		}
		i.node = i.tree.Root()
		i.position = process
		return true
	default:
		if i.node.left.NotNilNode() {
			i.node = i.node.left
			return true
		}
		if i.node.right.NotNilNode() {
			i.node = i.node.right
			return true
		}
		for i.node.parent.NotNilNode() {
			parent := i.node.parent
			if i.node.isLeftNode() && parent.right.NotNilNode() {
				i.node = parent.right
				return true
			}
			i.node = parent
		}
		i.position = end
		return false
	}
}

func (i *PreOrderIterator[T, K]) Prev() bool {
	switch i.position {
	case start:
		return false
	case end:
		if i.tree.length == 0 {
			return false
		}
		i.node = i.tree.Right()
		i.position = process
		return true
	default:
		if i.node.parent.NotNilNode() {
			if i.node.isLeftNode() {
				i.node = i.node.parent
				return true
			} else {
				i.node = i.node.parent
				if i.node.left.NotNilNode() {
					i.node = i.node.left
					for i.node.right.NotNilNode() {
						i.node = i.node.right
					}
					for i.node.left.NotNilNode() {
						i.node = i.node.left
					}
				}
				return true

			}
		}
		i.position = start
		return false
	}
}

func (i *PreOrderIterator[T, K]) Value() K {
	return i.node.value
}

func (i *PreOrderIterator[T, K]) Key() T {
	return i.node.key
}

func (i *PreOrderIterator[T, K]) Node() *Node[T, K] {
	return i.node
}

func (i *PreOrderIterator[T, K]) Start() {
	i.node = nil
	i.position = start
}

func (i *PreOrderIterator[T, K]) End() {
	i.node = nil
	i.position = end
}

func (i *PreOrderIterator[T, K]) First() bool {
	i.Start()
	return i.Next()
}

func (i *PreOrderIterator[T, K]) Last() bool {
	i.End()
	return i.Prev()
}
