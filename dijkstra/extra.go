package main

import (
	"container/heap"
	"fmt"
)

// Node represents a node in the graph
type Node struct {
	name     string
	distance int
	index    int
}

// PriorityQueue represents a priority queue of nodes
type PriorityQueue []*Node

// Len returns the length of the priority queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less compares the distance of two nodes
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

// Swap swaps two nodes in the priority queue
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push pushes a node onto the priority queue
func (pq *PriorityQueue) Push(n interface{}) {
	node := n.(*Node)
	node.index = len(*pq)
	*pq = append(*pq, node)
}

// Pop removes and returns the node with the smallest distance
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.index = -1
	*pq = old[0 : n-1]
	return node
}

// ShortestPath returns the shortest path between two nodes
func ShortestPath(graph map[string]map[string]int, start, end string) (int, []string) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	distances := make(map[string]int)
	previous := make(map[string]string)
	visited := make(map[string]bool)

	for node := range graph {
		distances[node] = 1<<31 - 1
		previous[node] = ""
		visited[node] = false
		heap.Push(&pq, &Node{node, distances[node], 0})
	}

	distances[start] = 0
	// heap.Fix(&pq, &Node{start, 0, 0})
	heap.Fix(&pq, pq.indexOf(start))

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Node)

		if visited[current.name] {
			continue
		}

		visited[current.name] = true

		for neighbor := range graph[current.name] {
			distance := graph[current.name][neighbor]
			if distances[neighbor] > distances[current.name]+distance {
				distances[neighbor] = distances[current.name] + distance
				previous[neighbor] = current.name
				heap.Fix(&pq, &Node{neighbor, distances[neighbor], 0})
			}
		}
	}
	path := make([]string, 0)
	for node := end; node != start; node = previous[node] {
		path = append([]string{node}, path...)
		if node == start {
			break
		}
	}
	path = append([]string{start}, path...)
	return distances[end], path
}
func (pq PriorityQueue) indexOf(name string) int {
    for i, node := range pq {
        if node.name == name {
            return i
        }
    }
    return -1
}

func main() {
	graph := map[string]map[string]int{
		"A": map[string]int{"B": 5, "C": 1},
		"B": map[string]int{"A": 5, "C": 2, "D": 1},
		"C": map[string]int{"A": 1, "B": 2, "D": 4},
		"D": map[string]int{"B": 1, "C": 4},
	}
	distance, path := ShortestPath(graph, "A", "D")
	fmt.Println("Shortest distance:", distance)
	fmt.Println("Shortest path:", path)
}
