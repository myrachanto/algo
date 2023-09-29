package main

import "fmt"

type shape struct {
	X, Y float64
}
type LinkedList struct {
	Length int
	Tail   *LinkedList
}

func (l *LinkedList) sum() int {
	if l == nil {
		return 0
	}
	return l.Length + l.Tail.sum()
}

func (s *shape) Area() float64 {
	return s.X * s.Y
}
func (s shape) Distance() float64 {
	return s.X + s.Y
}
func main() {
	ptr := &shape{5, 5}
	fmt.Println(ptr.Area())
	fmt.Println(ptr.Distance())
	nonptr := shape{4, 4}
	fmt.Println(nonptr.Area())
	fmt.Println(nonptr.Distance())

	l := &LinkedList{5, &LinkedList{10, &LinkedList{20, &LinkedList{}}}}
	fmt.Println(l.sum())
}
