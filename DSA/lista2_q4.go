package main

import (
	"DSA/datastructures"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// --- Funções Auxiliares ---

type sortFunction func([]int)
type sortFunctionMerge func([]int) []int

func benchmarkSort(fn sortFunction, data []int, name string) time.Duration {
	dataCopy := make([]int, len(data))
	copy(dataCopy, data)
	start := time.Now()
	fn(dataCopy)
	duration := time.Since(start)
	if !sort.IntsAreSorted(dataCopy) {
		fmt.Printf("ERRO: %s não ordenou corretamente!\n", name)
	}
	return duration
}

func benchmarkSortMerge(fn sortFunctionMerge, data []int, name string) time.Duration {
	dataCopy := make([]int, len(data))
	copy(dataCopy, data)
	start := time.Now()
	sortedData := fn(dataCopy)
	duration := time.Since(start)
	if !sort.IntsAreSorted(sortedData) {
		fmt.Printf("ERRO: %s não ordenou corretamente!\n", name)
	}
	return duration
}

// --- QuickSort Lomuto (Sem Randomização - para o item D) ---
func quickSortLomutoHelper(data []int, low int, high int) {
	if low < high {
		pivotIndex := partitionLomuto(data, low, high)
		quickSortLomutoHelper(data, low, pivotIndex-1)
		quickSortLomutoHelper(data, pivotIndex+1, high)
	}
}
func partitionLomuto(data []int, low int, high int) int {
	pivot := data[high]
	pIndex := low
	for i := low; i < high; i++ {
		if data[i] <= pivot {
			data[i], data[pIndex] = data[pIndex], data[i]
			pIndex++
		}
	}
	data[pIndex], data[high] = data[high], data[pIndex]
	return pIndex
}
func QuickSortLomuto(data []int) {
	quickSortLomutoHelper(data, 0, len(data)-1)
}

// --- Função Principal ---

func main() {
	fmt.Println("Executando experimentos de ordenação...")

	// --- Vetores ---
	vetorPequenoAleatorio := []int{8, 3, 5, 1, 9, 4, 7, 2, 6, 0} // n=10

	nGrande := 10000
	kFavoravel := nGrande
	kDesfavoravel := nGrande * 100 // k = 1.000.000

	vetorGrandeAleatorio := make([]int, nGrande)
	vetorGrandeCrescente := make([]int, nGrande)
	vetorGrandeDecrescente := make([]int, nGrande)
	// Novo vetor para CS Desfavorável, com n=10000 e k=1M
	vetorGrandeCSDesfavoravel := make([]int, nGrande)

	for i := 0; i < nGrande; i++ {
		vetorGrandeAleatorio[i] = rand.Intn(kFavoravel) // k=n
		vetorGrandeCrescente[i] = i
		vetorGrandeDecrescente[i] = nGrande - 1 - i
		vetorGrandeCSDesfavoravel[i] = rand.Intn(kDesfavoravel) // k=n*100
	}

	fmt.Printf("Tamanho Pequeno (n=%d), Tamanho Grande (n=%d)\n\n", len(vetorPequenoAleatorio), nGrande)

	// --- A. Vetores Pequenos ---
	fmt.Println("--- A. Vetores Pequenos ---")
	tempoIS_A := benchmarkSort(datastructures.InsertionSort, vetorPequenoAleatorio, "IS (A)")
	tempoMS_A := benchmarkSortMerge(datastructures.MergeSort, vetorPequenoAleatorio, "MS (A)")
	tempoCS_A := benchmarkSortMerge(datastructures.CountingSort, vetorPequenoAleatorio, "CS (A)")
	fmt.Printf("  Pequeno Aleatório -> IS: %v, MS: %v, CS: %v\n", tempoIS_A, tempoMS_A, tempoCS_A)

	// --- B. Vetores Grandes ---
	fmt.Println("\n--- B. Vetores Grandes ---")
	// B.1: Aleatório
	tempoIS_B_Rand := benchmarkSort(datastructures.InsertionSort, vetorGrandeAleatorio, "IS (B-Rand)")
	tempoMS_B_Rand := benchmarkSortMerge(datastructures.MergeSort, vetorGrandeAleatorio, "MS (B-Rand)")
	tempoCS_B_Rand := benchmarkSortMerge(datastructures.CountingSort, vetorGrandeAleatorio, "CS (B-Rand Fav)")
	fmt.Printf("  Grande Aleatório (k=n) -> IS: %v, MS: %v, CS: %v\n", tempoIS_B_Rand, tempoMS_B_Rand, tempoCS_B_Rand)
	// B.2: Crescente
	tempoIS_B_Cresc := benchmarkSort(datastructures.InsertionSort, vetorGrandeCrescente, "IS (B-Cresc)")
	tempoMS_B_Cresc := benchmarkSortMerge(datastructures.MergeSort, vetorGrandeCrescente, "MS (B-Cresc)")
	tempoCS_B_Cresc := benchmarkSortMerge(datastructures.CountingSort, vetorGrandeCrescente, "CS (B-Cresc)")
	fmt.Printf("  Grande Crescente      -> IS: %v, MS: %v, CS: %v\n", tempoIS_B_Cresc, tempoMS_B_Cresc, tempoCS_B_Cresc)
	// B.3: Decrescente
	tempoIS_B_Decresc := benchmarkSort(datastructures.InsertionSort, vetorGrandeDecrescente, "IS (B-Decresc)")
	tempoMS_B_Decresc := benchmarkSortMerge(datastructures.MergeSort, vetorGrandeDecrescente, "MS (B-Decresc)")
	tempoCS_B_Decresc := benchmarkSortMerge(datastructures.CountingSort, vetorGrandeDecrescente, "CS (B-Decresc)")
	fmt.Printf("  Grande Decrescente    -> IS: %v, MS: %v, CS: %v\n", tempoIS_B_Decresc, tempoMS_B_Decresc, tempoCS_B_Decresc)

	// --- C. Consistência do MergeSort ---
	fmt.Println("\n--- C. Consistência do MergeSort ---")
	fmt.Printf("  (Resultados já mostrados no item B)\n")
	fmt.Printf("  Aleatório: %v, Crescente: %v, Decrescente: %v\n", tempoMS_B_Rand, tempoMS_B_Cresc, tempoMS_B_Decresc)

	// --- D. Pior Caso QuickSort vs Randomização ---
	fmt.Println("\n--- D. Pior Caso QuickSort vs Randomização ---")
	tempoQSLomuto_D_Cresc := benchmarkSort(QuickSortLomuto, vetorGrandeCrescente, "QSLomuto (D-Cresc)")
	tempoQS_D_Cresc := benchmarkSort(datastructures.QuickSort, vetorGrandeCrescente, "QS (D-Cresc)")
	fmt.Printf("  Vetor Crescente -> QS (Pivô Fixo): %v, QS (Random): %v\n", tempoQSLomuto_D_Cresc, tempoQS_D_Cresc)
	tempoQSLomuto_D_Decresc := benchmarkSort(QuickSortLomuto, vetorGrandeDecrescente, "QSLomuto (D-Decresc)")
	tempoQS_D_Decresc := benchmarkSort(datastructures.QuickSort, vetorGrandeDecrescente, "QS (D-Decresc)")
	fmt.Printf("  Vetor Decrescente -> QS (Pivô Fixo): %v, QS (Random): %v\n", tempoQSLomuto_D_Decresc, tempoQS_D_Decresc)
	tempoQSLomuto_D_Rand := benchmarkSort(QuickSortLomuto, vetorGrandeAleatorio, "QSLomuto (D-Rand)")
	tempoQS_D_Rand := benchmarkSort(datastructures.QuickSort, vetorGrandeAleatorio, "QS (D-Rand)")
	fmt.Printf("  Vetor Aleatório -> QS (Pivô Fixo): %v, QS (Random): %v\n", tempoQSLomuto_D_Rand, tempoQS_D_Rand)

	// --- E. Desempenho do CountingSort ---
	fmt.Println("\n--- E. Desempenho do CountingSort ---")
	fmt.Printf("  Caso Favorável (n=%d, k=n=%d) -> CS: %v (Resultado do item B.1)\n", nGrande, kFavoravel, tempoCS_B_Rand)
	tempoCS_E_Desfav := benchmarkSortMerge(datastructures.CountingSort, vetorGrandeCSDesfavoravel, "CS (E-Desfav)")
	fmt.Printf("  Caso Desfavorável (n=%d, k=n*100=%d) -> CS: %v\n", nGrande, kDesfavoravel, tempoCS_E_Desfav)

	fmt.Println("\nExperimentos concluídos.")
}
