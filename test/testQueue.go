package test

import (
	"fmt"
	"golang_algorithm/algorithm/queue"
	"math/rand"
	"time"
)

func Queue() {
	que := queue.New()
	cont := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 1000000; i++ {
		v := r.Intn(112312312312)
		if v&1 == 1 {
			que.Push(v)
			cont = append(cont, v)
		} else {
			if que.Empty() {
				continue
			}
			if que.Pop() != cont[0] {
				fmt.Println("queue failed to pass the test")
				return
			}
			cont = cont[1:]
		}
	}
	fmt.Println("queue passed the test")
	printSplitLine()
}
