package datastructures

import "errors"

// ==== Interface da Pilha ====

type Stack interface {
	Push(e int)
	Pop() (int, error)
	Peek() (int, error) // topo()
	Size() int
	IsEmpty() bool
}

// ==== Implementação de Pilha usando ideia de Índice ====

// StackWithIndex é uma implementação que gerencia o topo com um índice explícito.
type StackWithIndex struct {
	items []int // O slice que armazena os dados
	top   int   // O índice que aponta para a PRÓXIMA posição livre
}

// NewStackWithIndex cria uma nova pilha com uma capacidade inicial.
func NewStackWithIndex(initialCapacity int) *StackWithIndex {
	return &StackWithIndex{
		items: make([]int, initialCapacity), // Aloca o slice com a capacidade definida
		top:   0,                            // O topo começa no índice 0
	}
}

// Push implementa a lógica de "add" e "update top".
func (s *StackWithIndex) Push(e int) {
	// Se o topo atingiu a capacidade do slice, não podemos mais inserir.
	// (Uma implementação mais robusta poderia redimensionar o slice aqui)
	if s.top == len(s.items) {
		// Por enquanto, vamos ignorar a inserção se estiver cheio.
		// Em um caso real, retornaríamos um erro.
		return
	}

	// Passo 1: "add"
	s.items[s.top] = e
	// Passo 2: "update top"
	s.top++
}

// Pop faz o processo inverso.
func (s *StackWithIndex) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("a pilha está vazia")
	}
	// Passo 1: Decrementa o topo para apontar para o último elemento válido.
	s.top--
	// Passo 2: Retorna o elemento que estava no topo.
	item := s.items[s.top]
	return item, nil
}

// Peek "espia" o elemento na posição `top - 1`.
func (s *StackWithIndex) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("a pilha está vazia")
	}
	return s.items[s.top-1], nil
}

// Size é simplesmente o valor do nosso índice `top`.
func (s *StackWithIndex) Size() int {
	return s.top
}

// IsEmpty verifica se o `top` é 0.
func (s *StackWithIndex) IsEmpty() bool {
	return s.top == 0
}

// ==== Implementação de Pilha usando ArrayList ====

type StackArrayList struct {
	list *ArrayList
}

// NewStackArrayList cria uma nova pilha.
func NewStackArrayList() *StackArrayList {
	newList := &ArrayList{}
	newList.Init(10)

	return &StackArrayList{list: newList}
}

// Size retorna o número de elementos na pilha.
func (s *StackArrayList) Size() int {
	return s.list.Size()
}

// IsEmpty verifica se a pilha está vazia.     '-'
func (s *StackArrayList) IsEmpty() bool {
	return s.list.Size() == 0
}

// Push adiciona um elemento no topo da pilha.
func (s *StackArrayList) Push(e int) {
	s.list.Add(e)
}

// Peek (ou topo) retorna o elemento do topo sem removê-lo.
func (s *StackArrayList) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("a pilha está vazia")
	}
	return s.list.Get(s.list.Size() - 1)
}

// Pop remove e retorna o elemento do topo da pilha.
func (s *StackArrayList) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("a pilha está vazia")
	}
	val, err := s.list.Get(s.list.Size() - 1)
	if err != nil {
		return 0, err
	}
	s.list.Remove(s.list.Size() - 1)
	return val, nil
}

// ==== Implementação de Pilha usando LinkedList ====
type StackLinkedList struct {
	list *LinkedList
}

// NewStackLinkedList cria uma nova pilha.
func NewStackLinkedList() *StackLinkedList {
	return &StackLinkedList{list: &LinkedList{}}
}

// Size retorna o número de elementos na pilha.
func (s *StackLinkedList) Size() int {
	return s.list.Size()
}

// IsEmpty verifica se a pilha está vazia.
func (s *StackLinkedList) IsEmpty() bool {
	return s.list.Size() == 0
}

// Push adiciona um elemento no topo da pilha.
func (s *StackLinkedList) Push(e int) {
	s.list.AddOnIndex(e, 0)
}

// Peek (ou topo) retorna o elemento do topo sem removê-lo.
func (s *StackLinkedList) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("a pilha está vazia")
	}
	return s.list.Get(0)
}

// Pop remove e retorna o elemento do topo da pilha.
func (s *StackLinkedList) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("a pilha está vazia")
	}
	val, _ := s.list.Get(0)
	s.list.Remove(0)
	return val, nil
}

type StackRaw struct {
	data []int
	top  int
}

// NewStackRaw cria uma nova pilha.
func NewStackRaw() *StackRaw {
	return &StackRaw{data: make([]int, 10), top: -1}
}

// Size retorna o número de elementos na pilha.
func (s *StackRaw) Size() int {
	return s.top + 1
}

// IsEmpty verifica se a pilha está vazia.
func (s *StackRaw) IsEmpty() bool {
	return s.top == -1
}

// Push adiciona um elemento no topo da pilha.
func (s *StackRaw) Push(e int) {
	if s.top+1 == len(s.data) {
		newData := make([]int, len(s.data)*2)
		copy(newData, s.data)
		s.data = newData
	}
	s.top++
	s.data[s.top] = e
}

// Peek (ou topo) retorna o elemento do topo sem removê-lo.
func (s *StackRaw) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("a pilha está vazia")
	}
	return s.data[s.top], nil
}

// Pop remove e retorna o elemento do topo da pilha.
func (s *StackRaw) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("a pilha está vazia")
	}
	val := s.data[s.top]
	s.top--
	return val, nil
}
