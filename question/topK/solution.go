package topK

import (
	"container/heap"
	"math/rand"
	"sync"
	"time"
)

const tot = 100000000

/*
把 100 000 000 个整数分成100个大组
再把大组分成100个小组
每个小组10000个整数
*/

type node []int

func (bar node) Len() int {
	return len(bar)
}

func (bar node) Less(i, j int) bool {
	return bar[i] < bar[j]
}

func (bar node) Swap(i, j int) {
	bar[i], bar[j] = bar[j], bar[i]
}

func (bar *node) Push(x interface{}) {
	num := x.(int)
	*bar = append(*bar, num)
}

func (bar *node) Pop() interface{} {
	old := *bar
	n := bar.Len()
	ret := old[n-1]
	*bar = old[0 : n-1]
	return ret
}

func Generate() []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < tot; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func Solve(nums []int) (ans []int) {
	var getTopHundred func([]int) []int
	getTopHundred = func(num []int) []int {
		pq := make(node, 0)
		for i := 0; i < 100; i++ {
			pq = append(pq, num[i])
		}
		heap.Init(&pq)
		for i := 100; i < len(num); i++ {
			if num[i] > pq[0] {
				heap.Pop(&pq)
				heap.Push(&pq, num[i])
			}
		}
		return pq
	}

	var WG sync.WaitGroup
	var MU sync.Mutex
	temp := make([]int, 0)

	for i := 0; i < 100; i++ {
		ret := make([]int, 0)
		var mu sync.Mutex
		var wg sync.WaitGroup

		sub := nums[i*1000000 : (i+1)*1000000]
		for j := 0; j < 100; j++ {
			subSub := sub[j*10000 : (j+1)*10000]
			wg.Add(1)
			go func() {
				tmp := getTopHundred(subSub)
				mu.Lock()
				ret = append(ret, tmp...)
				mu.Unlock()
				wg.Done()
			}()
		}
		wg.Wait()

		WG.Add(1)
		go func() {
			tmp := getTopHundred(ret)
			MU.Lock()
			temp = append(temp, tmp...)
			MU.Unlock()
			WG.Done()
		}()
	}
	WG.Wait()
	ans = getTopHundred(temp)

	return
}

func Force(nums []int) (ans []int) {
	var getTopHundred func([]int) []int
	getTopHundred = func(num []int) []int {
		pq := make(node, 0)
		for i := 0; i < 100; i++ {
			pq = append(pq, num[i])
		}
		heap.Init(&pq)
		for i := 100; i < len(num); i++ {
			if num[i] > pq[0] {
				heap.Pop(&pq)
				heap.Push(&pq, num[i])
			}
		}
		return pq
	}

	ans = getTopHundred(nums)
	return
}
