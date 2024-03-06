package rbtree

func (t *Tree[T, K]) GetRoot() *Node[T, K] {
	return t.root
}

func (t *Tree[T, K]) GetSize() int {
	return t.length
}

func (t *Tree[T, K]) ChangeCmpFunc(fn func(a, b T) bool) {
	t.cmp = fn
}

func (t *Tree[T, K]) GetNode(key T) (*Node[T, K], bool) {
	current := t.root
	for current.NotNilNode() && current != nil {
		switch {
		case key == current.key:
			return current, true
		case t.cmp(key, current.key):
			current = current.right
		default:
			current = current.left
		}
	}

	return nil, false
}

func (t *Tree[T, K]) Update(key T, value K) bool {
	node, ok := t.GetNode(key)
	if !ok {
		return false
	}
	node.value = value
	return true
}

func (t *Tree[T, K]) Insert(key T, value K) *Node[T, K] {
	current := t.root
	parent := nilNode[T, K]()
	for current.NotNilNode() && current != nil {
		parent = current
		switch {
		case key == current.key:
			return current
		case t.cmp(key, current.key):
			current = current.right
		default:
			current = current.left
		}
	}

	newNode := NewNode(key, value)
	newNode.parent = parent

	switch {
	case !parent.NotNilNode():
		t.root = newNode
	case t.cmp(key, parent.key):
		parent.right = newNode
	default:
		parent.left = newNode
	}

	t.balanceInsert(newNode)
	t.length++
	return nil
}

func (t *Tree[T, K]) Remove(key T) {
	deletedNode, ok := t.GetNode(key)
	if !ok {
		return
	}

	var child *Node[T, K]
	removedColor := deletedNode.color
	if deletedNode.getChildrenCount() < 2 { // Если у удаляемого узла 1 или 0 дочерних элементов тогда производим удаление
		child = deletedNode.getChildOrNil()
		t.swapNode(deletedNode, child)
	} else { // Если у него 2 дочерних элемента, тогда находим минимальный слева и меняем удаляемый на него. После удаляем минимальный
		swappedRight := deletedNode.right.getMin()
		deletedNode.key = swappedRight.key
		deletedNode.value = swappedRight.value
		removedColor = swappedRight.color
		child = swappedRight.getChildOrNil()
		t.swapNode(swappedRight, child)
	}
	// Если у удаляемого узла красный цвет то балансировать нам нечего. Если черный, тогда нужно балансировать
	if removedColor == black {
		t.balanceRemove(child)
	}
	t.length--
}
