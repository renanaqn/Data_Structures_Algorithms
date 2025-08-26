package datastructures

import "errors" // Você precisará importar isso para usar errors.New

// ==== Interface da Pilha ====

type Stack interface {
	Push(e int)
	Pop() (int, error)
	Peek() (int, error) // topo()
	Size() int
	IsEmpty() bool
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
