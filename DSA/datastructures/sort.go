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
