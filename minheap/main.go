package main

import "fmt"

type MinHeap struct {
	Store []int
}

func (m *MinHeap) Insert(index int) {
	m.Store = append(m.Store, index)
	m.MinHeapfyup(len(m.Store) - 1)
}
func (m *MinHeap) MinHeapfyup(index int) {
	for m.Store[parent(index)] > m.Store[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}

}
func (m *MinHeap) swap(x, y int) {
	m.Store[x], m.Store[y] = m.Store[y], m.Store[x]
}
func (m *MinHeap) Extract() (res int) {
	res = m.Store[0]
	l := len(m.Store) - 1
	if l == 0 {
		fmt.Println("The length of the array cannot be Zero")
		return -1
	}
	m.Store[0] = m.Store[l]
	m.Store = m.Store[:l]
	m.HeapFyDOwn(0)
	return res
}
func (m *MinHeap) HeapFyDOwn(i int) {

	lastindex := len(m.Store) - 1
	r, l := right(i), left(i)
	childToCompare := 0
	// lopp while the index has at least one child
	for l <= lastindex {
		// when the index has only one child
		if l == lastindex {
			childToCompare = l
		} else if m.Store[l] < m.Store[r] {
			// when the left child is smaller
			childToCompare = l
		} else {
			// when the right child is smaller
			childToCompare = r
		}
		if m.Store[i] > m.Store[childToCompare] {
			m.swap(i, childToCompare)
			i = childToCompare
			l, r = left(i), right(i)
		} else {
			return
		}
	}

}

func right(index int) int {
	return index*2 + 2
}
func left(index int) int {
	return index*2 + 1
}
func parent(index int) int {
	return (index - 1) / 2
}
func main() {
	m := &MinHeap{}
	as := []int{10, 14, 19, 26, 31, 44, 35, 33, 19, 42, 27}
	for _, v := range as {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
