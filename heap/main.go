package main

import "fmt"

type Maxheap struct {
	array []int
}

func (h *Maxheap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxheapfyUp(len(h.array) - 1)
}

// maxheapfyUp will heapify from bottom top
func (h *Maxheap) maxheapfyUp(i int) {
	for h.array[parent(i)] < h.array[i] {
		h.swap(parent(i), i)
		i = parent(i)
	}

}

//get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

//get the left child index
func left(i int) int {
	return 2*i + 1
}

// get the right child index
func right(i int) int {
	return 2*i + 2
}

// swap the keys in array
func (h *Maxheap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func (h *Maxheap) Extract() int {
	extraced := h.array[0]
	l := len(h.array) - 1
	if l == 0 {
		fmt.Println("The length of the array cannot be Zero")
		return -1
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.heapMaxifyDown(0)
	return extraced
}
func (h *Maxheap) heapMaxifyDown(i int) {

	lastindex := len(h.array) - 1
	r, l := right(i), left(i)
	childoCompare := 0
	// lopp while the index has at least one child
	for l <= lastindex {
		// when the index ha only one child
		if l == lastindex {
			childoCompare = l
		} else if h.array[l] > h.array[r] {
			// when the left child is larger
			childoCompare = l
		} else {
			// when the right child is larger
			childoCompare = r
		}
		if h.array[i] < h.array[childoCompare] {
			h.swap(i, childoCompare)
			i = childoCompare
			l, r = left(i), right(i)
		} else {
			return
		}
	}
}

func main() {
	m := &Maxheap{}
	as := []int{10, 20, 30, 12, 45, 6, 7, 46, 78}
	for _, v := range as {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
