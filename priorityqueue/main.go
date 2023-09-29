package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Value    string
	Priority int
	Index    int
}
type PriolityList []*Item

func (pl PriolityList) Len() int {
	return len(pl)
}
func (pl PriolityList) Less(i, j int) bool {
	return pl[i].Priority < pl[j].Priority
}
func (pl PriolityList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
	pl[i].Index = i
	pl[j].Index = j
}
func (pl *PriolityList) Push(i interface{}) {
	item := i.(*Item)
	item.Index = pl.Len()
	*pl = append(*pl, item)
}
func (pl *PriolityList) Pop() interface{} {
	old := *pl
	n := len(old)
	item := old[n-1]
	*pl = old[0 : n-1]
	return item

}
func main() {
	pq := make(PriolityList, 0)

	// Push items with priorities into the priority queue.
	items := []*Item{
		{Value: "foo", Priority: 3},
		{Value: "bar", Priority: 1},
		{Value: "baz", Priority: 2},
	}

	for _, item := range items {
		heap.Push(&pq, item)
	}

	// Pop items from the priority queue in order of priority.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Value: %s, Priority: %d\n", item.Value, item.Priority)
	}
}
