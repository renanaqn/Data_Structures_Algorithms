package main

import (
	"DSA/datastructures"
	"fmt"
)

func testaList() {
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
}

func testaStack() {
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
}

func testaQueue() {
	fmt.Println("\n--- Testando Fila Circular com Array ---")

	filaCircular := datastructures.NewCircularQueue(5)

	fmt.Println("Fila Circular está vazia?", filaCircular.IsEmpty())

	filaCircular.Enqueue(5)
	filaCircular.Enqueue(9)
	filaCircular.Enqueue(25)

	head, _ := filaCircular.Peek()
	fmt.Println("Head:", head)

	filaCircular.Dequeue()
	valor, _ := filaCircular.Dequeue()
	fmt.Println("Dequeue:", valor)

	fmt.Println("Fila Circular está vazia?", filaCircular.IsEmpty())
	fmt.Println("Tamanho da Fila Circular:", filaCircular.Size())

	head, _ = filaCircular.Peek()
	fmt.Println("Head:", head)

	filaCircular.Dequeue()
	fmt.Println("Fila esvaziada. Está vazia?", filaCircular.IsEmpty())

	// ==== Testando a Fila com ArrayList ====
	fmt.Println("\n--- Testando Fila com ArrayList ---")

	filaArray := datastructures.NewQueueArrayList()

	fmt.Println("Fila Array está vazia?", filaArray.IsEmpty())

	filaArray.Enqueue(5)
	filaArray.Enqueue(9)
	filaArray.Enqueue(25)

	headArray, _ := filaArray.Peek()
	fmt.Println("Head:", headArray)

	filaArray.Dequeue() // operação LENTA
	valorArray, _ := filaArray.Dequeue()
	fmt.Println("Dequeue:", valorArray)

	fmt.Println("Fila Array está vazia?", filaArray.IsEmpty())
	fmt.Println("Tamanho da Fila Array:", filaArray.Size())

	headArray, _ = filaArray.Peek()
	fmt.Println("Head:", headArray)

	filaArray.Dequeue()
	fmt.Println("Fila esvaziada. Está vazia?", filaArray.IsEmpty())

	// ==== Testando a Fila com DoublyLinkedList ====
	fmt.Println("\n--- Testando Fila com DoublyLinkedList ---")

	filaLinked := datastructures.NewQueueLinkedList()

	fmt.Println("Fila Linked está vazia?", filaLinked.IsEmpty())

	filaLinked.Enqueue(5)
	filaLinked.Enqueue(9)
	filaLinked.Enqueue(25)

	headLinked, _ := filaLinked.Peek()
	fmt.Println("Head:", headLinked)

	filaLinked.Dequeue()
	valorLinked, _ := filaLinked.Dequeue()
	fmt.Println("Dequeue:", valorLinked)

	fmt.Println("Fila Linked está vazia?", filaLinked.IsEmpty())
	fmt.Println("Tamanho da Fila Linked:", filaLinked.Size())

	headLinked, _ = filaLinked.Peek()
	fmt.Println("Head:", headLinked)

	filaLinked.Dequeue()
	fmt.Println("Fila esvaziada. Está vazia?", filaLinked.IsEmpty())

}

