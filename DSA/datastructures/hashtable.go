package datastructures

import (
	"fmt"
)

type HashTableInterface interface {
	Put(key string, value int)
	Get(key string) (int, bool)
	Remove(key string) bool
	Print()
	SizeHT() int
	CapacityHT() int
	LoadFactor() float64
}

// HashNode representa um par chave-valor em um bucket (lista ligada).
type HashNode struct {
	Key   string
	Value int
	Next  *HashNode
}

// HashTable representa a tabela hash.
type HashTable struct {
	Buckets  []*HashNode
	Size     int // Quantidade total de elementos inseridos
	Capacity int // Tamanho do array de buckets
}

// NewHashTable cria uma nova tabela com uma capacidade inicial.
func NewHashTable(capacity int) *HashTable {
	return &HashTable{
		Buckets:  make([]*HashNode, capacity),
		Capacity: capacity,
		Size:     0,
	}
}

func (ht *HashTable) SizeHT() int {
	return ht.Size
}

func (ht *HashTable) CapacityHT() int {
	return ht.Capacity
}
func (ht *HashTable) LoadFactor() float64 {
	return float64(ht.Size) / float64(ht.Capacity)
}

// hashFunction gera um índice para a chave.
// Usa um algoritmo simples de soma ponderada para espalhar as chaves.
func (ht *HashTable) hashFunction(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		// (hash * 31 + char) é uma fórmula comum para strings
		hash = (hash*31 + int(key[i])) % ht.Capacity
	}
	return hash
}

// Put insere ou atualiza um valor na tabela.
func (ht *HashTable) Put(key string, value int) {
	index := ht.hashFunction(key)
	node := ht.Buckets[index]

	// 1. Verifica se a chave já existe (Atualização)
	for node != nil {
		if node.Key == key {
			node.Value = value // Atualiza valor
			return
		}
		node = node.Next
	}

	// 2. Se não existe, insere no início da lista (Colisão tratada aqui)
	newNode := &HashNode{
		Key:   key,
		Value: value,
		Next:  ht.Buckets[index], // O próximo do novo é o antigo primeiro
	}
	ht.Buckets[index] = newNode
	ht.Size++
}

// Get retorna o valor associado à chave e true, ou 0 e false se não existir.
func (ht *HashTable) Get(key string) (int, bool) {
	index := ht.hashFunction(key)
	node := ht.Buckets[index]

	for node != nil {
		if node.Key == key {
			return node.Value, true
		}
		node = node.Next
	}
	return 0, false
}

// Remove remove uma chave da tabela.
func (ht *HashTable) Remove(key string) bool {
	index := ht.hashFunction(key)
	current := ht.Buckets[index]
	var prev *HashNode

	for current != nil {
		if current.Key == key {
			// Achou o nó a remover
			if prev == nil {
				// É o primeiro da lista
				ht.Buckets[index] = current.Next
			} else {
				// Está no meio ou fim
				prev.Next = current.Next
			}
			ht.Size--
			return true
		}
		prev = current
		current = current.Next
	}
	return false // Chave não encontrada
}

// Print exibe a estrutura interna da tabela
func (ht *HashTable) Print() {
	fmt.Println("--- Estado da Tabela Hash ---")
	for i, node := range ht.Buckets {
		if node != nil {
			fmt.Printf("Bucket %d: ", i)
			for node != nil {
				fmt.Printf("[%s: %d] -> ", node.Key, node.Value)
				node = node.Next
			}
			fmt.Println("nil")
		}
	}
	fmt.Printf("Total de elementos: %d\n", ht.Size)
	fmt.Println("-----------------------------")
}
