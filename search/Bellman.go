package main

import (
	"fmt"
	"math"
)

type Edge struct {
	source, destination, weight int
}
type BGraph struct {
	vertices, edges int
	edgeList        []Edge
}

func NewBGrapgh(vertices, edges int) *BGraph {
	return &BGraph{
		vertices: vertices,
		edges:    edges,
		edgeList: make([]Edge, edges),
	}
}
func (b *BGraph) AddEdge(source, destination, weight int) {
	b.edgeList = append(b.edgeList, Edge{source: source, destination: destination, weight: weight})
}
func Bellmanford(graph *BGraph, startVertex int) ([]int, []int) {
	// Initialize distance and predecessor arrays
	dist := make([]int, graph.vertices)
	prev := make([]int, graph.vertices)
	for i := range dist {
		dist[i] = math.MaxInt64
		prev[i] = -1
	}

	dist[startVertex] = 0
	// Relax edges repeatedly
	for i := 1; i <= graph.vertices-1; i++ {
		for _, edge := range graph.edgeList {
			u := edge.source
			v := edge.destination
			w := edge.weight
			if dist[u] != math.MaxInt64 && dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				prev[v] = u
			}
		}
	}
	// checking for negative weight cycles
	// negativeCycles := false
	for _, edge := range graph.edgeList {
		u := edge.source
		v := edge.destination
		w := edge.weight
		if dist[u] != math.MaxInt64 && dist[u]+w < dist[v] {
			// negativeCycles = true
			break
		}

	}
	return dist, prev
}
func bellmanner() {
	vertices := 5
	edges := 8
	graph := NewBGrapgh(vertices, edges)

	// Add edges to the graph
	graph.AddEdge(0, 1, -1)
	graph.AddEdge(0, 2, 4)
	graph.AddEdge(1, 2, 3)
	graph.AddEdge(1, 3, 2)
	graph.AddEdge(1, 4, 2)
	graph.AddEdge(3, 2, 5)
	graph.AddEdge(3, 1, 1)
	graph.AddEdge(4, 3, -3)

	startVertex := 0
	dist, prev := Bellmanford(graph, startVertex)

	if dist != nil {
		fmt.Println("Shortest distances from vertex", startVertex, "to all other vertices:")
		for i, d := range dist {
			fmt.Printf("Vertex %d: Distance %d, Previous Vertex: %d\n", i, d, prev[i])
		}
	} else {
		fmt.Println("Negative-weight cycle detected. The algorithm cannot proceed.")
	}
}
