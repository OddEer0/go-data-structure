package redblacktree

import "fmt"

func (n *Node[T, K]) String() string {
	return fmt.Sprintf("key: %v, value: %v", n.Key, n.Value)
}

func (n *Node[T, K]) Key() T {
	return n.key
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

func (n *Node[T, K]) Parent() *Node[T, K] {
	return n.parent
}

func (n *Node[T, K]) IsBlack() bool {
	return n.color == black
}

func (n *Node[T, K]) IsRed() bool {
	return n.color == red
}

func (n *Node[T, K]) NotNilNode() bool {
	return n != nilNode[T, K]()
}

func (n *Node[T, K]) NilNode() bool {
	return n == nilNode[T, K]()
}

func (n *Node[T, K]) isLeftNode() bool {
	return n == n.parent.left
}

func (n *Node[T, K]) isRightNode() bool {
	return n == n.parent.right
}

func (n *Node[T, K]) getChildrenCount() int {
	result := 0

	if n.left.NotNilNode() {
		result++
	}
	if n.right.NotNilNode() {
		result++
	}

	return result
}

func (n *Node[T, K]) getChildOrNil() *Node[T, K] {
	if n.left.NotNilNode() {
		return n.left
	}
	return n.right
}
