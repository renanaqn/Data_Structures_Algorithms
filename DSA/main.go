package main

import (
	"DSA/datastructures"
	"fmt"
)

func main() {
	fmt.Println("Trabalhando com Pilha em cima de ArrayList")

	// Para usar a pilha, é preciso lembrar de prefixar com o nome do pacote
	pilha := datastructures.NewStackArrayList()

	fmt.Println("Pilha está vazia?", pilha.IsEmpty())

	pilha.Push(5)
	pilha.Push(9)
	pilha.Push(25)

	topo, _ := pilha.Peek()
	fmt.Println("Topo:", topo)

	pilha.Pop()
	valor, _ := pilha.Pop()
	fmt.Println("Pop:", valor)

	fmt.Println("Pilha está vazia?", pilha.IsEmpty())

	fmt.Println("Tamanho da Pilha:", pilha.Size())

	topo, _ = pilha.Peek()
	fmt.Println("Topo:", topo)

	pilha.Pop()

	fmt.Println("")
	fmt.Println("Trabalhando com Pilha em cima de LinkedList")

	pilha2 := datastructures.NewStackLinkedList()

	fmt.Println("Pilha 2 está vazia?", pilha2.IsEmpty())

	pilha2.Push(7)
	pilha2.Push(15)
	pilha2.Push(23)

	topo2, _ := pilha2.Peek()
	fmt.Println("Topo:", topo2) // Deve ser 23

	valor2, _ := pilha2.Pop()
	fmt.Println("Pop:", valor2) // Deve ser 23

	fmt.Println("Tamanho:", pilha2.Size()) // Deve ser 2

	fmt.Println("Pilha 2 está vazia agora?", pilha2.IsEmpty())

	fmt.Println("")
	fmt.Println("Trabalhando com Pilha em cima de ArrayList de forma manual")

	// ===========================================
	/*
		fmt.Println("==== Trabalhando com ArrayList ==== ")

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
		fmt.Println("==== Trabalhando com LinkedList ==== ")

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
		fmt.Println("==== Trabalhando com DoublyLinkedList ==== ")

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
	*/
}
