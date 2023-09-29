package main

import "fmt"

type graph struct {
	numberofNodes int
	list          map[int][]int
}

func NewGraph() *graph {
	return &graph{
		list: make(map[int][]int),
	}
}
func (g *graph) addvertex(key int) {
	g.list[key] = []int{}
	g.numberofNodes++
}
func (g *graph) addEdges(vertex, edge int) {
	g.list[vertex] = append(g.list[vertex], edge)
	g.list[edge] = append(g.list[edge], vertex)
}
func (g *graph) showCOnnections() {
	for k, g := range g.list {
		fmt.Printf("%d has the following vertexes %v \n", k, g)
	}
}
func main() {
	g := NewGraph()
	g.addvertex(0)
	g.addvertex(1)
	g.addvertex(2)
	g.addvertex(3)
	g.addEdges(0, 1)
	g.addEdges(0, 3)
	g.addEdges(1, 3)
	g.addEdges(2, 3)
	g.showCOnnections()

}
