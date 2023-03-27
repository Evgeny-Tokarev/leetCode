package main

import "fmt"

func main() {
	array := []int{11, 15, 14, 7, 4, 9, 21}
	//BubbleSort(array)
	//MergeSort(array)
	//fmt.Println(array)
	inputArray := []int{11, 15, 14, 7, 4, 9, 21}
	var heap2 = new(HeapMy)
	heap2.Sort(array)
	fmt.Println(array)
	var heap = new(Heap)
	heap.HeapSort(inputArray)
	fmt.Println(inputArray)
}
