package redblacktree

func (t *RedBlackTree[T, K]) Root() *Node[T, K] {
	return t.root
}

func (t *RedBlackTree[T, K]) Size() int {
	return t.length
}

func (t *RedBlackTree[T, K]) Clear() {
	t.root = nil
	t.length = 0
}

func (t *RedBlackTree[T, K]) IsEmpty() bool {
	return t.length == 0
}

func (t *RedBlackTree[T, K]) GetNode(key T) (*Node[T, K], bool) {
	current := t.root
	for current.NotNilNode() && current != nil {
		switch {
		case key == current.key:
			return current, true
		case t.cmp(key, current.key):
			current = current.left
		default:
			current = current.right
		}
	}

	return nil, false
}

func (t *RedBlackTree[T, K]) Get(key T) (K, bool) {
	node, ok := t.GetNode(key)
	if ok {
		return node.value, true
	}
	var val K
	return val, false
}

func (t *RedBlackTree[T, K]) Update(key T, value K) bool {
	node, ok := t.GetNode(key)
	if !ok {
		return false
	}
	node.value = value
	return true
}

func (t *RedBlackTree[T, K]) Insert(key T, value K) *Node[T, K] {
	current := t.root
	parent := nilNode[T, K]()
	for current.NotNilNode() && current != nil {
		parent = current
		switch {
		case key == current.key:
			return current
		case t.cmp(key, current.key):
			current = current.left
		default:
			current = current.right
		}
	}

	newNode := NewNode(key, value)
	newNode.parent = parent

	switch {
	case !parent.NotNilNode():
		t.root = newNode
	case t.cmp(key, parent.key):
		parent.left = newNode
	default:
		parent.right = newNode
	}

	t.balanceInsert(newNode)
	t.length++
	return nil
}

func (t *RedBlackTree[T, K]) InsertOrUpdate(key T, value K) {
	node := t.Insert(key, value)
	if node != nil {
		node.value = value
	}
}

func (t *RedBlackTree[T, K]) InsertMany(entries ...*Entry[T, K]) {
	for _, entry := range entries {
		t.Insert(entry.Key, entry.Value)
	}
}

func (t *RedBlackTree[T, K]) InsertOrUpdateMany(entries ...*Entry[T, K]) {
	for _, entry := range entries {
		node := t.Insert(entry.Key, entry.Value)
		if node != nil {
			node.value = entry.Value
		}
	}
}

func (t *RedBlackTree[T, K]) Remove(key T) bool {
	deletedNode, ok := t.GetNode(key)
	if !ok {
		return false
	}

	var child *Node[T, K]
	removedColor := deletedNode.color
	if deletedNode.getChildrenCount() < 2 { // Если у удаляемого узла 1 или 0 дочерних элементов тогда производим удаление
		child = deletedNode.getChildOrNil()
		t.swapNode(deletedNode, child)
	} else { // Если у него 2 дочерних элемента, тогда находим минимальный слева и меняем удаляемый на него. После удаляем минимальный
		swappedRight := t.getRightSwappedNode(deletedNode.right)
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
	return true
}

func (t *RedBlackTree[T, K]) RemoveMany(keys ...T) {
	for _, key := range keys {
		t.Remove(key)
	}
}

func (t *RedBlackTree[T, K]) Left() *Node[T, K] {
	if t.length == 0 {
		return nil
	}
	current := t.root
	for current.left.NotNilNode() {
		current = current.left
	}
	return current
}

func (t *RedBlackTree[T, K]) Right() *Node[T, K] {
	if t.length == 0 {
		return nil
	}
	current := t.root
	for current.right.NotNilNode() {
		current = current.right
	}
	return current
}
