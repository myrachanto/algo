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

type queuelinked struct {
	First  *node
	Last   *node
	length int
}
type node struct {
	yeild int
	next  *node
}

func (q *queuelinked) Enqueue(yeild int) {
	newnode := &node{yeild: yeild}
	if q.length == 0 {
		q.First = newnode
		q.Last = newnode
		q.length = 1
	}
	// holder := q.Last
	q.Last.next = newnode
	q.Last = newnode
	q.length++
}
func (q *queuelinked) Peek() *node {
	return q.First
}
func (q *queuelinked) Dequeue() {
	if q.First == nil {
		return
	}
	if q.length == 1 {
		q.Last = nil
		q.First = nil
		q.length = 0
		return
	}
	// holding := q.Top
	q.First = q.First.next
	q.length--

}
