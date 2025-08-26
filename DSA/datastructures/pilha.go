package datastructures

import "errors" // Você precisará importar isso para usar errors.New

// ==== Interface da Pilha ====

type Stack interface {
	Push(e int)
	Pop() (int, error)
	Peek() (int, error) // Equivalente ao topo()
	Size() int
	IsEmpty() bool
}

// ==== Implementação de Pilha usando ArrayList ====

type StackArrayList struct {
	list *ArrayList
}

// NewStackArrayList cria uma nova pilha.
func NewStackArrayList() *StackArrayList {
	// Cria uma nova instância de ArrayList
	newList := &ArrayList{}
	// Inicializa a lista com uma capacidade padrão (ex: 10)
	newList.Init(10)
	// Retorna a nova pilha contendo a lista inicializada
	return &StackArrayList{list: newList}
}

// Size retorna o número de elementos na pilha.
func (s *StackArrayList) Size() int {
	return s.list.Size()
}

// IsEmpty verifica se a pilha está vazia.
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
