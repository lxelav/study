package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) Insert(data int) {
	bst.Root = insert(bst.Root, data)
}

func (bst *BST) Search(data int) *Node {
	return search(bst.Root, data)
}

func (bst *BST) Delete(data int) {
	bst.Root = delete(bst.Root, data)
}

func insert(node *Node, data int) *Node {
	if node == nil {
		return &Node{Data: data}
	}
	if data < node.Data {
		node.Left = insert(node.Left, data)
	} else if data > node.Data {
		node.Right = insert(node.Right, data)
	}
	return node
}

func search(node *Node, data int) *Node {
	if node == nil || node.Data == data {
		return node
	} else if data < node.Data {
		return search(node.Left, data)
	} else if data > node.Data {
		return search(node.Right, data)
	}
	return nil
}

func delete(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	if data < node.Data {
		node.Left = delete(node.Left, data)
	} else if data > node.Data {
		node.Right = delete(node.Right, data)
	} else {
		if node.Left == nil
	}
}

func main() {
	bst := &BST{}

	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, value := range values {
		bst.Insert(value)
	}

	fmt.Println(bst.Search(3).Data)
}
