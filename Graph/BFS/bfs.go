package main

import (
	"fmt"
)

// Graph представляет собой граф, используя список смежности
type Graph struct {
	vertices map[int][]int
}

// NewGraph создает новый экземпляр графа
func NewGraph() *Graph {
	return &Graph{vertices: make(map[int][]int)}
}

// AddEdge добавляет ребро в граф
func (g *Graph) AddEdge(v, w int) {
	g.vertices[v] = append(g.vertices[v], w)
}

// BFS реализует поиск в ширину
func (g *Graph) BFS(start int) {
	// Очередь для хранения вершин, ожидающих обработки
	var queue []int

	// Маркируем начальную вершину как посещенную и помещаем в очередь
	visited := make(map[int]bool)
	visited[start] = true
	queue = append(queue, start)

	for len(queue) > 0 {
		// Извлекаем вершину из очереди и обрабатываем её
		vertex := queue[0]
		queue = queue[1:] // удаляем первый элемент из очереди
		fmt.Printf("%d ", vertex)

		// Получаем всех соседей извлеченной вершины
		for _, neighbor := range g.vertices[vertex] {
			if !visited[neighbor] {
				// Если сосед ещё не был посещен, маркируем его как посещенного и добавляем в очередь
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	// Создаем граф и добавляем ребра
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	// Выполняем BFS начиная с узла 2
	fmt.Println("Следующий порядок посещения вершин BFS:")
	g.BFS(2)
}
