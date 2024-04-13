package main

import (
	"fmt"

	"github.com/JFKongphop/DSA-golang/tree"
)

func main() {
	root  := tree.NewTreeNode(10)
	root.Insert(18)
	root.Insert(11)
	root.Insert(10)
	root.Insert(26)
	root.Insert(14)
	root.Insert(13)
	root.Insert(21)
	root.Insert(20)

	fmt.Println("PreOrder:")
	root.PreOrder()
	fmt.Println("\nInOrder:")
	root.InOrder()
	fmt.Println("\nPostOrder:")
	root.PostOrder()

	fmt.Println("\n\nAfter deleting 26")
	root.Delete(26)
	fmt.Println("PreOrder:")
	root.PreOrder()
	fmt.Println("\nInOrder:")
	root.InOrder()
	fmt.Println("\nPostOrder:")
	root.PostOrder()

	fmt.Println("\n\nAfter deleting 16")
	root.Delete(16)
	fmt.Println("PreOrder:")
	root.PreOrder()
	fmt.Println("\nInOrder:")
	root.InOrder()
	fmt.Println("\nPostOrder:")
	root.PostOrder()
}