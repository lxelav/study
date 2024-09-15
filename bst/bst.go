package bst

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

type Bst struct {
	Root *Node
}

func (bst *Bst) Insert(value int) {
	bst.Root = insert(bst.Root, value)
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}
	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	}

	return node
}

func (bst *Bst) Search(value int) *Node {
	return search(bst.Root, value)
}

func search(node *Node, value int) *Node {
	var result *Node
	if node == nil {
		fmt.Println("Узла нет")
		return nil
	}
	if node.Value == value {
		result = node
		return result
	}
	if value < node.Value {
		result = search(node.Left, value)
	} else if value > node.Value {
		result = search(node.Right, value)
	}

	return result
}

func (bst *Bst) Delete(value int) {
	bst.Root = deleteNode(bst.Root, value)
}

func deleteNode(node *Node, value int) *Node {
	// Базовый случай
	if node == nil {
		return node
	}

	// Если ключ, который нужно удалить, меньше ключа корня,
	// тогда он находится в левом поддереве
	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
		return node
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
		return node
	}

	// Если ключ совпадает с ключом корня, то это узел, который нужно удалить
	// Узел с одним ребенком или без детей
	if node.Left == nil {
		temp := node.Right
		node = nil
		return temp
	} else if node.Right == nil {
		temp := node.Left
		node = nil
		return temp
	}

	// Узел с двумя детьми: Получаем inorder-наследника (наименьший
	// в правом поддереве)
	succParent := node
	succ := node.Right
	for succ.Left != nil {
		succParent = succ
		succ = succ.Left
	}

	node.Value = succ.Value

	if succParent.Left == succ {
		succParent.Left = succ.Right
	} else {
		succParent.Right = succ.Right
	}

	succ = nil
	return node
}

func Depth(node *Node) {
	if node == nil {
		return
	}

	Depth(node.Left)
	fmt.Println(node.Value)
	Depth(node.Right)
}

func width(node *Node) {
	if node == nil {
		fmt.Println("Error: root bst is nil")
		return
	}

	p := []*Node{node}
	for len(p) != 0 {
		current := p[0]
		p = p[1:]

		fmt.Println(current)

		if current.Left != nil {
			p = append(p, current.Left)
		}
		if current.Right != nil {
			p = append(p, current.Right)
		}
	}
}
