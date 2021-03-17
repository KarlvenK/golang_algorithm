package main

import (
	"algorithm/algorithm/list"
	"algorithm/algorithm/sort"
	"fmt"
	"math/rand"
	"time"
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

func testSort() {
	arr := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	t := 10000

	for i := 0; i < t; i++ {
		n := r.Intn(100000000)
		arr = append(arr, n)
	}

	fmt.Println(arr[:30])
	sort.Sort(intSlice(arr)) // we do sort here

	if checkSort(arr) {
		fmt.Println("ok")
		fmt.Println(arr[:30])
	} else {
		fmt.Println("no")
	}
	printSplitLine()
}

func checkSort(arr intSlice) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func testList() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	printSplitLine()
}

func main() {
	testSort()
	testList()
}
func printSplitLine() {
	fmt.Println("====================================================")
}
