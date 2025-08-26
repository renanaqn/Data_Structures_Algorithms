package main

import (
	"DSA/datastructures"
	"fmt"
)

func main() {
	fmt.Println("Trabalhando com Pilha")

	// Para usar a pilha, é preciso lembrar de prefixar com o nome do pacote
	pilha := datastructures.NewStackArrayList()

	fmt.Println("Pilha está vazia?", pilha.IsEmpty())

	pilha.Push(10)
	pilha.Push(20)

	topo, _ := pilha.Peek()
	fmt.Println("Topo:", topo) // Deve ser 20

	valor, _ := pilha.Pop()
	fmt.Println("Pop:", valor) // Deve ser 20

	fmt.Println("Tamanho:", pilha.Size()) // Deve ser 1

	fmt.Println("Trabalhando com ArrayList")

	l := &datastructures.ArrayList{}
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

	l2 := &datastructures.LinkedList{}
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

	l3 := &datastructures.DoublyLinkedList{}
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
