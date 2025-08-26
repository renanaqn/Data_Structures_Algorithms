package datastructures

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int
	Get(index int) (int, error)
	Add(e int)
	AddOnIndex(e int, index int) error
	Remove(index int) error
}

// ==== Implementação de ArrayList ====

type ArrayList struct {
	v        []int
	inserted int
}

// Init cria e retorna uma nova instância de ArrayList.
// Complexidade: O(1) no pior caso, pois o tempo não depende da entrada
func (l *ArrayList) Init(size int) {
	l.v = make([]int, size)
}

// Size retorna o tamanho do Array
// Complexidade: theta(1) no melhor e pior caso, pois o tempo não depende da entrada
func (list *ArrayList) Size() int {
	return list.inserted
}

// Get retorna o elemento da posição index
// Complexidade: theta(1) no pior caso, pois o tempo não depende da entrada
func (list *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		return list.v[index], nil
	} else {
		return -1, errors.New(fmt.Sprintf("Index invalido: %dA lista tem %d elementos.", index, list.inserted))
	}
}

// doubleV duplica a capacidade do Array
// Complexidade: theta(n) no melhor e pior caso, pois o tempo depende de uma multiplicação
func (list *ArrayList) doubleV() { //Theta(n)
	newV := make([]int, list.inserted*2)
	for i := 0; i < len(list.v); i++ {
		newV[i] = list.v[i]
	}
	list.v = newV
}

// Add adiciona um novo elemento no final
// Complexidade: O(n) no pior caso, pois o tempo depende de uma multiplicação da doubleV
func (list *ArrayList) Add(val int) { //O(n), Ômega(1)
	if list.inserted == len(list.v) {
		list.doubleV()
	}
	list.v[list.inserted] = val
	list.inserted++
}

// AddOnIndex adiciona um novo valor em um local específico
// Complexidade: O(n) no pior caso, pois o tempo depende de uma multiplicação da doubleV
func (list *ArrayList) AddOnIndex(e int, index int) error {
	if index < 0 || index > list.inserted {
		return errors.New(fmt.Sprintf("Índice inválido: %d. O índice deve estar entre 0 e %d.", index, list.inserted))
	}

	if list.inserted == len(list.v) {
		list.doubleV()
	}
	//Desloca os elementos para a direita
	for i := list.inserted; i > index; i-- {
		list.v[i] = list.v[i-1]
	}

	// Insere o novo elemento na posição correta
	list.v[index] = e

	// Incrementa o contador de elementos inseridos
	list.inserted++

	return nil
}

// Remove remove o elemento de um índice específico.
// Complexidade: O(n) no pior caso, pois pode precisar deslocar todos os elementos.
func (list *ArrayList) Remove(index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New(fmt.Sprintf("Índice inválido: %d. A lista tem %d elementos.", index, list.inserted))
	}
	// Desloca os elementos
	for i := index; i < list.inserted-1; i++ {
		list.v[i] = list.v[i+1]
	}
	list.inserted--

	list.v[list.inserted] = 0 // Define o valor zero para o tipo int

	return nil
}

// Print imprime todos os elementos do ArrayList
// Complexidade: O(n) no pior caso, pois o tempo depende do número de elementos
func (list *ArrayList) Print() {
	for i := 0; i < list.inserted; i++ {
		fmt.Print(list.v[i], " ")
	}
	fmt.Println()
}

// ==== Implementação de LinkedList ====

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head     *Node
	inserted int
}

func (list *LinkedList) Size() int { //Theta(1)
	return list.inserted
}

