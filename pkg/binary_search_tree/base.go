package binarysearchtree

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

func (t *Tree[T, K]) ChangeCmpFunc(fn func(a, b T) bool) {
	t.cmp = fn
}

func (t *Tree[T, K]) Insert(key T, value K) bool {
	if t.root == nil {
		t.root = &Node[T, K]{key, value, nil, nil}
		t.length++
		return true
	}

	var reduceFn func(*Node[T, K]) bool
	reduceFn = func(node *Node[T, K]) bool {
		if key == node.key {
			return false
		}

		if t.cmp(key, node.key) {
			if node.right == nil {
				node.right = &Node[T, K]{key, value, nil, nil}
				return true
			}
			return reduceFn(node.right)
		} else {
			if node.left == nil {
				node.left = &Node[T, K]{key, value, nil, nil}
				return true
			}
			return reduceFn(node.left)
		}
	}

	ok := reduceFn(t.root)
	if ok {
		t.length++
	}
	return ok
}
