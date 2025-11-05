package datastructures

import (
	"fmt"
	"math"
)

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
	IsBst() bool
	Size() int
	SetRoot(node *TreeNode)
	SetSize(size int)
	ConvertToBalancedBst([]int) *BST
	CountEven() int
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

// Adicionados para permitir que o main.go construa uma BST
// a partir do nó raiz retornado pela função única.
func (t *BST) SetRoot(node *TreeNode) {
	t.root = node
}
func (t *BST) SetSize(size int) {
	t.size = size
}

// === Adição de Elemento ===

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

// === Procura de Elemento ===

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

// === Valores Mínimo/Máximo de Elemento ===

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

// === Altura Árvore ===

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

/* === Diferenças de Implementação ===
Professor navegou fazendo uma verificação no nó da esquerda
e no nó da direita em cada iteração da função:
func PreOrder(node *TreeNode) {
	fmrt.Print(node.Value, " ")
	if node.Left != nil {
		PreOrder(node.Left)
	}
	if node.Right != nil {
		PreOrder(node.Right)
	}
}
Assim como ele não usa duas funções separadas
*/

// === Travessias ===

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

// === Remoção de Elemento ===

func (t *BST) Remove(value int) {
	var removed bool
	t.root, removed = Remove(t.root, value)
	if removed {
		t.size--
	}
}

func Remove(node *TreeNode, value int) (*TreeNode, bool) {
	removed := false // Flag para rastrear se a remoção ocorreu nesta sub-árvore
	if node == nil {
		return nil, false // Valor não encontrado, nenhuma remoção
	}
	if value < node.Value {
		node.Left, removed = Remove(node.Left, value)
	} else if value > node.Value {
		node.Right, removed = Remove(node.Right, value)
	} else {
		removed = true // Remoção vai acontecer (ou já aconteceu nos casos abaixo)
		// Nó com apenas um filho ou nenhum filho
		if node.Left == nil {
			return node.Right, true
		} else if node.Right == nil {
			return node.Left, true
		}
		// Nó com dois filhos: obtém o sucessor inorder (menor na subárvore direita)
		minRight := Min(node.Right)
		node.Value = minRight
		node.Right, _ = Remove(node.Right, minRight)
	}
	return node, removed
}

// === Verificação de BST ===

// IsBst é o método público que inicia a verificação a partir do nó receptor.
// Ele chama a função auxiliar com os limites máximo e mínimo possíveis.
func (bstNode *TreeNode) IsBst() bool {
	// Foi usado o math.MinInt e o math.MaxInt para representar -infinito e +infinito
	// como limites iniciais para a raiz.
	return isBstRecursive(bstNode, math.MinInt, math.MaxInt)
}

// isBstRecursive é a função recursiva que faz a verificação.
func isBstRecursive(node *TreeNode, min int, max int) bool {
	// Caso 1: Uma árvore vazia (nó nulo) é considerada uma BST válida.
	if node == nil {
		return true
	}

	// Caso 2: Verifica se o valor do nó atual está fora dos limites permitidos.
	// Está sendo usado a convenção: left <= node < right (ou seja, valores iguais podem ir para a esquerda)
	// Se o valor for menor ou igual ao mínimo OU maior que o máximo, não é BST.
	if node.Value <= min || node.Value > max {
		return false
	}

	// Passo recursivo:
	// 1. Verifica a sub-árvore esquerda:
	//    Todos os nós à esquerda devem ser menores ou iguais ao nó atual (novo 'max')
	//    e maiores que o limite 'min' herdado.
	// 2. Verifica a sub-árvore direita:
	//    Todos os nós à direita devem ser maiores que o nó atual (novo 'min')
	//    e menores ou iguais ao limite 'max' herdado.
	// A árvore só é BST se o nó atual for válido E ambas as sub-árvores forem BSTs válidas.
	return isBstRecursive(node.Left, min, node.Value) && isBstRecursive(node.Right, node.Value, max)
}

// === Size ===

// Size retorna o número de nós na árvore BST.
func (t *BST) Size() int {
	return t.size // Simplesmente retorna o valor do campo size
}

// === Conversão de Array Ordenado para BST ===

// ConvertToBalancedBst constrói uma BST balanceada a partir de um slice
// Lembrar de passar os índices corretos (ex: 0 e len(v)-1).
func ConvertToBalancedBst(v []int, ini int, fim int) *TreeNode {
	// Caso Base: Se o sub-array é inválido/vazio, retorna nil.
	if ini > fim {
		return nil
	}

	// Encontra o elemento do meio
	mid := ini + (fim-ini)/2

	// O elemento do meio se torna a raiz desta sub-árvore
	rootNode := &TreeNode{Value: v[mid]}

	// Constrói recursivamente a sub-árvore esquerda
	// (com os elementos de 'ini' até 'mid-1')
	rootNode.Left = ConvertToBalancedBst(v, ini, mid-1)

	// Constrói recursivamente a sub-árvore direita
	// (com os elementos de 'mid+1' até 'fim')
	rootNode.Right = ConvertToBalancedBst(v, mid+1, fim)

	// Retorna a raiz da sub-árvore criada
	return rootNode
}

// CountEven é o método público na BST (wrapper).
func (t *BST) CountEven() int {
	// Chama a função recursiva 'Par' no nó raiz
	return Par(t.root)
}

// Par (conforme protótipo da Questão 8) é a função recursiva
// que conta quantos nós na sub-árvore (iniciando em 'node')
// possuem um valor par.
func Par(node *TreeNode) int {
	// Caso Base: Se o nó é nulo, a contagem de pares é 0.
	if node == nil {
		return 0
	}

	// 1. Conta recursivamente na sub-árvore esquerda
	countLeft := Par(node.Left)

	// 2. Conta recursivamente na sub-árvore direita
	countRight := Par(node.Right)

	// 3. Verifica o nó atual
	countCurrent := 0
	if node.Value%2 == 0 {
		countCurrent = 1
	}

	// 4. Retorna a soma total
	return countCurrent + countLeft + countRight
}

// TODO: melhorar os comentários das implementações (principalmente as recursivas)
