package avltree

func (t *AvlTree[T, K]) Root() *Node[T, K] {
	return t.root
}

func (t *AvlTree[T, K]) Size() int {
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

func (n *Node[T, K]) GetMax() *Node[T, K] {
	switch {
	case n == nil:
		return nil
	case n.right == nil:
		return n
	}
	return n.right.GetMax()
}

func (t *AvlTree[T, K]) Insert(key T, value K) *Node[T, K] {
	if t.root == nil {
		t.root = &Node[T, K]{key, value, nil, nil, 0}
		t.length++
		return nil
	}

	var reduceFn func(*Node[T, K]) *Node[T, K]
	reduceFn = func(node *Node[T, K]) *Node[T, K] {
		var res *Node[T, K] = nil
		cmpValue := t.cmp(key, node.key)
		switch {
		case cmpValue == 0:
			res = node
		case cmpValue < 0:
			if node.left == nil {
				node.left = &Node[T, K]{key, value, nil, nil, 0}
			} else {
				res = reduceFn(node.left)
			}
		default:
			if node.right == nil {
				node.right = &Node[T, K]{key, value, nil, nil, 0}
			} else {
				res = reduceFn(node.right)
			}
		}
		node.updateHeight()
		node.balance()
		return res
	}

	node := reduceFn(t.root)
	if node == nil {
		t.length++
	}
	return node
}

func (t *AvlTree[T, K]) GetNode(key T) (*Node[T, K], bool) {
	var reduceFn func(*Node[T, K]) *Node[T, K]
	ok := false

	reduceFn = func(node *Node[T, K]) *Node[T, K] {
		if node == nil {
			return nil
		}
		cmpValue := t.cmp(key, node.key)
		switch {
		case cmpValue == 0:
			ok = true
			return node
		case cmpValue < 0:
			return reduceFn(node.left)
		default:
			return reduceFn(node.right)
		}
	}

	return reduceFn(t.root), ok
}

func (t *AvlTree[T, K]) Update(key T, value K) bool {
	node, ok := t.GetNode(key)
	if !ok {
		return false
	}
	node.value = value
	return true
}

func (t *AvlTree[T, K]) Remove(key T) {
	if t.root == nil {
		return
	}

	var reduceFn func(*Node[T, K], T) *Node[T, K]
	reduceFn = func(node *Node[T, K], k T) *Node[T, K] {
		if node == nil {
			return nil
		}
		cmpValue := t.cmp(k, node.key)
		switch {
		case cmpValue == 0:
			if node.left == nil {
				node = node.right
				t.length--
			} else if node.right == nil {
				node = node.left
				t.length--
			} else {
				maxInLeft := node.left.GetMax()
				node.key = maxInLeft.key
				node.value = maxInLeft.value
				node.left = reduceFn(node.left, maxInLeft.key)
			}
		case cmpValue < 0:
			node.left = reduceFn(node.left, k)
		default:
			node.right = reduceFn(node.right, k)
		}
		if node != nil {
			node.updateHeight()
			node.balance()
		}
		return node
	}
	t.root = reduceFn(t.root, key)
}
