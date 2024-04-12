package main

import (
	"fmt"

	"github.com/JFKongphop/DSA-golang/tree"
)

func main() {
	t := tree.NewTreeNode(10)
	t.Insert(5)
	t.Insert(4)
	t.Insert(3)

	fmt.Println(t.Min().Data)
}