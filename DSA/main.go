package main

import (
	"DSA/datastructures"
	"fmt"
)

func main() {

	fmt.Println("Trabalhando com pilha em cima de Array com Índice")

	pilhaIndice := datastructures.NewStackWithIndex(10)

	fmt.Println("pilha com Índice está vazia?", pilhaIndice.IsEmpty())

	pilhaIndice.Push(5)
	pilhaIndice.Push(9)
	pilhaIndice.Push(25)

	topoIndice, _ := pilhaIndice.Peek()
	fmt.Println("Topo:", topoIndice)

	pilhaIndice.Pop()
	valorIndice, _ := pilhaIndice.Pop()
	fmt.Println("Pop:", valorIndice)

	fmt.Println("pilhaIndice está vazia?", pilhaIndice.IsEmpty())

	fmt.Println("Tamanho da pilha com Índice:", pilhaIndice.Size())

	topoIndice, _ = pilhaIndice.Peek()
	fmt.Println("Topo:", topoIndice)

	fmt.Println("")
	fmt.Println("Trabalhando com pilhaArray em cima de ArrayList")

	pilhaArray := datastructures.NewStackArrayList()

	fmt.Println("pilhaArray está vazia?", pilhaArray.IsEmpty())

	pilhaArray.Push(5)
	pilhaArray.Push(9)
	pilhaArray.Push(25)

	topoArray, _ := pilhaArray.Peek()
	fmt.Println("Topo:", topoArray)

	pilhaArray.Pop()
	valorArray, _ := pilhaArray.Pop()
	fmt.Println("Pop:", valorArray)

	fmt.Println("pilhaArray está vazia?", pilhaArray.IsEmpty())

	fmt.Println("Tamanho:", pilhaArray.Size())

	topoArray, _ = pilhaArray.Peek()
	fmt.Println("topoArray:", topoArray)

	pilhaArray.Pop()

	fmt.Println("")
	fmt.Println("Trabalhando com pilhaArray em cima de LinkedList")

	pilhaLinked := datastructures.NewStackLinkedList()

	fmt.Println("pilhaLinked está vazia?", pilhaLinked.IsEmpty())

	pilhaLinked.Push(5)
	pilhaLinked.Push(9)
	pilhaLinked.Push(25)

	topoLinked, _ := pilhaLinked.Peek()
	fmt.Println("Topo:", topoLinked)

	pilhaLinked.Pop()
	valorLinked, _ := pilhaLinked.Pop()
	fmt.Println("Pop:", valorLinked)

	fmt.Println("pilhaLinked está vazia agora?", pilhaLinked.IsEmpty())

	fmt.Println("Tamanho:", pilhaLinked.Size())

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
