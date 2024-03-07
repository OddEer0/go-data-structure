package main

import (
	"fmt"

	"github.com/OddEer0/go-data-structure/pkg/redblacktree"
)

func printRedBlackTree(root *redblacktree.Node[int, int], indent string, last bool) {
	if root == nil || root.NilNode() {
		return
	}

	fmt.Print(indent)
	if last {
		fmt.Print("R----")
		indent += "     "
	} else {
		fmt.Print("L----")
		indent += "|    "
	}

	color := "RED"
	if root.IsBlack() {
		color = "BLACK"
	}

	fmt.Printf("%d (%s)\n", root.Value(), color)

	printRedBlackTree(root.Left(), indent, false)
	printRedBlackTree(root.Right(), indent, true)
}

func main() {
	tree := redblacktree.New[int, int]()

	for i := 0; i < 700; i++ {
		tree.Insert(i+1, i+1)
	}

	// printRedBlackTree(tree.Root(), "", true)
	// fmt.Println("ROOT: ", tree.Root().Value())

	for i := 50; i < 100; i++ {
		tree.Remove(i + 1)
	}

	for i := 400; i < 550; i++ {
		tree.Remove(i + 1)
	}

	for i := 200; i < 270; i++ {
		tree.Remove(i + 1)
	}

	printRedBlackTree(tree.Root(), "", true)
	fmt.Println("ROOT: ", tree.Root().Value())
}