func testaDeque() {
	fmt.Println("\n--- Testando Deque com Array Circular ---")

	// Criamos um deque com capacidade para 5 elementos.
	deque := datastructures.NewArrayDeque(5)

	fmt.Println("Está vazio?", deque.IsEmpty())

	// Testando EnqueueRear (comportamento de Fila)
	deque.EnqueueRear(10) // Deque: [10]
	deque.EnqueueRear(20) // Deque: [10, 20]

	// Testando EnqueueFront (comportamento de Pilha)
	deque.EnqueueFront(5) // Deque: [5, 10, 20]
	deque.EnqueueFront(1) // Deque: [1, 5, 10, 20]

	front, _ := deque.Front()
	rear, _ := deque.Rear()
	fmt.Printf("Estado atual: Frente=%d, Traseira=%d, Tamanho=%d\n", front, rear, deque.Size())

	// Testando Dequeues
	val, _ := deque.DequeueFront()
	fmt.Println("DequeueFront retornou:", val)

	val, _ = deque.DequeueRear()
	fmt.Println("DequeueRear retornou:", val)

	front, _ = deque.Front()
	rear, _ = deque.Rear()
	fmt.Printf("Estado final: Frente=%d, Traseira=%d, Tamanho=%d\n", front, rear, deque.Size())

	fmt.Println("\n--- Testando Deque com Lista Duplamente Ligada ---")
	deque2 := datastructures.NewDoublyLinkedListDeque()

	fmt.Println("Está vazio?", deque.IsEmpty())

	// Mesma sequência de testes
	deque2.EnqueueRear(10)
	deque2.EnqueueRear(20)
	deque2.EnqueueFront(5)
	deque2.EnqueueFront(1)

	front, _ = deque2.Front()
	rear, _ = deque2.Rear()
	fmt.Printf("Estado atual: Frente=%d, Traseira=%d, Tamanho=%d\n", front, rear, deque2.Size())

	val, _ = deque2.DequeueFront()
	fmt.Println("DequeueFront retornou:", val)

	val, _ = deque2.DequeueRear()
	fmt.Println("DequeueRear retornou:", val)

	front, _ = deque2.Front()
	rear, _ = deque2.Rear()
	fmt.Printf("Estado final: Frente=%d, Traseira=%d, Tamanho=%d\n", front, rear, deque2.Size())
}

func testaSearch() {
	fmt.Println("\n--- Testando Algoritmos de Busca ---")

	unsortedData := []int{3, 41, 2, 9, 33, 15, 28, 49}
	sortedData := []int{2, 9, 15, 28, 33, 41, 49}
	target := 33

	// --- Teste da Busca Linear ---
	fmt.Println("\n[Busca Linear]")
	index1 := datastructures.LinearSearch(unsortedData, target)
	fmt.Printf("Buscando %d em dados não ordenados. Encontrado no índice: %d\n", target, index1)

	index2 := datastructures.LinearSearch(sortedData, target)
	fmt.Printf("Buscando %d em dados ordenados. Encontrado no índice: %d\n", target, index2)

	// --- Teste da Busca Binária ---
	fmt.Println("\n[Busca Binária]")
	index3 := datastructures.BinarySearch(sortedData, target)
	fmt.Printf("Buscando %d em dados ORDENADOS. Encontrado no índice: %d\n", target, index3)

	// Exemplo de por que a Busca Binária PRECISA de dados ordenados
	index4 := datastructures.BinarySearch(unsortedData, target)
	fmt.Printf("!! Tentando buscar %d em dados NÃO ORDENADOS. Resultado (não confiável): %d\n", target, index4)
}

func testaOrdenacao() {
	fmt.Println("\n--- Testando Algoritmos de Ordenação ---")

	// --- Selection Sort ---
	fmt.Println("\n[Selection Sort]")
	dados1 := []int{8, 2, 4, 3, 7}
	fmt.Println("Array Original:", dados1)
	datastructures.SelectionSort(dados1)
	fmt.Println("Array Ordenado:", dados1)

	// --- Bubble Sort ---
	fmt.Println("\n[Bubble Sort]")
	dados2 := []int{2, 8, 6, 10, 4, 5, 3}
	fmt.Println("Array Original:", dados2)
	datastructures.BubbleSort(dados2)
	fmt.Println("Array Ordenado:", dados2)

	// --- Insertion Sort ---
	fmt.Println("\n[Insertion Sort]")
	dados3 := []int{8, 2, 4, 3, 7}
	fmt.Println("Array Original:", dados3)
	datastructures.InsertionSort(dados3)
	fmt.Println("Array Ordenado:", dados3)

	// --- Merge Sort ---
	fmt.Println("\n[Merge Sort]")
	dados4 := []int{9, 4, 3, 6, 3, 2, 5, 7, 1, 8}
	fmt.Println("Array Original:", dados4)
	// MergeSort retorna um novo slice, então precisamos atribuir o resultado.
	dadosOrdenados := datastructures.MergeSort(dados4)
	fmt.Println("Array Ordenado:", dadosOrdenados)

	// --- Quick Sort ---
	fmt.Println("\n[Quick Sort]")
	dados5 := []int{9, 4, 3, 6, 3, 2, 8, 7, 1, 5}
	fmt.Println("Array Original:", dados5)
	datastructures.QuickSort(dados5)
	fmt.Println("Array Ordenado:", dados5)

	// --- Counting Sort ---
	fmt.Println("\n[Counting Sort]")
	dados6 := []int{4, 1, 3, 4, 3, 1, 2}
	fmt.Println("Array Original:", dados6)
	datastructures.CountingSort(dados6) // Esta versão é in-place
	fmt.Println("Array Ordenado:", dados6)
}

