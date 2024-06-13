package main

import (
	"fmt"
	"math"
)

type Graph struct {
	Vertices int
	Edges    [][]int
}

func Newgraph(vertices int) *Graph {
	edges := make([][]int, vertices)
	for e := range edges {
		edges[e] = make([]int, vertices)
	}
	return &Graph{
		Edges:    edges,
		Vertices: vertices,
	}
}
func (g *Graph) AddEdge(u, v, weight int) {
	g.Edges[v][u] = weight
	g.Edges[u][v] = weight // for undirected graphs
}
func Dijkstra(graph *Graph, startVertex int) []int {
	dist := make([]int, graph.Vertices)
	visited := make([]bool, graph.Vertices)

	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[startVertex] = 0

	for count := 0; count < graph.Vertices-1; count++ {
		u := minDistance(dist, visited)
		visited[u] = true
		for v := 0; v < graph.Vertices; v++ {
			if !visited[v] && graph.Edges[u][v] != 0 && dist[u]+graph.Edges[u][v] < dist[v] {
				dist[v] = dist[u] + graph.Edges[u][v]
			}

		}
	}
	return dist
}
func minDistance(dist []int, visited []bool) int {
	min := math.MaxInt64
	minIndex := -1

	for v := 0; v < len(dist); v++ {
		if !visited[v] && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}
	return minIndex
}
func grapher() {
	graph := Newgraph(9)
	graph.AddEdge(0, 1, 4)
	graph.AddEdge(0, 7, 8)
	graph.AddEdge(1, 2, 8)
	graph.AddEdge(1, 7, 11)
	graph.AddEdge(2, 3, 7)
	graph.AddEdge(2, 8, 2)
	graph.AddEdge(2, 5, 4)
	graph.AddEdge(3, 4, 9)
	graph.AddEdge(3, 5, 14)
	graph.AddEdge(4, 5, 10)
	graph.AddEdge(5, 6, 2)
	graph.AddEdge(6, 7, 1)
	graph.AddEdge(6, 8, 6)
	graph.AddEdge(7, 8, 7)

	startVertex := 0
	dist := Dijkstra(graph, startVertex)

	fmt.Println("Shortest distances from vertex", startVertex, "to all other vertices:")
	for i, d := range dist {
		fmt.Printf("Vertex %d: Distance %d\n", i, d)
	}
}
