package main

import (
	"fmt"
)

// merge объединяет два упорядоченных среза в один
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	// Используем два указателя для отслеживания текущих позиций в срезах
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы, если таковые имеются
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// mergeSort выполняет сортировку слиянием
func mergeSort(arr []int) []int {
	// Базовый случай: если массив пустой или содержит один элемент, он уже отсортирован
	if len(arr) <= 1 {
		return arr
	}

	// Разбиваем массив на две половины
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	// Рекурсивно сортируем каждую половину
	left = mergeSort(left)
	right = mergeSort(right)

	// Объединяем отсортированные половины
	return merge(left, right)
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Исходный массив:", arr)

	sortedArr := mergeSort(arr)
	fmt.Println("Отсортированный массив:", sortedArr)
}
