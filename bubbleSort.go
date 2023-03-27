package main

func BubbleSort(ar []int) {
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
