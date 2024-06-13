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

func (t *tree) treecreator() {
	t.insert(9)
	t.insert(4)
	t.insert(6)
	t.insert(20)
	t.insert(170)
	t.insert(15)
	t.insert(1)
}

func main() {
	t := &tree{}
	t.treecreator()
	res := t.DFSInorder()
	fmt.Println("----->", res)
	res = t.DFSPreOrder()
	fmt.Println("----->", res)
	res = t.DFSPostOrder()
	fmt.Println("----->", res)
	// grapher()
	bellmanner()


}
func (t *tree) DFSInorder() []int {
	ls := []int{}
	if t.root == nil {
		return ls
	}
	var dfs func(n *node)
	dfs = func(n *node) {
		if n == nil {
			return
		}
		dfs(n.left)
		ls = append(ls, n.value)
		dfs(n.right)
	}
	dfs(t.root)
	return ls
}
func (t *tree) DFSPreOrder() []int {
	ls := []int{}
	if t.root == nil {
		return ls
	}
	var dfs func(n *node)
	dfs = func(n *node) {
		if n == nil {
			return
		}
		ls = append(ls, n.value)
		dfs(n.left)
		dfs(n.right)
	}
	dfs(t.root)
	return ls
}
func (t *tree) DFSPostOrder() []int {
	ls := []int{}
	if t.root == nil {
		return ls
	}
	var dfs func(n *node)
	dfs = func(n *node) {
		if n == nil {
			return
		}
		dfs(n.left)
		dfs(n.right)
		ls = append(ls, n.value)
	}
	dfs(t.root)
	return ls
}
func (t *tree) BreadthFirstSearch() []int {
	current := t.root
	results := []int{} 
	queue := []*node{}
	queue = append(queue, current)
	for len(queue) > 0 {
		current, queue = dequeue(queue)
		results = append(results, current.value)
		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
	return results
}
func dequeue(ls []*node) (*node, []*node) {
	noded := ls[0]
	results := []*node{}
	results = append(results, ls[1:]...)
	return noded, results
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
