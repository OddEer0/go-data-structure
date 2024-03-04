package rbtree

func (t *Tree[T, K]) GetRoot() *Node[T, K] {
	return t.root
}

func (t *Tree[T, K]) GetSize() int {
	return t.length
}

func (n *Node[T, K]) Value() K {
	return n.value
}

func (n *Node[T, K]) Right() *Node[T, K] {
	return n.right
}

func (n *Node[T, K]) Left() *Node[T, K] {
	return n.left
}

func (n *Node[T, K]) NotNilNode() bool {
	return n != nilNode[T, K]()
}

func (t *Tree[T, K]) ChangeCmpFunc(fn func(a, b T) bool) {
	t.cmp = fn
}

func (t *Tree[T, K]) Insert(key T, value K) *Node[T, K] {
	current := t.root
	parent := nilNode[T, K]()
	for current.NotNilNode() && current != nil {
		parent = current
		if key == current.key {
			return current
		}
		if t.cmp(key, current.key) {
			current = current.right
		} else {
			current = current.left
		}
	}

	newNode := NewNode(key, value)
	newNode.parent = parent
	if !parent.NotNilNode() {
		t.root = newNode
	} else if t.cmp(key, parent.key) {
		parent.right = newNode
	} else {
		parent.left = newNode
	}

	t.balance(newNode)
	t.length++
	return nil
}