func testaArvores() {
	// --- Árvore de Busca Binária (BST) ---
	fmt.Println("\n[Árvore de Busca Binária]")
	// arvore := &datastructures.BST{}
	arvore := datastructures.NewBST()

	// Inserindo elementos
	// valores := []int{10, 5, 15, 3, 7, 12, 18}
	// valores := []int{40, 20, 25, 10, 50, 45, 60}
	// valores := []int{20, 10, 30, 5, 15, 8, 40}
	valores := []int{50, 30, 45, 32, 10, 90, 55, 49, -5, 88, 110, 40}
	//TODO: testar com a seguinte sequencia: 20, 10, 30, 5, 15, 8, 40
	// onde o 8 vai ficar? No lugar do 5 ou do 10?

	fmt.Println("Inserindo os valores:", valores)
	for _, value := range valores {
		arvore.Add(value)
	}

	// Testando a busca
	// fmt.Println("Buscando o valor 7...", arvore.Search(7))
	// fmt.Println("Buscando o valor 18...", arvore.Search(45))
	// fmt.Println("Buscando o valor 99...", arvore.Search(99))
	// fmt.Println("Buscando o valor 1...", arvore.Search(1))

	// //Testando Min e Max
	// fmt.Println("Valor Mínimo na árvore:", arvore.Min())
	// fmt.Println("Valor Máximo na árvore:", arvore.Max())

	// Testando altura
	fmt.Println("Altura da árvore:", arvore.Height())
	// Testando travessias
	// fmt.Println("Travessia PreOrder:")
	// arvore.PreOrder()
	// fmt.Println("\nTravessia InOrder:")
	// arvore.InOrder()
	// fmt.Println("\nTravessia PosOrder:")
	// arvore.PosOrder()
	fmt.Println("\nTravessia LevelOrder:")
	arvore.LevelOrder()
	fmt.Println("\nTamanho da árvore:", arvore.Size())

	// Testando remoção
	fmt.Println("\nRemovendo o valor 5...")
	arvore.Remove(45)
	fmt.Println("\nTravessia LevelOrder:")
	arvore.LevelOrder()
	fmt.Println("\nTamanho da árvore:", arvore.Size())
	fmt.Println()

}

