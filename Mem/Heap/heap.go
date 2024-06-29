package main

import "fmt"

// Queue представляет собой структуру очереди.
type Queue struct {
	elements []int
}

// Enqueue добавляет элемент в конец очереди.
func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

// Dequeue удаляет и возвращает первый элемент из очереди. Возвращает 0, если очередь пуста.
func (q *Queue) Dequeue() int {
	if len(q.elements) == 0 {
		fmt.Println("Очередь пуста")
		return 0
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value
}

func main() {
	queue := &Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println(queue.Dequeue()) // 10
	fmt.Println(queue.Dequeue()) // 20
	fmt.Println(queue.Dequeue()) // 30
	fmt.Println(queue.Dequeue()) // Очередь пуста, 0
}
