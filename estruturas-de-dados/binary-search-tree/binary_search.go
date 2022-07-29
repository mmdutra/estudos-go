package binarysearchtree

import (
	"strconv"
)

// Node represents the components of a binary search tree
type Node struct {
	Key   int
	left  *Node
	right *Node
}

// Insert will add a node to the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.right == nil {
			n.right = &Node{Key: k}
		} else {
			n.right.Insert(k)
		}
	} else if n.Key > k {
		if n.left == nil {
			n.right = &Node{Key: k}
		} else {
			n.left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.right.Search(k)
	} else if n.Key > k {
		return n.left.Search(k)
	}

	return true
}

func Print(n *Node) string {
	var str string

	if n == nil {
		return str
	}

	str += Print(n.left)
	str += (strconv.Itoa(n.Key) + " ")
	str += Print(n.right)

	return str
}
