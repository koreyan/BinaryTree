package BT

import "fmt"

type Element int

type Node struct {
	Data  Element
	Left  *Node
	Right *Node
}

type binaryTree struct {
	head *Node
}

func New(rootData Element) *binaryTree {
	root := &Node{Data: rootData, Left: nil, Right: nil}
	head := &Node{Left: root}
	return &binaryTree{head: head} //head is parent of root
}

func isThere(node *Node) bool {
	if node != nil {
		return true
	} else {
		return false
	}
}

//                                tree start
// head- Node pointing root -- root -  left node
//          \ nil                   \  right node

// tree.head.Left points root instance
func (tree *binaryTree) Search(data Element) (*Node, bool) {
	return search(tree.head.Left, data)
}

func (tree *binaryTree) Insert(data Element) bool {
	return insert(tree.head.Left, data)
}

func (tree *binaryTree) Remove(data Element) bool {

	return remove(tree.head.Left, tree.head, data)

}
func (tree *binaryTree) PrintTree() {
	print(tree.head.Left)
}

// For Search data in Binary Tree. if this function can't find data in tree, it returns nil and false otherwise adrress of node and true
func search(node *Node, target Element) (*Node, bool) {
	if node != nil {
		if target < node.Data {
			return search(node.Left, target)
		} else if target > node.Data {
			return search(node.Right, target)
		} else {
			return node, true
		}
	}

	return nil, false
}

// For Insert data in Binary Tree. if this function can't insert data in tree, it returns false otherwise true
func insert(node *Node, data Element) bool {

	if isThere(node) == false {
		fmt.Println("There is no root")
		return false
	}
	if data > node.Data {
		if isThere(node.Right) == false {
			node.Right = &Node{Data: data, Left: nil, Right: nil}
			return true
		} else {
			return insert(node.Right, data)
		}
	} else if data < node.Data {
		if isThere(node.Left) == false {
			node.Left = &Node{Data: data, Left: nil, Right: nil}
			return true
		} else {
			return insert(node.Left, data)
		}
	} else {
		return true
	}
}

// For Remove data in Binary Tree if this function can't remove data in tree, it returns false otherwise true
func remove(currentNode *Node, parentNode *Node, target Element) bool {
	if target < currentNode.Data {
		if isThere(currentNode.Left) == false {
			fmt.Println("Can not remove")
			return false
		}
		return remove(currentNode.Left, currentNode, target)

	} else if target > currentNode.Data {
		if isThere(currentNode.Right) == false {
			fmt.Println("Can not remove")
			return false
		}
		return remove(currentNode.Right, currentNode, target)

	} else {
		if currentNode.Left == nil && currentNode.Right == nil {
			if parentNode.Left == currentNode {
				parentNode.Left = nil
			} else {
				parentNode.Right = nil
			}
			return true
		} else if currentNode.Left == nil {
			if parentNode.Left == currentNode {
				parentNode.Left = currentNode.Right
			} else {
				parentNode.Right = currentNode.Right
			}
			return true
		} else if currentNode.Right == nil {
			if parentNode.Left == currentNode {
				parentNode.Left = currentNode.Left
			} else {
				parentNode.Right = currentNode.Left
			}
			return true
		} else { // if node to remove has two nodes
			smallestNode := currentNode.Right
			for smallestNode.Left != nil { //Find node has smallest data
				smallestNode = smallestNode.Left
			}

			tmpData := smallestNode.Data                       // backup smallest data
			remove(currentNode, parentNode, smallestNode.Data) // remove smallest node
			currentNode.Data = tmpData                         // copying smallest data to Data of currentNode
			return true
		}

	}
}

func print(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Data, " ")
	print(node.Left)
	print(node.Right)
}
