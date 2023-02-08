package main

import (
	"fmt"
	"math"
	"os"
	"sync"
)

// create a weighted graph ing golang
// structures
type (
	Node struct {
		Name    string
		Value   int
		Through *Node
	}
	Edge struct {
		Node   *Node
		Weight int
	}
	WeightedGraph struct {
		Nodes []*Node
		Edges map[string][]*Edge
		Mutex sync.RWMutex
	}
	Heap struct {
		Elements []*Node
		Mutex    sync.RWMutex
	}
)

func NewGraph() *WeightedGraph {
	return &WeightedGraph{
		Edges: make(map[string][]*Edge),
	}
}

// GetNode
func (g *WeightedGraph) GetNode(name string) (node *Node) {
	g.Mutex.RLock()
	defer g.Mutex.RUnlock()
	for _, n := range g.Nodes {
		if n.Name == name {
			node = n
		}
	}
	return
}

// AddNode to the graph
func (g *WeightedGraph) AddNode(n *Node) {
	g.Mutex.RLock()
	defer g.Mutex.RUnlock()
	g.Nodes = append(g.Nodes, n)
}
func AddNodes(graph *WeightedGraph, names ...string) (nodes map[string]*Node) {
	nodes = make(map[string]*Node)
	for _, name := range names {
		n := &Node{name, math.MaxInt, nil}
		graph.AddNode(n)
		nodes[name] = n
	}
	return
}
func (g *WeightedGraph) AddEdge(n1, n2 *Node, weight int) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	g.Edges[n1.Name] = append(g.Edges[n1.Name], &Edge{n2, weight})
	g.Edges[n2.Name] = append(g.Edges[n2.Name], &Edge{n1, weight})
}

// min heap
func (h *Heap) Size() int {
	h.Mutex.RLock()
	defer h.Mutex.RUnlock()
	return len(h.Elements)
}

// Push an element into the heap and heapify
func (h *Heap) Push(element *Node) {
	h.Mutex.RLock()
	defer h.Mutex.RUnlock()
	h.Elements = append(h.Elements, element)
	// addedElementIndex := len(h.Elements) - 1
	i := len(h.Elements) - 1
	for h.Elements[i].Value < h.Elements[parent(i)].Value {
		h.swap(i, parent(i))
	}

}
func (h *Heap) Pop() (i *Node) {
	h.Mutex.RLock()
	defer h.Mutex.RUnlock()
	i = h.Elements[0]
	h.Elements[0] = h.Elements[len(h.Elements)-1]
	h.Elements = h.Elements[:len(h.Elements)-1]
	h.Rearrage(0)
	return

}
func buildGraph() *WeightedGraph {
	graph := NewGraph()
	nodes := AddNodes(graph,
		"London",
		"Paris",
		"Amsterdam",
		"Luxembourg",
		"Zurich",
		"Rome",
		"Berlin",
		"Vienna",
		"Warsaw",
		"Istanbul",
	)
	graph.AddEdge(nodes["London"], nodes["Paris"], 80)
	graph.AddEdge(nodes["London"], nodes["Luxembourg"], 75)
	graph.AddEdge(nodes["London"], nodes["Amsterdam"], 75)
	graph.AddEdge(nodes["Paris"], nodes["Luxembourg"], 60)
	graph.AddEdge(nodes["Paris"], nodes["Rome"], 125)
	graph.AddEdge(nodes["Luxembourg"], nodes["Berlin"], 90)
	graph.AddEdge(nodes["Luxembourg"], nodes["Zurich"], 60)
	graph.AddEdge(nodes["Luxembourg"], nodes["Amsterdam"], 55)
	graph.AddEdge(nodes["Zurich"], nodes["Vienna"], 80)
	graph.AddEdge(nodes["Zurich"], nodes["Rome"], 90)
	graph.AddEdge(nodes["Zurich"], nodes["Berlin"], 85)
	graph.AddEdge(nodes["Berlin"], nodes["Amsterdam"], 85)
	graph.AddEdge(nodes["Berlin"], nodes["Vienna"], 75)
	graph.AddEdge(nodes["Vienna"], nodes["Rome"], 100)
	graph.AddEdge(nodes["Vienna"], nodes["Istanbul"], 130)
	graph.AddEdge(nodes["Warsaw"], nodes["Berlin"], 80)
	graph.AddEdge(nodes["Warsaw"], nodes["Istanbul"], 180)
	graph.AddEdge(nodes["Rome"], nodes["Istanbul"], 155)
	return graph
}

// Dijkstras algorithym
func Dijkstras(graph *WeightedGraph, city string) {
	visited := make(map[string]bool)
	heap := &Heap{}

	startNode := graph.GetNode(city)
	startNode.Value = 0
	heap.Push(startNode)

	for heap.Size() > 0 {
		current := heap.Pop()
		visited[current.Name] = true
		edges := graph.Edges[current.Name]
		for _, edge := range edges {
			if !visited[edge.Node.Name] {
				heap.Push(edge.Node)
				if current.Value+edge.Weight < edge.Node.Value {
					edge.Node.Value = current.Value + edge.Weight
					edge.Node.Through = current
				}
			}
		}
	}

}

// Rearrage the heap
func (h *Heap) Rearrage(i int) {
	smallest := i
	l, r, s := left(i), right(i), len(h.Elements)
	if l < s && h.Elements[l].Value < h.Elements[smallest].Value {
		smallest = l
	}
	if r < s && h.Elements[r].Value < h.Elements[smallest].Value {
		smallest = r
	}
	if smallest != i {
		h.swap(i, smallest)
		h.Rearrage(smallest)
	}
}
func (h *Heap) swap(i, j int) {
	h.Elements[i], h.Elements[j] = h.Elements[j], h.Elements[i]
}
func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return i*2 + 1
}
func right(i int) int {
	return i*2 + 2
}

func main() {
	// build and run Dijkstra algorithm on the graph
	graph := buildGraph()
	city := os.Args[1]
	Dijkstras(graph, city)
	for _, node := range graph.Nodes {
		fmt.Printf("Shotrest time from %s to %s is %d \n", city, node.Name, node.Value)
		for n := node; n.Through != nil; n = n.Through {
			fmt.Print(n.Name, "<-")

		}
		fmt.Println(city)
		fmt.Println()
	}

}
// Shotrest time from Istanbul to London is 345
// London<-Luxembourg<-Zurich<-Vienna<-Istanbul

// Shotrest time from Istanbul to Paris is 280
// Paris<-Rome<-Istanbul

// Shotrest time from Istanbul to Amsterdam is 290
// Amsterdam<-Berlin<-Vienna<-Istanbul

// Shotrest time from Istanbul to Luxembourg is 270
// Luxembourg<-Zurich<-Vienna<-Istanbul

// Shotrest time from Istanbul to Zurich is 210
// Zurich<-Vienna<-Istanbul

// Shotrest time from Istanbul to Rome is 155
// Rome<-Istanbul

// Shotrest time from Istanbul to Berlin is 205
// Berlin<-Vienna<-Istanbul

// Shotrest time from Istanbul to Vienna is 130
// Vienna<-Istanbul

// Shotrest time from Istanbul to Warsaw is 180
// Warsaw<-Istanbul

// Shotrest time from Istanbul to Istanbul is 0
// Istanbul