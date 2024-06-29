package main

import (
	"fmt"
)

// partition разделяет массив на две части и возвращает индекс опорного элемента
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// quickSort реализует алгоритм быстрой сортировки
func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)

		// Рекурсивно сортируем подмассивы
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Исходный массив:", arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Отсортированный массив:", arr)
}
