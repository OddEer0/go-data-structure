package rbtree

func (n *Node[T, K]) Value() K {
	return n.value
}

func (n *Node[T, K]) Right() *Node[T, K] {
	return n.right
}

func (n *Node[T, K]) Left() *Node[T, K] {
	return n.left
}

func (n *Node[T, K]) Parent() *Node[T, K] {
	return n.parent
}

func (n *Node[T, K]) IsBlack() bool {
	return n.color == black
}

func (n *Node[T, K]) IsRed() bool {
	return n.color == red
}

func (n *Node[T, K]) NotNilNode() bool {
	return n != nilNode[T, K]()
}

func (n *Node[T, K]) NilNode() bool {
	return n == nilNode[T, K]()
}

func (n *Node[T, K]) isLeftNode() bool {
	return n == n.parent.left
}

func (n *Node[T, K]) isRightNode() bool {
	return n == n.parent.right
}

func (n *Node[T, K]) getSwappedLeft() *Node[T, K] {
	switch {
	case n.NilNode():
		return n
	case n.right.NilNode():
		return n
	}
	return n.right.getSwappedLeft()
}

func (n *Node[T, K]) getChildrenCount() int {
	result := 0

	if n.left.NotNilNode() {
		result++
	}
	if n.right.NotNilNode() {
		result++
	}

	return result
}

func (n *Node[T, K]) getChildOrNil() *Node[T, K] {
	if n.left.NotNilNode() {
		return n.left
	}
	return n.right
}

func (n *Node[T, K]) getMin() *Node[T, K] {
	current := n
	for current.NotNilNode() {
		if n.left.NilNode() {
			return n
		}
		current = current.left
	}
	return n
}
