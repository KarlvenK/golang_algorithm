package main

import (
	"algorithm/algorithm"
	"fmt"
)

type intSlice []int

func (data intSlice) Len() int {
	return len(data)
}

func (data intSlice) Less(i, j int) bool {
	return data[i] < data[j]
}

func (data intSlice) Swap(i, j int) {
	data[i], data[j] = data[j], data[i]
}

func main() {
	arr := []int{3, 4, 1, 2, 5, 8, 5}
	fmt.Println(arr)
	algorithm.Sort(intSlice(arr))
	fmt.Print(arr)
}
