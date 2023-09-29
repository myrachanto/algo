package main

import "fmt"

type Lister[T comparable] struct {
	Store []T
}

func (l *Lister[T]) add(item T) {
	l.Store = append(l.Store, item)
}
func (l *Lister[T]) remove(item T) {
	updatedList := []T{}
	for _, v := range l.Store {
		if v != item {
			updatedList = append(updatedList, v)
		}
	}
	l.Store = updatedList
}
func (l *Lister[T]) clear() {
	l.Store = l.Store[:0]
}
func main() {
	res := Lister[int]{}
	// res2 := Lister[string]{}
	res.add(1)
	res.add(2)
	res.add(3)
	res.add(4)
	fmt.Println("--------------", res)
	res.remove(3)
	fmt.Println("--------------", res)
	res.clear()
	fmt.Println("--------------", res)
}
