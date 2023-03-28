package main

import (
	"math/rand"
	"time"
)

func getRand(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

func CreateRandomArr(cols, rows, rangeLimit int) ([][]int, int, int) {
	var num = getRand(rangeLimit)
	ok := false
	var maxLine int
	max := num
	arr := make([][]int, rows)
	arrMap := make(map[int]interface{})
	for i := 0; i < rows; i++ {
		arr[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			for ok {
				num = getRand(rangeLimit)
				_, ok = arrMap[num]
			}
			ok = true
			if num > max {
				max = num
				maxLine = i
			}
			arrMap[num] = 0
			arr[i][j] = num
		}
	}
	return arr, maxLine, max
}
