package main

import "fmt"

func bubbleSort(ar []int) {
	var tmp int
	for i := 0; i < len(ar)-1; i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar)[i] > (ar)[j] {
				tmp = (ar)[i]
				(ar)[i] = (ar)[j]
				(ar)[j] = tmp
			}
		}
	}
}
func merge(arr []int, leftStart, leftEnd, rightStart, rightEnd int) {
	leftSize := leftEnd - leftStart + 1
	rightSize := rightEnd - rightStart + 1
	temp := make([]int, leftSize+rightSize)

	i, j, k := leftStart, rightStart, 0
	for i <= leftEnd && j <= rightEnd {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= leftEnd {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= rightEnd {
		temp[k] = arr[j]
		j++
		k++
	}
	for i := 0; i < len(temp); i++ {
		arr[leftStart+i] = temp[i]
	}
}
func mergeSort(arr []int) {
	start, end := 0, len(arr)-1
	if start < end {
		mid := (start + end) / 2
		mergeSort(arr[start : mid+1])
		mergeSort(arr[mid+1 : end+1])
		merge(arr, start, mid, mid+1, end)
	}
}
func main() {
	array := []int{1, 3, 2, 7, 4, 9, 21}
	//bubbleSort(array)
	mergeSort(array)
	fmt.Println(array)
}