func testaArvores2() {
	fmt.Println("\n--- Testando a Função isBst ---")

	// --- Árvore 1: É uma BST válida ---
	//       10
	//      /  \
	//     5    15
	//    / \   / \
	//   3   7 12 18
	fmt.Println("\nÁrvore 1 (Válida):")
	raizValida := &datastructures.TreeNode{Value: 10}
	raizValida.Left = &datastructures.TreeNode{Value: 5}
	raizValida.Right = &datastructures.TreeNode{Value: 15}
	raizValida.Left.Left = &datastructures.TreeNode{Value: 3}
	raizValida.Left.Right = &datastructures.TreeNode{Value: 7}
	raizValida.Right.Left = &datastructures.TreeNode{Value: 12}
	raizValida.Right.Right = &datastructures.TreeNode{Value: 18}

	fmt.Println("Verificando se é BST:", raizValida.IsBst()) // Esperado: true

	// --- Árvore 2: NÃO é uma BST válida ---
	//       10
	//      /  \
	//     5    15
	//    / \
	//   3   12  <-- Erro: 12 > 10, não deveria estar na sub-árvore esquerda
	fmt.Println("\nÁrvore 2 (Inválida):")
	raizInvalida := &datastructures.TreeNode{Value: 10}
	raizInvalida.Left = &datastructures.TreeNode{Value: 5}
	raizInvalida.Right = &datastructures.TreeNode{Value: 15}
	raizInvalida.Left.Left = &datastructures.TreeNode{Value: 3}
	raizInvalida.Left.Right = &datastructures.TreeNode{Value: 12} // Violação da BST

	fmt.Println("Verificando se é BST:", raizInvalida.IsBst()) // Esperado: false

	// --- Árvore 3: Outra inválida ---
	//       20
	//      /  \
	//     10   30
	//         /
	//        5   <-- Erro: 5 < 20, não deveria estar na sub-árvore direita
	fmt.Println("\nÁrvore 3 (Inválida):")
	raizInvalida2 := &datastructures.TreeNode{Value: 20}
	raizInvalida2.Left = &datastructures.TreeNode{Value: 10}
	raizInvalida2.Right = &datastructures.TreeNode{Value: 30}
	raizInvalida2.Right.Left = &datastructures.TreeNode{Value: 5} // Violação da BST

	fmt.Println("Verificando se é BST:", raizInvalida2.IsBst()) // Esperado: false

	// --- Árvore 4: Vazia (Considerada válida) ---
	fmt.Println("\nÁrvore 4 (Vazia):")
	var raizVazia *datastructures.TreeNode = nil
	fmt.Println("Verificando se é BST:", raizVazia.IsBst()) // Esperado: true

	// --- Árvore 5: Apenas Raiz (Considerada válida) ---
	fmt.Println("\nÁrvore 5 (Apenas Raiz):")
	raizSo := &datastructures.TreeNode{Value: 50}
	fmt.Println("Verificando se é BST:", raizSo.IsBst()) // Esperado: true

}

func testaConversaoBST() {
	fmt.Println("\n--- Testando Conversão de Array Ordenado para BST Balanceada ---")

	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println("Array ordenado de entrada:", sortedArray)

	rootNode := datastructures.ConvertToBalancedBst(sortedArray, 0, len(sortedArray)-1)

	// O usuário recebe um *TreeNode*, não a struct *BST*
	//    Então ele precisa criar a struct manualmente:
	balancedTree := datastructures.NewBST()
	balancedTree.SetRoot(rootNode)         // (Precisaríamos criar um método SetRoot)
	balancedTree.SetSize(len(sortedArray)) // (Precisaríamos criar um método SetSize)

	balancedTree.LevelOrder()                              // Esperado: [6 3 9 1 4 7 10 2 5 8 11] (ou similar)
	fmt.Println("Tamanho da árvore:", balancedTree.Size()) // Esperado: 11
	fmt.Println("Altura:", balancedTree.Height())          // Esperado: 3 (O(log n))
}

// testaPares (Questão 8)
func testaPares() {
	fmt.Println("\n--- Testando Questão 8: Contagem de Pares ---")

	vetor := []int{2, 4, 1, 10, 12, 5}
	fmt.Println("Vetor de entrada:", vetor)

	// 1. Construir a BST
	arvore := datastructures.NewBST()
	for _, v := range vetor {
		arvore.Add(v)
	}

	// 2. Visualizar a árvore (LevelOrder é bom para isso)
	fmt.Print("Árvore resultante (LevelOrder): ")
	arvore.LevelOrder()

	// 3. Chamar a função CountEven()
	// A árvore é:
	//       2
	//      / \
	//     1   4
	//          \
	//          10
	//          / \
	//         5   12
	//
	// Valores pares: 2, 4, 10, 12.
	// Resultado esperado: 4

	contagem := arvore.CountEven()
	fmt.Println("Quantidade de nós com valores pares:", contagem) // Esperado: 4
}

