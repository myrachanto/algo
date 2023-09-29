package main

import (
	"fmt"
)

type Linklist struct {
	head   *node
	length int
}
type node struct {
	name string
	next *node
	prev *node
}

func (l *Linklist) append(str string) {
	newnode := &node{name: str, next: nil, prev: nil}
	if l.head == nil {
		l.head = newnode
		l.length = 1

		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	newnode.prev = current
	current.next = newnode
	l.length++
}
func (l *Linklist) prepend(str string) {
	newnode := &node{name: str, next: l.head}
	l.head = newnode
	l.length++
}
func (l *Linklist) insert(index int, str string) {
	if index >= l.length {
		l.append(str)
		return
	}
	if index == 0 {
		l.prepend(str)
		return
	}
	leader := l.GetNextNode(index - 1)
	follower := leader.next
	newnode := &node{name: str, next: follower, prev: leader}
	leader.next = newnode
	follower.prev = newnode
	l.length++

}
func (l *Linklist) delete(index int) {
	if index >= l.length {
		return
	}
	leader := l.GetNextNode(index - 1)
	followerafter := leader.next.next
	leader.next = followerafter
	followerafter.prev = leader
	l.length--
}
func (l *Linklist) GetNextNode(index int) *node {
	counter := 0
	current := l.head
	for counter != index {
		current = current.next
		counter++
	}
	return current
}
func (l *Linklist) remove(str string) {
	current := l.head
	for current.next != nil {
		if current.next.next == nil {
			current.next = nil
			l.length--
			return
		}
		if current.next.name == str {
			current.next = current.next.next
		}
		current = current.next

	}
	l.length--
}
func (l *Linklist) List() {
	current := l.head
	for current.next != nil {
		fmt.Println(current.name)
		current = current.next
	}
	fmt.Println(current.name)

}

// func (l *Linklist) reverse() *Linklist {
// 	if l.head == nil {
// 		return l
// 	}
// 	reversedlist := Linklist{}

// }
func main() {
	l := Linklist{}
	l.append("first")
	l.append("second")
	l.append("third")
	l.prepend("heading")
	// l.remove("second")
	l.remove("third")
	l.delete(2)
	// fmt.Println(l.length)
	l.insert(1, "mic testing")
	l.List()
}
