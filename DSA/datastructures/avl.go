package datastructures

import (
	"errors"
	"fmt"
)

// Interface para AVL Tree
type AVLTreeInterface interface {
	Add(int)
	Search(int) bool
	Min() (int, error)
	Max() (int, error)
	Height() int
	PreOrder()
	InOrder()
	PosOrder()
	LevelOrder()
	Remove(int) error
}

// AVLTree estrutura principal
type AVLTree struct {
	root *AVLNode
	size int
}

// AVLNode estrutura do nó da AVL
type AVLNode struct {
	Value  int
	Left   *AVLNode
	Right  *AVLNode
	Height int
	BF     int // Balance Factor
}

func NewAVL() *AVLTree {
	return &AVLTree{}
}

// Função auxiliar para criar nó
func createAVLNode(val int) *AVLNode {
	return &AVLNode{Value: val, Left: nil, Right: nil}
}

// === Auxiliares ===

// UpdateProperties atualiza altura e fator de balanço
func (node *AVLNode) UpdateProperties() {
	hLeft := 0
	hRight := 0

	if node.Left != nil {
		hLeft = node.Left.Height
	}
	if node.Right != nil {
		hRight = node.Right.Height
	}

	// Fator de Balanço = Dir - Esq
	node.BF = hRight - hLeft

	// Altura
	if node.Left == nil && node.Right == nil {
		node.Height = 0 // Folha tem altura 0
	} else {
		if hLeft > hRight {
			node.Height = hLeft + 1
		} else {
			node.Height = hRight + 1
		}
	}
}

// === Rotações ===

func (root *AVLNode) RotRight() *AVLNode {
	newRoot := root.Left
	root.Left = newRoot.Right
	newRoot.Right = root

	root.UpdateProperties()
	newRoot.UpdateProperties()

	return newRoot
}

func (root *AVLNode) RotLeft() *AVLNode {
	newRoot := root.Right
	root.Right = newRoot.Left
	newRoot.Left = root

	root.UpdateProperties()
	newRoot.UpdateProperties()

	return newRoot
}

// === Rebalanceamento ===

func (root *AVLNode) RebalanceLeftLeft() *AVLNode {
	return root.RotRight()
}

func (root *AVLNode) RebalanceLeftNeutral() *AVLNode {
	return root.RotRight()
}

func (root *AVLNode) RebalanceLeftRight() *AVLNode {
	root.Left = root.Left.RotLeft()
	return root.RotRight()
}

func (root *AVLNode) RebalanceRightRight() *AVLNode {
	return root.RotLeft()
}

func (root *AVLNode) RebalanceRightNeutral() *AVLNode {
	return root.RotLeft()
}

func (root *AVLNode) RebalanceRightLeft() *AVLNode {
	root.Right = root.Right.RotRight()
	return root.RotLeft()
}

// Rebalance decide qual rotação aplicar
func (root *AVLNode) Rebalance() *AVLNode {
	if root == nil {
		return nil
	}

	// Desbalanceado para a Esquerda
	if root.BF == -2 {
		if root.Left.BF == -1 { // Caso 1: Esq-Esq
			return root.RebalanceLeftLeft()
		} else if root.Left.BF == 0 { // Caso 1b: Esq-Neutro (após remoção)
			return root.RebalanceLeftNeutral()
		} else if root.Left.BF == 1 { // Caso 2: Esq-Dir
			return root.RebalanceLeftRight()
		}
	} else if root.BF == 2 { // Desbalanceado para a Direita
		if root.Right.BF == 1 { // Caso 3: Dir-Dir
			return root.RebalanceRightRight()
		} else if root.Right.BF == 0 { // Caso 3b: Dir-Neutro
			return root.RebalanceRightNeutral()
		} else if root.Right.BF == -1 { // Caso 4: Dir-Esq
			return root.RebalanceRightLeft()
		}
	}
	return root
}

// === Inserção (Add) ===

func (t *AVLTree) Add(val int) {
	if t.root == nil {
		t.root = createAVLNode(val)
	} else {
		t.root = t.root.Add(val)
	}
	t.size++
}

func (node *AVLNode) Add(val int) *AVLNode {
	if val <= node.Value {
		if node.Left == nil {
			node.Left = createAVLNode(val)
		} else {
			node.Left = node.Left.Add(val)
		}
	} else {
		if node.Right == nil {
			node.Right = createAVLNode(val)
		} else {
			node.Right = node.Right.Add(val)
		}
	}

	node.UpdateProperties()
	return node.Rebalance()
}

// === Busca (Search) ===

func (t *AVLTree) Search(val int) bool {
	if t.root == nil {
		return false
	}
	return t.root.Search(val)
}

func (node *AVLNode) Search(val int) bool {
	if val == node.Value {
		return true
	} else if val < node.Value {
		if node.Left == nil {
			return false
		}
		return node.Left.Search(val)
	} else {
		if node.Right == nil {
			return false
		}
		return node.Right.Search(val)
	}
}

// === Remoção ===

func (t *AVLTree) Remove(val int) error {
	if t.root == nil {
		return errors.New("impossivel remover de arvore vazia")
	} else {
		t.size--
		t.root = t.root.Remove(val)
		return nil
	}
}

func (node *AVLNode) Remove(val int) *AVLNode {
	if val < node.Value {
		if node.Left != nil {
			node.Left = node.Left.Remove(val)
		} else {
			return nil
		}
	} else if val > node.Value {
		if node.Right != nil {
			node.Right = node.Right.Remove(val)
		} else {
			return nil
		}
	} else {
		// Achamos o nó
		if node.Left == nil && node.Right == nil {
			return nil // Folha
		} else if node.Left != nil && node.Right == nil {
			return node.Left // 1 filho à esquerda
		} else if node.Left == nil && node.Right != nil {
			return node.Right // 1 filho à direita
		} else {
			// 2 filhos
			min, _ := node.Right.Min()
			node.Value = min
			node.Right = node.Right.Remove(min)
		}
	}

	node.UpdateProperties()
	return node.Rebalance()
}

// === Auxiliares de Busca (Min/Max) ===

func (node *AVLNode) Min() (int, error) {
	min := node
	for min.Left != nil {
		min = min.Left
	}
	return min.Value, nil
}

func (node *AVLNode) Max() (int, error) {
	max := node
	for max.Right != nil {
		max = max.Right
	}
	return max.Value, nil
}

// === Travessias ===

func (t *AVLTree) InOrder() {
	if t.root != nil {
		t.root.InOrder()
	}
	fmt.Println()
}

func (node *AVLNode) InOrder() {
	if node != nil {
		node.Left.InOrder()
		fmt.Printf("v:%d(h:%d bf:%d) ", node.Value, node.Height, node.BF)
		node.Right.InOrder()
	}
}

// LevelOrder para AVL
func (t *AVLTree) LevelOrder() {
	if t.root == nil {
		return
	}
	queue := []*AVLNode{t.root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Printf("v:%d(h:%d bf:%d) ", current.Value, current.Height, current.BF)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	fmt.Println()
}
