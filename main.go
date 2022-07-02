package main

import (
	"BinaryTree/BT"
	"fmt"
)

func main() {
	tree := BT.New(50)

	tree.Insert(30)
	tree.Insert(60)
	tree.Insert(80)
	tree.Insert(25)
	tree.Insert(12)
	tree.Insert(45)
	tree.Insert(17)

	tree.PrintTree()

	s, ok := tree.Search(17)
	fmt.Printf("\nseach node: %d\n", s.Data)

	s, ok = tree.Search(39)
	if !ok {
		fmt.Println("node has 39 data do not exist")
	}

	fmt.Println("remove 30, 17, 60")
	tree.Remove(30)
	tree.Remove(17)
	tree.Remove(60)

	tree.PrintTree()
}
