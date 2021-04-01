package test

import (
	"golang_algorithm/algorithm/avlTree"
	"math/rand"
	"time"
)

type class int

func (c class) Cmp(a avlTree.Interface) int {
	if c == a.(class) {
		return 0
	}
	if c > a.(class) {
		return 1
	}
	return -1
}

func AvlTree() {
	tree := avlTree.New()
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 1000; i++ {
		key := r.Intn(100000)
		data := r.Intn(100000)
		tree.Add(class(key), data)
		/*
			if i % 20 == 0 {
				if tree.Check() == false {
					fmt.Print("error")
					return
				}
			}*/
	}
	tree.DisplayInorder()
	printSplitLine()
}
