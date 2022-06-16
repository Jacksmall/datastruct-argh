package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) insert(k int) {
	if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.insert(k)
		}
	} else if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.insert(k)
		}
	}
}

func (n *Node) search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.search(k)
	} else if n.Key > k {
		return n.Left.search(k)
	}

	return true
}

func main() {
	fmt.Println("Binary Search Tree")

	bst := &Node{Key: 100}
	bst.insert(50)
	bst.insert(200)
	bst.insert(88)
	bst.insert(90)
	bst.insert(108)
	fmt.Println(bst)
	fmt.Println(bst.search(80))
}
