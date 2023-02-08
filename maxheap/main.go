package main

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) Push(val int) {
	h.data = append(h.data, val)
	h.heapifyUp()
}

func (h *MaxHeap) Pop() int {
	if len(h.data) == 0 {
		return 0
	}
	val := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown()
	return val
}

func (h *MaxHeap) heapifyUp() {
	idx := len(h.data) - 1
	for h.hasParent(idx) && h.parent(idx) < h.data[idx] {
		h.swap(h.parentIndex(idx), idx)
		idx = h.parentIndex(idx)
	}
}

func (h *MaxHeap) heapifyDown() {
	idx := 0
	for h.hasLeftChild(idx) {
		largerChildIndex := h.leftChildIndex(idx)
		if h.hasRightChild(idx) && h.rightChild(idx) > h.leftChild(idx) {
			largerChildIndex = h.rightChildIndex(idx)
		}
		if h.data[idx] > h.data[largerChildIndex] {
			break
		} else {
			h.swap(idx, largerChildIndex)
		}
		idx = largerChildIndex
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MaxHeap) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) leftChildIndex(i int) int {
	return i*2 + 1
}

func (h *MaxHeap) rightChildIndex(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) hasParent(i int) bool {
	return h.parentIndex(i) >= 0
}

func (h *MaxHeap) hasLeftChild(i int) bool {
	return h.leftChildIndex(i) < len(h.data)
}

func (h *MaxHeap) hasRightChild(i int) bool {
	return h.rightChildIndex(i) < len(h.data)
}

func (h *MaxHeap) parent(i int) int {
	return h.data[h.parentIndex(i)]
}

func (h *MaxHeap) leftChild(i int) int {
	return h.data[h.leftChildIndex(i)]
}

func (h *MaxHeap) rightChild(i int) int {
	return h.data[h.rightChildIndex(i)]
}
