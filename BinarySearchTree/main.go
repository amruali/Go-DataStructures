package main

import "fmt"

type Node struct {
	key    int
	left   *Node
	right  *Node
	parent *Node
}

func main() {

	BSTree := NewTree(10)
	BSTree.Insert(3)
	BSTree.Insert(6)
	BSTree.Insert(1)
	BSTree.Insert(0)
	BSTree.Insert(2)
	BSTree.Insert(4)
	BSTree.Insert(5)
	BSTree.Insert(8)
	BSTree.Insert(7)
	BSTree.Insert(9)
	BSTree.Insert(13)
	BSTree.Insert(12)
	BSTree.Insert(15)
	BSTree.Insert(14)
	BSTree.Insert(16)
	BSTree.Insert(20)
	BSTree.Insert(18)
	BSTree.Insert(22)
	fmt.Println(BSTree)

	/*
	   									   10
	   								        /      \
	   								       3       13
	                                                            	     /   \    /  \
	                          				            1     6  12   15
	                                				   / \   / \     /  \
	                               					  0   2 4   8   14  16
	   								       \   / \	       \
	   								     	5 7   9        20
	   										      /  \
	                        				                             18  22

	*/

	BSTree.InOrderTraversalBSTree()
	fmt.Println("\n------------------")

	BSTree.PreOrderTraversalBSTree()
	fmt.Println("\n------------------")

	BSTree.PostOrderTraversalBSTree()
	fmt.Println("\n------------------")

	fmt.Println("Minimum is : ", BSTree.Min())
	fmt.Print("Maximum is : ", BSTree.Max())

	fmt.Println("\n------------------")
	if _, Ok := BSTree.find(99); !Ok {
		fmt.Print("Not Found")
	} else {
		fmt.Print("Found")
	}

	fmt.Println("\n------------------")

	fmt.Print(BSTree.Successor(20))
	fmt.Println("\n------------------")
	fmt.Print(BSTree.Predecessor(1))

	fmt.Println("\n------------------")

	fmt.Print(BSTree.Delete(10))

	fmt.Println("\n------------------")

	BSTree.InOrderTraversalBSTree()
}

func NewTree(key int) *Node {
	return &Node{
		key:    key,
		parent: nil,
	}
}

func (n *Node) Insert(key int) {
	if key <= n.key {
		if n.left == nil {
			n.left = &Node{key: key, parent: n}
		} else {
			n.left.Insert(key)
		}
	} else if key > n.key {
		if n.right == nil {
			n.right = &Node{key: key, parent: n}
		} else {
			n.right.Insert(key)
		}
	}
}

func (n *Node) find(key int) (*Node, bool) {
	if n == nil {
		return nil, false
	}
	if n.key == key {
		return n, true
	}
	if key < n.key {
		return n.left.find(key)
	} else {
		return n.right.find(key)
	}
}

// Root - Left - Right
func (n *Node) PreOrderTraversalBSTree() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.key)
	n.left.PreOrderTraversalBSTree()
	n.right.PreOrderTraversalBSTree()
}

// Left - Root - Right
func (n *Node) InOrderTraversalBSTree() {
	if n == nil {
		return
	}
	if n.left != nil {
		n.left.InOrderTraversalBSTree()
	}
	fmt.Print(n.key, " ")
	n.right.InOrderTraversalBSTree()
}

// Left - Right - Root
func (n *Node) PostOrderTraversalBSTree() {
	if n == nil {
		return
	}
	n.left.PostOrderTraversalBSTree()
	n.right.PostOrderTraversalBSTree()
	fmt.Print(n.key, " ")
}

// B.S.Tree Minimum key
func (n *Node) Min() int {
	if n.left != nil {
		return n.left.Min()
	}
	return n.key
}

// B.S.Tree Maximum Key
func (n *Node) Max() int {
	if n.right != nil {
		return n.right.Max()
	}
	return n.key
}

// smallest node with key value larger than node with key parameter
func (n *Node) Successor(key int) int {
	KeyNode, Ok := n.find(key)
	// Node is not found or node key is the maximum in the tree
	if !Ok || KeyNode.key == n.Max() {
		return -1
	}
	if KeyNode.right != nil {
		return KeyNode.right.Min()
	}
	p := KeyNode.parent
	for p != nil && KeyNode == p.right {
		KeyNode = p
		p = p.parent
	}
	return p.key
}

// the largest node with key value smaller than node with key parameter
func (n *Node) Predecessor(key int) int {
	KeyNode, Ok := n.find(key)
	// This Means node is not found or node key is the minimum in the tree
	if !Ok || KeyNode.key == n.Min() {
		return -1
	}
	if KeyNode.left != nil {
		return KeyNode.left.Max()
	}

	p := KeyNode.parent
	for p != nil && KeyNode == p.left {
		KeyNode = p
		p = p.parent
	}
	return p.key
}

func (n *Node) Delete(key int) error {
	n, Ok := n.find(key)
	if !Ok {
		return fmt.Errorf("node is not found")
	}
	p := n.parent
	// Case 1 ( No left or right children nodes )
	if n.left == nil && n.right == nil {
		// left node
		if p.left != nil && p.left.key == key {
			p.left = nil
		} else {
			p.right = nil
		}
		return nil
	}
	if n.left == nil {
		n.right.parent = p
		if p.right == n {
			p.right = n.right
			return nil
		}
		p.left = n.right
		return nil
	}
	if n.right == nil {
		n.left.parent = p
		if p.left == n {
			p.left = n.left
			return nil
		}
		p.right = n.left
		return nil
	}
	// Case 3 (2 Nodes with or without subtrees)
	predecessorKey := n.Predecessor(n.key) // We can also use predecessorKey := n.left.Max()
	predecessorNode, _ := n.find(predecessorKey)
	n.key = predecessorKey
	return predecessorNode.Delete(predecessorKey)
}
