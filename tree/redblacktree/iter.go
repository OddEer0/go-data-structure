package redblacktree

func (t *RedBlackTree[T, K]) PreOrderFunc(callback func(*Node[T, K])) {
	if t.root == nil {
		return
	}
	var reduce func(*Node[T, K])
	reduce = func(node *Node[T, K]) {
		if node.NotNilNode() {
			callback(node)
			reduce(node.left)
			reduce(node.right)
		}
	}
	reduce(t.root)
}

func (t *RedBlackTree[T, K]) InOrderFunc(callback func(*Node[T, K])) {
	if t.root == nil {
		return
	}
	var reduce func(*Node[T, K])
	reduce = func(node *Node[T, K]) {
		if node.NotNilNode() {
			reduce(node.left)
			callback(node)
			reduce(node.right)
		}
	}
	reduce(t.root)
}

func (t *RedBlackTree[T, K]) PostOrderFunc(callback func(*Node[T, K])) {
	if t.root == nil {
		return
	}
	var reduce func(*Node[T, K])
	reduce = func(node *Node[T, K]) {
		if node.NotNilNode() {
			reduce(node.left)
			reduce(node.right)
			callback(node)
		}
	}
	reduce(t.root)
}
