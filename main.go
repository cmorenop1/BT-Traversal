package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) PreOrderTraversal() {
	//base case
	if node == nil {
		return
	}

	// Pre-order traversal (root, left, right)
	fmt.Println("node.Value:", node.Value)
	node.Left.PreOrderTraversal()
	node.Right.PreOrderTraversal()
}

func (node *Node) InOrderTraversal() {
	//base case
	if node == nil {
		return
	}

	// In-order traversal (left, root, right)
	node.Left.InOrderTraversal()
	fmt.Println("node.Value:", node.Value)
	node.Right.InOrderTraversal()
}

func (node *Node) PostOrderTraversal() {
	//base case
	if node == nil {
		return
	}

	// Post-order traversal (left, right, root)
	node.Left.PostOrderTraversal()
	node.Right.PostOrderTraversal()
	fmt.Println("node.Value:", node.Value)
}

func CreateNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func (node *Node) AddChildrenAt(targetValue int, left, right *Node) {
	//base case 1
	if node == nil {
		return
	}

	//base case 2
	if node.Value == targetValue {
		node.Left = left
		node.Right = right
		return
	}

	node.Left.AddChildrenAt(targetValue, left, right)
	node.Right.AddChildrenAt(targetValue, left, right)
}

func main() {
	tree := CreateNode(7)
	tree.AddChildrenAt(7, CreateNode(23), CreateNode(3))
	tree.AddChildrenAt(23, CreateNode(5), CreateNode(4))
	tree.AddChildrenAt(3, CreateNode(18), CreateNode(21))

	fmt.Println("Pre-order traversal:")
	tree.PreOrderTraversal()





	fmt.Println("In-order traversal:")
	tree.InOrderTraversal()
	fmt.Println("Post-order traversal:")
	tree.PostOrderTraversal()
}
