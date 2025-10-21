package datastructures

// SelectionSort ordena um slice de inteiros usando o algoritmo Selection Sort (in-place).
// A lógica é a mesma descrita na sua nota de aula.
func SelectionSort(data []int) {
	tamanho := len(data)
	// O loop externo vai até o penúltimo elemento.
	for i := 0; i < tamanho-1; i++ {

		// Assumimos que o menor elemento está na posição inicial 'i'.
		iMenor := i

		// O loop interno busca pelo menor elemento no restante do slice.
		for j := i + 1; j < tamanho; j++ {
			if data[j] < data[iMenor] {
				// Se encontrarmos um elemento menor, atualizamos o índice 'iMenor'.
				iMenor = j
			}
		}

		// Ao final do loop interno, trocamos o menor elemento encontrado
		// com o elemento na posição 'i'.
		// Esta é a forma idiomática de fazer a troca (swap) em Go.
		if i != iMenor {
			data[i], data[iMenor] = data[iMenor], data[i]
		}
	}
}

// BubbleSort ordena um slice de inteiros usando o algoritmo Bubble Sort (in-place).
func BubbleSort(data []int) {
	tamanho := len(data)
	for i := 0; i < tamanho-1; i++ {
		trocou := false
		for j := 0; j < tamanho-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				trocou = true
			}
		}
		if !trocou {
			break
		}
	}
}

// InsertionSort ordena um slice de inteiros usando o algoritmo Insertion Sort (in-place).
func InsertionSort(data []int) {
	tamanho := len(data)
	for i := 1; i < tamanho; i++ {
		valor := data[i]
		j := i
		for j > 0 && data[j-1] > valor {
			data[j] = data[j-1]
			j--
		}
		data[j] = valor
	}
}

// MergeSort ordena um slice de inteiros usando o algoritmo Merge Sort
func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	meio := len(data) / 2
	esquerda := MergeSort(data[:meio])
	direita := MergeSort(data[meio:])
	return merge(esquerda, direita)
}

// merge combina dois slices ordenados
func merge(esquerda, direita []int) []int {
	resultado := make([]int, len(esquerda)+len(direita))
	i, j := 0, 0
	for i < len(esquerda) && j < len(direita) {
		if esquerda[i] <= direita[j] {
			resultado[i+j] = esquerda[i]
			i++
		} else {
			resultado[i+j] = direita[j]
			j++
		}
	}
	for i < len(esquerda) {
		resultado[i+j] = esquerda[i]
		i++
	}
	for j < len(direita) {
		resultado[i+j] = direita[j]
		j++
	}
	return resultado
}

//todo: montar uma versão sem usar quickSortRecursive

// QuickSort é a função pública que inicia a ordenação
func QuickSort(data []int) {
	quickSortRecursive(data, 0, len(data)-1)
}

// quickSortRecursive é a função recursiva
func quickSortRecursive(data []int, low int, high int) {
	if low < high {
		pivotIndex := partition(data, low, high)
		quickSortRecursive(data, low, pivotIndex-1)
		quickSortRecursive(data, pivotIndex+1, high)
	}
}

// partition reorganiza o slice em torno de um pivô.
func partition(data []int, low int, high int) int {
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

// CountingSort
func CountingSort(data []int) []int {
	if len(data) == 0 {
		return []int{}
	}
	k := data[0]
	for _, value := range data {
		if value > k {
			k = value
		}
	}
	count := make([]int, k+1)
	for _, value := range data {
		count[value]++
	}
	index := 0
	for value, frequency := range count {
		for i := 0; i < frequency; i++ {
			data[index] = value
			index++
		}
	}
	return data
}
