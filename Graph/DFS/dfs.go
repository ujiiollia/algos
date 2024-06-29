package main

import "fmt"

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

// DFS реализует поиск в глубину
func (g *Graph) DFS(start int, visited map[int]bool) {
	// Маркируем текущий узел как посещенный
	visited[start] = true
	fmt.Printf("%d ", start)

	// Рекурсивно посещаем всех соседей текущего узла
	for _, neighbor := range g.vertices[start] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
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

	// Инициализируем карту посещенных узлов
	visited := make(map[int]bool)

	// Выполняем DFS начиная с узла 2
	fmt.Println("Следующий порядок посещения вершин DFS:")
	g.DFS(2, visited)
}
