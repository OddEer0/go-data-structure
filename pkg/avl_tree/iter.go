package avltree

func (t *Tree[T, K]) PreOrderFunc(callback func(*Node[T, K])) {
	var reduce func(*Node[T, K])
	reduce = func(node *Node[T, K]) {
		if node != nil {
			callback(node)
			reduce(node.left)
			reduce(node.right)
		}
	}
	reduce(t.root)
}

func (t *Tree[T, K]) InOrderFunc(callback func(*Node[T, K])) {
	var reduce func(*Node[T, K])
	reduce = func(node *Node[T, K]) {
		if node != nil {
			reduce(node.left)
			callback(node)
			reduce(node.right)
		}
	}
	reduce(t.root)
}

func (t *Tree[T, K]) PostOrderFunc(callback func(*Node[T, K])) {
	var reduce func(*Node[T, K])
	reduce = func(node *Node[T, K]) {
		if node != nil {
			reduce(node.left)
			reduce(node.right)
			callback(node)
		}
	}
	reduce(t.root)
}
