package datastructures

// LinearSearch percorre o slice para encontrar o alvo.
// Retorna o índice se encontrar, ou -1 se não encontrar.
func LinearSearch(data []int, target int) int {
	for index, value := range data {
		if value == target {
			return index
		}
	}
	return -1
}

// BinarySearch busca o alvo em um slice ORDENADO.
// Retorna o índice se encontrar, ou -1 se não encontrar.
func BinarySearch(sortedData []int, target int) int {
	low := 0
	high := len(sortedData) - 1

	for low <= high {
		mid := low + (high-low)/2

		if sortedData[mid] == target {
			return mid
		} else if sortedData[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
