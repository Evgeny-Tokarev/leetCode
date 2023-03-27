package main

type HeapMy struct {
}

func (h *HeapMy) Sort(arr []int) {
	h.CreateHeap(arr)
	for i := len(arr); i > 1; i-- {
		h.SwapBotton(arr, i)
	}
}
func (h *HeapMy) CreateHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		h.Heapify2(arr, i, len(arr))
	}
}
func (h *HeapMy) SwapBotton(arr []int, length int) {
	arr[0], arr[length-1] = arr[length-1], arr[0]
	h.Heapify2(arr, 0, length-1)
}

func (h *HeapMy) Heapify2(arr []int, root, length int) {
	var max = root
	var l, r = h.left(root), h.right(root)
	if l < length && arr[max] < arr[l] {
		max = l
	}
	if r < length && arr[max] < arr[r] {
		max = r
	}
	if max != root {
		arr[root], arr[max] = arr[max], arr[root]
		h.Heapify2(arr, max, length)
	}
}

func (*HeapMy) left(parent int) int {
	return (parent * 2) + 1
}
func (*HeapMy) right(parent int) int {
	return (parent * 2) + 2
}
