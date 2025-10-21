package datastructures

import "fmt"

type BinarySearchTree interface {
	Add(int)
	Search(int) bool
	Min() int
	Max() int
	Height() int
	PreOrder()
	InOrder()
	PosOrder()
	LevelOrder()
	Remove(int)
}

// TreeNode representa um nó na árvore.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BST representa a árvore de busca binária
type BST struct {
	root *TreeNode
	size int
}

func NewBST() *BST {
	return &BST{}
}

// Add inicia a inserção a partir da raiz
func (t *BST) Add(value int) {
	t.root = Add(t.root, value)
	t.size++
}

// professor usou uma interpretação diferente, deixando separado essas ações
// ou seja, em uma função, se o nó é nulo, cria um novo nó
// e em outra função, decide se vai para a esquerda ou para a direita
// eu juntei as duas ações em uma só função recursiva
// caso base 1: se o nó é nulo, encontramos a posição. Então é criado e retornado o novo nó
// caso base 2: se o nó não é nulo, decide se vai para a esquerda ou para a direita

// Add é a função auxiliar recursiva.
// todo: como essa função é interna, pode ser deixada com letra minúscula
func Add(node *TreeNode, value int) *TreeNode {
	// Caso base: se o nó é nulo, encontramos a posição. Então é criado e retornado o novo nó
	if node == nil {
		return &TreeNode{Value: value}
	}
	// Passo recursivo: decide se vai para a esquerda ou para a direita
	if value <= node.Value {
		node.Left = Add(node.Left, value)
	} else {
		node.Right = Add(node.Right, value)
	}
	return node
}

func (t *BST) Search(value int) bool {
	return Search(t.root, value)
}

func Search(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}

	if node.Value == value {
		return true
	} else if value < node.Value {
		return Search(node.Left, value)
	} else {
		return Search(node.Right, value)
	}
}

func (t *BST) Min() int {
	return Min(t.root)
}

func Min(node *TreeNode) int {
	if node == nil {
		return -1 // ou algum valor que indique que a árvore está vazia
	}
	for node.Left != nil {
		node = node.Left
	}
	return node.Value
}

func (t *BST) Max() int {
	return Max(t.root)
}
func Max(node *TreeNode) int {
	if node == nil {
		return -1 // ou algum valor que indique que a árvore está vazia
	}
	for node.Right != nil {
		node = node.Right
	}
	return node.Value
}
func (t *BST) Height() int {
	return Height(t.root)
}

func Height(node *TreeNode) int {
	if node == nil {
		return -1
	}
	leftHeight := Height(node.Left)
	rightHeight := Height(node.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func (t *BST) PreOrder() {
	PreOrder(t.root)
}

func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Value, " ")
	PreOrder(node.Left)
	PreOrder(node.Right)
}
func (t *BST) InOrder() {
	InOrder(t.root)
}

func InOrder(node *TreeNode) {
	if node == nil {
		return
	}
	InOrder(node.Left)
	fmt.Print(node.Value, " ")
	InOrder(node.Right)
}

func (t *BST) PosOrder() {
	PosOrder(t.root)
}

func PosOrder(node *TreeNode) {
	if node == nil {
		return
	}
	PosOrder(node.Left)
	PosOrder(node.Right)
	fmt.Print(node.Value, " ")
}

func (t *BST) LevelOrder() {
	LevelOrder(t.root)
}

func LevelOrder(node *TreeNode) {
	if node == nil {
		return
	}
	queue := []*TreeNode{node}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Print(current.Value, " ")
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}
func (t *BST) Remove(value int) {
	t.root = Remove(t.root, value)
}

func Remove(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = Remove(node.Left, value)
	} else if value > node.Value {
		node.Right = Remove(node.Right, value)
	} else {
		// Nó com apenas um filho ou nenhum filho
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		// Nó com dois filhos: obtém o sucessor inorder (menor na subárvore direita)
		minRight := Min(node.Right)
		node.Value = minRight
		node.Right = Remove(node.Right, minRight)
	}
	return node
}
