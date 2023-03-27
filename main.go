package main

import "fmt"

func main() {
	array := []int{11, 15, 14, 7, 4, 9, 21}
	//BubbleSort(array)
	//MergeSort(array)
	//fmt.Println(array)
	//inputArray := []int{6, 5, 3, 7, 2, 8}
	var heap = new(HeapMy)
	heap.CreateHeap(array)
	fmt.Println(array)
	heap.Sort(array)
	fmt.Println(array)
}
