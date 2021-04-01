package test

import (
	"golang_algorithm/algorithm/avlTree"
	"math/rand"
	"time"
)

func AvlTree() {
	tree := avlTree.AVLTree{}
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 1000; i++ {
		key := r.Intn(100000)
		data := r.Intn(100000)
		tree.Add(key, data)
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
