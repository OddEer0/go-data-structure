package avltree

import "math"

func (n *Node[T, K]) getHeight(node *Node[T, K]) int {
	if node == nil {
		return -1
	}
	return node.height
}

func (n *Node[T, K]) updateHeight() {
	n.height = int(math.Max(float64(n.getHeight(n.left)), float64(n.getHeight(n.right)))) + 1
}

func (n *Node[T, K]) getBalance(node *Node[T, K]) int {
	if node == nil {
		return 0
	}
	return node.getHeight(node.right) - node.getHeight(node.left)
}

func (n *Node[T, K]) swap(a, b *Node[T, K]) {
	a.key, b.key = b.key, a.key
	a.value, b.value = b.value, a.value
}

func (n *Node[T, K]) rightRotate() {
	n.swap(n, n.left)
	tmp := n.right
	n.right = n.left
	n.left = n.right.left
	n.right.left = n.right.right
	n.right.right = tmp
	n.right.updateHeight()
	n.updateHeight()
}

func (n *Node[T, K]) leftRotate() {
	n.swap(n, n.right)
	tmp := n.left
	n.left = n.right
	n.right = n.left.right
	n.left.right = n.left.left
	n.left.left = tmp
	n.left.updateHeight()
	n.updateHeight()
}

func (n *Node[T, K]) balance() {
	balance := n.getBalance(n)
	switch balance {
	case -2:
		if n.getBalance(n.left) == 1 {
			n.left.leftRotate()
		}
		n.rightRotate()
	case 2:
		if n.getBalance(n.left) == -1 {
			n.right.rightRotate()
		}
		n.leftRotate()
	}
}