func (list *LinkedList) Get(index int) (int, error) { //O(n), Ômega(1)
	if index >= 0 && index < list.inserted {
		aux := list.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		return aux.val, nil
	} else {
		return -1, errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

func (list *LinkedList) Add(val int) { //O(n), Ômega(1)
	newNode := &Node{val: val}
	if list.head == nil {
		list.head = newNode
	} else {
		aux := list.head
		for aux.next != nil {
			aux = aux.next
		}
		aux.next = newNode
	}
	list.inserted++
}

func (list *LinkedList) AddOnIndex(val int, index int) error { //O(n), Ômega(1)
	if index >= 0 && index <= list.inserted {
		newNode := &Node{val: val}
		if index == 0 {
			newNode.next = list.head
			list.head = newNode
		} else {
			aux := list.head
			for i := 0; i < index-1; i++ {
				aux = aux.next
			}
			newNode.next = aux.next
			aux.next = newNode
		}
		list.inserted++
		return nil
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

func (list *LinkedList) Remove(index int) error { //Ômega(1), O(n)
	if index >= 0 && index < list.inserted {
		if index == 0 {
			list.head = list.head.next
		} else {
			aux := list.head
			for i := 0; i < index-1; i++ {
				aux = aux.next
			}
			aux.next = aux.next.next
		}
		list.inserted--
		return nil
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

func (list *LinkedList) Print() {
	aux := list.head
	for aux != nil {
		fmt.Print(aux.val, " ")
		aux = aux.next
	}
	fmt.Println()
}

// ==== Implementação de DoublyLinkedList ====

type DoublyNode struct {
	val  int
	next *DoublyNode
	prev *DoublyNode
}

type DoublyLinkedList struct {
	head     *DoublyNode
	tail     *DoublyNode
	inserted int
}

func (list *DoublyLinkedList) Size() int { //Theta(1)
	return list.inserted
}

func (list *DoublyLinkedList) Get(index int) (int, error) { //O(n), Ômega(1)
	if index >= 0 && index < list.inserted {
		aux := list.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		return aux.val, nil
	} else {
		return -1, errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

func (list *DoublyLinkedList) Add(val int) { //O(n), Ômega(1)
	newNode := &DoublyNode{val: val}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
	list.inserted++
}

func (list *DoublyLinkedList) AddOnIndex(val int, index int) error { //O(n), Ômega(1)
	if index >= 0 && index <= list.inserted {
		newNode := &DoublyNode{val: val}
		if index == 0 {
			newNode.next = list.head
			if list.head != nil {
				list.head.prev = newNode
			}
			list.head = newNode
			if list.tail == nil {
				list.tail = newNode
			}
		} else if index == list.inserted {
			newNode.prev = list.tail
			if list.tail != nil {
				list.tail.next = newNode
			}
			list.tail = newNode
		} else {
			aux := list.head
			for i := 0; i < index-1; i++ {
				aux = aux.next
			}
			newNode.next = aux.next
			newNode.prev = aux
			if aux.next != nil {
				aux.next.prev = newNode
			}
			aux.next = newNode
		}
		list.inserted++
		return nil
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

func (list *DoublyLinkedList) Remove(index int) error { //Ômega(1), O(n)
	if index >= 0 && index < list.inserted {
		if index == 0 {
			list.head = list.head.next
			if list.head != nil {
				list.head.prev = nil
			} else {
				list.tail = nil
			}
		} else if index == list.inserted-1 {
			list.tail = list.tail.prev
			if list.tail != nil {
				list.tail.next = nil
			} else {
				list.head = nil
			}
		} else {
			aux := list.head
			for i := 0; i < index; i++ {
				aux = aux.next
			}
			aux.prev.next = aux.next
			if aux.next != nil {
				aux.next.prev = aux.prev
			}
		}
		list.inserted--
		return nil
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

func (list *DoublyLinkedList) Print() {
	aux := list.head
	for aux != nil {
		fmt.Print(aux.val, " ")
		aux = aux.next
	}
	fmt.Println()
}

/*
func main() {
	fmt.Println("Trabalhando com ArrayList")

	l := &ArrayList{}
	l.Init(10)
	for i := 1; i <= 50; i++ {
		l.Add(i)
	}
	val, _ := l.Get(0)
	fmt.Println("Valor na posicao 0: ", val)

	val, _ = l.Get(49)
	fmt.Println("Valor na posicao 49: ", val)

	l.AddOnIndex(-1, 30)

	val, _ = l.Get(30)
	fmt.Println("Valor na posicao 30: ", val)

	l.Remove(30)

	val, _ = l.Get(0)
	fmt.Println("Valor na posicao 30: ", val)

	l.Print()

	fmt.Println("")
	fmt.Println("Trabalhando com LinkedList")

	l2 := &LinkedList{}
	for i := 1; i <= 50; i++ {
		l2.Add(i)
	}
	val, _ = l2.Get(0)
	fmt.Println("Valor na posicao 0: ", val)

	val, _ = l2.Get(30)
	fmt.Println("Valor na posicao 49: ", val)

	l2.AddOnIndex(-1, 0)

	val, _ = l2.Get(0)
	fmt.Println("Valor na posicao 0: ", val)

	l2.Remove(0)

	val, _ = l2.Get(0)
	fmt.Println("Valor na posicao 0: ", val)

	l2.Print()

	fmt.Println("")
	fmt.Println("Trabalhando com DoublyLinkedList")

	l3 := &DoublyLinkedList{}
	for i := 1; i <= 50; i++ {
		l3.Add(i)
	}
	val, _ = l3.Get(0)
	fmt.Println("Valor na posicao 0: ", val)

	val, _ = l3.Get(30)
	fmt.Println("Valor na posicao 49: ", val)

	l3.AddOnIndex(-1, 0)
	val, _ = l3.Get(0)
	fmt.Println("Valor na posicao 0: ", val)

	l3.Remove(0)
	val, _ = l3.Get(0)
	fmt.Println("Valor na posicao 0: ", val)

	l3.Print()
}
*/
// TODO: Mostrar também o endereço de memória de cada elemento para fazer uns testes
