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

func (n *Node[T, K]) GetMax() *Node[T, K] {
	switch {
	case n == nil:
		return nil
	case n.right == nil:
		return n
	}
	return n.right.GetMax()
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
		switch {
		case key == node.key:
			return false
		case t.cmp(key, node.key):
			if node.right == nil {
				node.right = &Node[T, K]{key, value, nil, nil}
				return true
			}
			return reduceFn(node.right)
		default:
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

func (t *Tree[T, K]) Remove(key T) {
	if t.root == nil {
		return
	}

	var reduceFn func(*Node[T, K], T) *Node[T, K]
	reduceFn = func(node *Node[T, K], k T) *Node[T, K] {
		switch {
		case node == nil:
			return nil
		case k == node.key:
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
		case t.cmp(k, node.key):
			node.right = reduceFn(node.right, k)
		default:
			node.left = reduceFn(node.left, k)
		}
		return node
	}
	t.root = reduceFn(t.root, key)
}

func (t *Tree[T, K]) GetNodeByKey(key T) (*Node[T, K], bool) {
	var reduceFn func(*Node[T, K]) *Node[T, K]
	ok := false

	reduceFn = func(node *Node[T, K]) *Node[T, K] {
		switch {
		case node == nil:
			return nil
		case node.key == key:
			ok = true
			return node
		case t.cmp(key, node.key):
			return reduceFn(node.right)
		default:
			return reduceFn(node.left)
		}
	}

	return reduceFn(t.root), ok
}

func (t *Tree[T, K]) Update(key T, value K) bool {
	node, ok := t.GetNodeByKey(key)
	if !ok {
		return false
	}
	node.value = value
	return true
}
