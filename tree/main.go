package main

import (
	"fmt"
)

type tree struct {
	root *node
}
type node struct {
	value int
	left  *node
	right *node
}

func (t *tree) insert(value int) {
	newNode := &node{value: value}
	if t.root == nil {
		t.root = newNode
		return
	}
	current := t.root
	for {
		if value > current.value {
			if current.right == nil {
				current.right = newNode
				break
			}
			current = current.right
		} else {
			if current.left == nil {
				current.left = newNode
				break
			}
			current = current.left
		}
	}
}
func (t *tree) lookup(value int) bool {
	current := t.root
	var res bool
	for current != nil {
		if value == current.value {
			res = true
			break
		}
		if value > current.value {
			current = current.right
		} else {
			current = current.left
		}
	}
	return res
}
func (t *tree) delete(value int) bool {
	current := t.root
	parent := &node{}
	var res bool
	for current != nil {
		if value == current.value {
			if current.right == nil {
				parent.left = current.left
				current = current.left
				res = true
			} else if current.left.value < current.right.left.value {
				parent.left = current.right.left
				current = current.right.left
				res = true
			} else if current.right.value > current.left.value {
				parent.left = current.right
				current = current.right
				res = true
			}
		}
		fmt.Println("=++++++++++++++", current)
		if current == nil {
			break
		}
		parent = current
		if value > current.value {
			current = current.right
		} else {
			current = current.left
		}
	}
	return res
}
func main() {
	t := &tree{}
	t.insert(10)
	fmt.Println(t.root)
	t.insert(11)
	// fmt.Println(t.root.right)
	t.insert(8)
	t.insert(9)
	t.insert(5)
	fmt.Println(t.root.left)
	fmt.Println(t.lookup(5))
	fmt.Println("deleted ", t.delete(5))
}
