package redblacktree

import (
	"math"
	"strings"
)

func (t *RedBlackTree[T, K]) String() string {
	var str strings.Builder
	str.WriteString("RedBlackTree\n")
	if !t.IsEmpty() {
		t.output(t.root, "", true, &str)
	}
	return str.String()
}

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
		cmpV := t.cmp(key, current.key)
		switch {
		case cmpV == 0:
			return current, true
		case cmpV < 0:
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
		cmpV := t.cmp(key, current.key)
		switch {
		case cmpV == 0:
			return current
		case cmpV < 0:
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
	case t.cmp(key, parent.key) < 0:
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

func (t *RedBlackTree[T, K]) Copy() Tree[T, K] {
	if t.root == nil {
		return &RedBlackTree[T, K]{
			root:   nil,
			cmp:    t.cmp,
			length: 0,
		}
	}

	newTree := &RedBlackTree[T, K]{
		root:   nil,
		cmp:    t.cmp,
		length: t.length,
	}

	newTree.root = &Node[T, K]{
		key:    t.root.key,
		value:  t.root.value,
		color:  t.root.color,
		parent: nilNode[T, K](),
	}

	depth := t.getRelativeMaxDepth()
	stack, newStack := make([]*Node[T, K], 0, depth), make([]*Node[T, K], 0, depth)
	stack = append(stack, t.root)
	newStack = append(newStack, newTree.root)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		newCurrent := newStack[len(newStack)-1]
		newStack = newStack[:len(newStack)-1]

		if current.left.NotNilNode() {
			newLeft := &Node[T, K]{
				key:    current.left.key,
				value:  current.left.value,
				color:  current.left.color,
				parent: newCurrent,
				left:   nilNode[T, K](),
				right:  nilNode[T, K](),
			}
			newCurrent.left = newLeft
			stack = append(stack, current.left)
			newStack = append(newStack, newLeft)
		}

		if current.right.NotNilNode() {
			newRight := &Node[T, K]{
				key:    current.right.key,
				value:  current.right.value,
				color:  current.right.color,
				parent: newCurrent,
				left:   nilNode[T, K](),
				right:  nilNode[T, K](),
			}
			newCurrent.right = newRight
			stack = append(stack, current.right)
			newStack = append(newStack, newRight)
		}
	}

	return newTree
}

func (t *RedBlackTree[T, K]) getRelativeMaxDepth() int {
	if t.Size() <= defaultMaxDepthLimit {
		return defaultMaxDepth
	}
	return int(math.Ceil(math.Log2(float64(t.Size()))*relativeToDepthMul) + 1)
}

func (t *RedBlackTree[T, K]) output(node *Node[T, K], prefix string, isTail bool, str *strings.Builder) {
	if node.right.NotNilNode() {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		t.output(node.right, newPrefix, false, str)
	}
	str.WriteString(prefix)
	if isTail {
		str.WriteString("└── ")
	} else {
		str.WriteString("┌── ")
	}
	str.WriteString(node.String())
	str.WriteRune('\n')
	if node.left.NotNilNode() {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		t.output(node.left, newPrefix, true, str)
	}
}
