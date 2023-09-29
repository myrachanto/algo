package main

import "fmt"

type stack struct {
	data []int
}

func (s *stack) Push(i int) {
	s.data = append(s.data, i)
}
func (s *stack) Pop() int {
	lastindex := len(s.data) - 1
	s.data = s.data[:lastindex]
	return lastindex
}
func (s *stack) Peek() int {
	return s.data[len(s.data)-1]
}
func main() {
	mydata := &stack{}
	mydata.Push(100)
	mydata.Push(200)
	mydata.Push(300)
	fmt.Println(mydata)
	mydata.Pop()
	fmt.Println(mydata)

}

type stacklinked struct {
	Top    *node
	Bottom *node
	length int
}
type node struct {
	yeild int
	next  *node
}

func (sl *stacklinked) Push(yeild int) {
	newnode := &node{yeild: yeild}
	if sl.length == 0 {
		sl.Top = newnode
		sl.Bottom = newnode
		sl.length = 1
	}
	holding := sl.Top
	sl.Top = newnode
	sl.Top.next = holding
	sl.length++
}
func (sl *stacklinked) Peek() *node {
	return sl.Top
}
func (sl *stacklinked) Pop() {
	if sl.Top == nil {
		return
	}
	if sl.length == 1 {
		sl.Top = nil
		sl.Bottom = nil
		sl.length = 0
		return
	}
	// holding := sl.Top
	sl.Top = sl.Top.next
	sl.length--

}
