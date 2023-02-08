package main

import "fmt"

type Queue interface {
	Push(i int)
	Dequeue() int
}
type queue struct {
	data []int
}

func (q *queue) Push(i int) {
	q.data = append(q.data, i)
}
func (q *queue) Dequeue() int {
	firstIndex := q.data[0]
	q.data = q.data[1:]
	return firstIndex
}

func main() {
	ques := &queue{}
	ques.Push(100)
	ques.Push(200)
	ques.Push(300)
	fmt.Println(ques)
	r := ques.Dequeue()
	fmt.Println(ques, r)
}
