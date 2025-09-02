package datastructures

import "errors"

// ==== Interface da Fila ====

type Queue interface {
	Enqueue(e int)         // Adiciona ao fim
	Dequeue() (int, error) // Remove do início
	Peek() (int, error)    // Espia o início
	Size() int
	IsEmpty() bool
}

type CircularQueue struct {
	items []int
	head  int
	tail  int
	size  int
}

// NewCircularQueue cria uma fila com capacidade fixa.
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items: make([]int, capacity),
		head:  0,
		tail:  -1, // Começa em -1 para indicar que está vazia
		size:  0,
	}
}

func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *CircularQueue) isFull() bool {
	return q.size == len(q.items)
}

func (q *CircularQueue) Size() int {
	return q.size
}

// Enqueue adiciona ao fim usando a lógica circular.
func (q *CircularQueue) Enqueue(e int) error {
	if q.isFull() {
		return errors.New("a fila está cheia")
	}
	// Avança a traseira de forma circular
	q.tail = (q.tail + 1) % len(q.items)
	q.items[q.tail] = e
	q.size++
	return nil
}

// Dequeue remove do início usando a lógica circular.
func (q *CircularQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("a fila está vazia")
	}
	item := q.items[q.head]
	// Avança a frente de forma circular
	q.head = (q.head + 1) % len(q.items)
	q.size--
	return item, nil
}

// Peek espia o elemento da frente.
func (q *CircularQueue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("a fila está vazia")
	}
	return q.items[q.head], nil
}

// ==== Implementação com ArrayList ====
// ineficiente para Dequeue

type QueueArrayList struct {
	list *ArrayList
}

func NewQueueArrayList() *QueueArrayList {
	list := &ArrayList{}
	list.Init(10) // Inicializa com capacidade 10
	return &QueueArrayList{list: list}
}

func (q *QueueArrayList) Size() int {
	return q.list.Size()
}

func (q *QueueArrayList) IsEmpty() bool {
	return q.list.Size() == 0
}

// Enqueue adiciona no fim do array - O(1) amortizado
func (q *QueueArrayList) Enqueue(e int) {
	q.list.Add(e)
}

// Dequeue remove do início do array - MUITO LENTO (O(n))
func (q *QueueArrayList) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("a fila está vazia")
	}
	val, _ := q.list.Get(0)
	q.list.Remove(0) // Esta operação desloca todos os elementos
	return val, nil
}

// Peek espia o início do array - Rápido (O(1))
func (q *QueueArrayList) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("a fila está vazia")
	}
	return q.list.Get(0)
}

// ==== Implementação com DoublyLinkedList ====
type QueueLinkedList struct {
	list *DoublyLinkedList
}

func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{list: &DoublyLinkedList{}}
}

func (q *QueueLinkedList) Size() int {
	return q.list.Size()
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.list.Size() == 0
}

// Enqueue adiciona no final da lista (tail) - O(1)
func (q *QueueLinkedList) Enqueue(e int) {
	q.list.Add(e)
}

// Dequeue remove do início da lista (head) - O(1)
func (q *QueueLinkedList) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("a fila está vazia")
	}
	val, _ := q.list.Get(0)
	q.list.Remove(0)
	return val, nil
}

// Peek espia o início da lista (head) - O(1)
func (q *QueueLinkedList) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("a fila está vazia")
	}
	return q.list.Get(0)
}
