package main

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

// Implementação de ArrayList

type ArrayList struct {
	v        []int
	inserted int
}

func (l *ArrayList) Init(size int) {
	l.v = make([]int, size)
}

func (list *ArrayList) Size() int { //Theta(1)
	return list.inserted
}

func (list *ArrayList) Get(index int) (int, error) { //Theta(1)
	if index >= 0 && index < list.inserted {
		return list.v[index], nil
	} else {
		return -1, errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

func (list *ArrayList) doubleV() { //Theta(n)
	newV := make([]int, list.inserted*2)
	for i := 0; i < len(list.v); i++ {
		newV[i] = list.v[i]
	}
	list.v = newV
}

func (list *ArrayList) Add(val int) { //O(n), Ômega(1)
	if list.inserted == len(list.v) {
		list.doubleV()
	}
	list.v[list.inserted] = val
	list.inserted++
}

func (list *ArrayList) AddOnIndex(val int, index int) error { //O(n), Ômega(1)
	if index >= 0 && index <= list.inserted {
		if list.inserted == len(list.v) {
			list.doubleV()
		}
		for i := list.inserted; i > index; i-- {
			list.v[i] = list.v[i-1]
		}
		list.v[index] = val
		list.inserted++
		return nil
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

func (list *ArrayList) Remove(index int) error { //Ômega(1), O(n)
	if index >= 0 && index < list.inserted {
		for i := index; i < list.inserted; i++ {
			list.v[i] = list.v[i+1]
		}
		list.inserted--
		return nil
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

func (list *ArrayList) Print() {
	for i := 0; i < list.inserted; i++ {
		fmt.Print(list.v[i], " ")
	}
	fmt.Println()
}

// Implementação de LinkedList

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

// Implementação de DoublyLinkedList

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

	l.AddOnIndex(-1, 0)

	val, _ = l.Get(0)
	fmt.Println("Valor na posicao 0: ", val)

	l.Remove(0)

	val, _ = l.Get(0)
	fmt.Println("Valor na posicao 0: ", val)

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

// TODO: Mostrar também o endereço de memória de cada elemento
