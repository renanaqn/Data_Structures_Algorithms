package main

import (
	"DSA/datastructures" // IMPORTANTE: Importe seu novo pacote!
	"fmt"
)

func main() {
	fmt.Println("Trabalhando com Pilha (usando o pacote datastructures)")

	// Agora, para usar sua pilha, você precisa prefixar com o nome do pacote.
	pilha := datastructures.NewStackArrayList()

	fmt.Println("Pilha está vazia?", pilha.IsEmpty())

	pilha.Push(10)
	pilha.Push(20)

	topo, _ := pilha.Peek()
	fmt.Println("Topo:", topo) // Deve ser 20

	valor, _ := pilha.Pop()
	fmt.Println("Pop:", valor) // Deve ser 20

	fmt.Println("Tamanho:", pilha.Size()) // Deve ser 1
}
