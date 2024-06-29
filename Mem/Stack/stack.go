package main

import "fmt"

// Stack представляет собой структуру стека.
type Stack struct {
	elements []int
}

// Push добавляет элемент в стек.
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop удаляет и возвращает верхний элемент из стека. Возвращает 0, если стек пуст.
func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		fmt.Println("Стек пуст")
		return 0
	}
	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return value
}

func main() {
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println(stack.Pop()) // 30
	fmt.Println(stack.Pop()) // 20
	fmt.Println(stack.Pop()) // 10
	fmt.Println(stack.Pop()) // Стек пуст, 0
}
