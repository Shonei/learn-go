package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	key   int
}

func main() {
	/*
				    27
		          /    \
		     	 /      \
				14	    35
			   /  \    /  \
			  10  19  31  42
	*/
	tree := Node{left: &Node{left: &Node{left: nil, right: nil, key: 19},
		right: &Node{left: nil, right: nil, key: 10}, key: 14},
		right: &Node{left: &Node{left: nil, right: nil, key: 31},
			right: &Node{left: nil, right: nil, key: 42}, key: 35},
		key: 27}

	fmt.Println(loop(tree))
}

func loop(node Node) bool {

	if node.key > node.right.key && node.key >= node.left.key {
		return false
	} else if node.left.key > node.right.key {
		return false
	}

	if node.left.left != nil {
		if v := loop(*node.left); !v {
			return false
		}
	}
	if node.right.right != nil {
		if v := loop(*node.right); !v {
			return false
		}
	}

	return true
}