func testaAVL() {
	fmt.Println("\n--- Testando Árvore AVL ---")
	avl := datastructures.NewAVL()

	// --- Teste 1: Inserção Ordenada (Pior caso para BST) ---
	fmt.Println("\n[Teste 1] Inserindo 1, 2, 3, 4, 5, 6, 7 sequencialmente...")
	// Isso forçaria várias rotações para manter a altura logarítmica
	inputs := []int{1, 2, 3, 4, 5, 6, 7}
	for _, val := range inputs {
		fmt.Printf("Inserindo %d...\n", val)
		avl.Add(val)
	}

	fmt.Println("\nEstado da Árvore após inserções (Em-Ordem):")
	// Esperado: Árvore balanceada.
	// A raiz deve ser o 4.
	// Estrutura esperada (visual):
	//       4
	//     /   \
	//    2     6
	//   / \   / \
	//  1   3 5   7
	// Altura esperada: 2 (considerando folha h=0) ou 3 (se contar nós)
	avl.InOrder()
	fmt.Println("\nEstado da Árvore após inserções (Nível):")
	avl.LevelOrder()

	// --- Teste 2: Remoção e Rebalanceamento ---
	fmt.Println("\n\n[Teste 2] Removendo elementos para forçar rotações...")

	// Remover o 4 (Raiz) deve forçar uma reestruturação
	fmt.Println("Removendo 4 (Raiz)...")
	avl.Remove(4)
	avl.InOrder()

	// Remover o 6 deve desbalancear o lado direito e forçar rotação
	fmt.Println("\nRemovendo 6...")
	avl.Remove(6)
	avl.InOrder()

	// Remover o 5
	fmt.Println("\nRemovendo 5...")
	avl.Remove(5)
	avl.InOrder()

	fmt.Println("\nTestes de AVL concluídos.")
}

// Função dedicada para testar a Tabela Hash
func testaHash() {
	fmt.Println("\n--- Testando Tabela Hash ---")

	// 1. Criação
	// Criamos uma tabela propositalmente pequena (capacidade 5)
	// para forçar colisões e visualizar o tratamento (listas ligadas).
	ht := datastructures.NewHashTable(5)

	// 2. Inserção (Put)
	fmt.Println("--- Inserindo dados ---")
	ht.Put("Alice", 25)
	ht.Put("Bob", 30)
	ht.Put("Carol", 35)

	fmt.Println("Fator de Carga (Load Factor):", ht.LoadFactor())
	fmt.Println("Capacidade da Tabela:", ht.CapacityHT())
	fmt.Println("Tamanho da Tabela (Size):", ht.SizeHT())

	ht.Put("Dave", 40) // Colisão provável (dependendo do hash)
	ht.Put("Eve", 45)  // Colisão provável

	// Exibe o estado atual para ver onde cada chave caiu
	ht.Print()

	// 3. Busca (Get)
	fmt.Println("\n--- Buscando ---")
	chaveBusca := "Bob"
	val, ok := ht.Get(chaveBusca)
	if ok {
		fmt.Printf("Chave '%s' encontrada! Valor: %d\n", chaveBusca, val)
	} else {
		fmt.Printf("Chave '%s' não encontrada.\n", chaveBusca)
	}

	// 4. Atualização (Put em chave existente)
	fmt.Println("\n--- Atualizando ---")
	fmt.Println("Atualizando valor de 'Alice' para 99...")
	ht.Put("Alice", 99)

	val, _ = ht.Get("Alice")
	fmt.Printf("Novo valor de Alice: %d\n", val)

	fmt.Println("Fator de Carga (Load Factor):", ht.LoadFactor())
	fmt.Println("Capacidade da Tabela:", ht.CapacityHT())
	fmt.Println("Tamanho da Tabela (Size):", ht.SizeHT())
	// 5. Remoção (Remove)
	fmt.Println("\n--- Removendo ---")
	fmt.Println("Removendo a chave 'Carol'...")
	removido := ht.Remove("Carol")
	fmt.Printf("Remoção bem-sucedida? %v\n", removido)

	// Exibe o estado final
	ht.Print()
}

func main() {
	// testaList()
	// testaStack()
	// testaQueue()
	// testaSearch()
	// testaDeque()
	// testaOrdenacao()
	// testaArvores()
	// testaArvores2()
	// testaConversaoBST()
	// testaPares()
	// testaAVL()
	testaHash()
}
