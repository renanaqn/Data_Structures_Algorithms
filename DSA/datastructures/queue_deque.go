package datastructures

import "errors"

// ==== Interface do Deque ====

// IDeque define a interface para uma Fila de Duas Pontas (Deque).
type IDeque interface {
	EnqueueFront(value int) error
	EnqueueRear(value int) error
	DequeueFront() (int, error)
	DequeueRear() (int, error)
	Front() (int, error)
	Rear() (int, error)
	IsEmpty() bool
	Size() int
}

// ==== Implementação com Array Circular (ArrayDeque) ====

// ArrayDeque implementa um Deque usando um array circular de capacidade fixa.
type ArrayDeque struct {
	items    []int
	front    int
	rear     int
	size     int
	capacity int
}

// NewArrayDeque cria um novo Deque com capacidade fixa.
func NewArrayDeque(capacity int) *ArrayDeque {
	return &ArrayDeque{
		items:    make([]int, capacity),
		front:    0,
		rear:     0,
		size:     0,
		capacity: capacity,
	}
}

func (d *ArrayDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *ArrayDeque) isFull() bool {
	return d.size == d.capacity
}

func (d *ArrayDeque) Size() int {
	return d.size
}

// EnqueueRear adiciona um elemento no final (na 'rear' da fila).
func (d *ArrayDeque) EnqueueRear(value int) error {
	if d.isFull() {
		return errors.New("deque está cheio")
	}
	d.items[d.rear] = value
	d.rear = (d.rear + 1) % d.capacity
	d.size++
	return nil
}

// EnqueueFront adiciona um elemento no início (na 'front' da fila).
func (d *ArrayDeque) EnqueueFront(value int) error {
	if d.isFull() {
		return errors.New("deque está cheio")
	}
	d.front = (d.front - 1 + d.capacity) % d.capacity
	d.items[d.front] = value
	d.size++
	return nil
}

// DequeueFront remove um elemento do início.
func (d *ArrayDeque) DequeueFront() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque está vazio")
	}
	item := d.items[d.front]
	d.front = (d.front + 1) % d.capacity
	d.size--
	return item, nil
}

// DequeueRear remove um elemento do final.
func (d *ArrayDeque) DequeueRear() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque está vazio")
	}
	d.rear = (d.rear - 1 + d.capacity) % d.capacity
	item := d.items[d.rear]
	d.size--
	return item, nil
}

// Front espia o primeiro elemento.
func (d *ArrayDeque) Front() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque está vazio")
	}
	return d.items[d.front], nil
}

// Rear espia o último elemento.
func (d *ArrayDeque) Rear() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque está vazio")
	}
	rearIndex := (d.rear - 1 + d.capacity) % d.capacity
	return d.items[rearIndex], nil
}

// ==== Implementação com Lista Duplamente Ligada (DoublyLinkedListDeque) ====

// DoublyLinkedListDeque implementa um Deque usando a DoublyLinkedList existente.
type DoublyLinkedListDeque struct {
	list *DoublyLinkedList
}

// NewDoublyLinkedListDeque cria um novo Deque dinâmico.
func NewDoublyLinkedListDeque() *DoublyLinkedListDeque {
	return &DoublyLinkedListDeque{list: &DoublyLinkedList{}}
}

func (d *DoublyLinkedListDeque) IsEmpty() bool {
	return d.list.Size() == 0
}

func (d *DoublyLinkedListDeque) Size() int {
	return d.list.Size() // Usa o Size() da lista interna.
}

// EnqueueFront adiciona no início da lista (no 'head').
func (d *DoublyLinkedListDeque) EnqueueFront(value int) error {
	return d.list.AddOnIndex(value, 0)
}

// EnqueueRear adiciona no final da lista (no 'tail').
func (d *DoublyLinkedListDeque) EnqueueRear(value int) error {
	d.list.Add(value)
	return nil
}

// DequeueFront remove do início da lista (o 'head').
func (d *DoublyLinkedListDeque) DequeueFront() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque está vazio")
	}
	val, _ := d.list.Get(0)
	d.list.Remove(0)
	return val, nil
}

// DequeueRear remove do final da lista (o 'tail').
func (d *DoublyLinkedListDeque) DequeueRear() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque está vazio")
	}
	lastIndex := d.list.Size() - 1
	val, _ := d.list.Get(lastIndex)
	d.list.Remove(lastIndex)
	return val, nil
}

// Front espia o primeiro elemento (o 'head' da lista).
func (d *DoublyLinkedListDeque) Front() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque está vazio")
	}
	return d.list.Get(0)
}

// Rear espia o último elemento (o 'tail' da lista).
func (d *DoublyLinkedListDeque) Rear() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque está vazio")
	}
	return d.list.Get(d.list.Size() - 1)
}
