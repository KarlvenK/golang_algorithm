package algorithm

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Interface) {
	n := data.Len()
	quickSort(data, 0, n, maxDepth(n))
}

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func quickSort(data Interface, a, b, maxDepth int) {
	if b-a > 12 { //if the len of slice > 12, use shellsort else use heapsort
		if maxDepth == 0 {
			heapSort(data, a, b)
			return
		}
	}
	maxDepth--
	mlo, mhi := doPivot(data, a, b)
	// data[a..mlo] <= data[pivot] < data[mhi..b]
	if mlo-a < b-mhi {
		quickSort(data, a, mlo, maxDepth)
		a = mhi
	} else {
		quickSort(data, mhi, b, maxDepth)
		b = mlo
	}
	if b-a > 1 {
		for i := a + 6; i < b; i++ {
			if data.Less(i, i-6) {
				data.Swap(i, i-6)
			}
		}
		insertionSort(data, a, b)
	}
}

func insertionSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func doPivot(data Interface, lo, hi int) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1) // avoid overflowing
	if hi-lo > 40 {
		s := (hi - lo) / 8
		medianOfThree(data, lo, lo+s, lo+s*2)
		medianOfThree(data, m, m-s, m+s)
		medianOfThree(data, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThree(data, lo, m, hi-1)
	/*
		data[lo] = pivot(set up by chiisePivot)
		data[lo < i < a] < pivot
		data[a <= i < b] <= pivot
		data[b <= i < c] unexamined
		data[c <= i < hi - 1] > pivot
		data[hi - 1] >= pivot
	*/
	pivot := lo
	a, c := lo+1, hi-1
	for ; a < c && data.Less(a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !data.Less(pivot, b); b++ {
		}
		for ; b < c && data.Less(pivot, c-1); c-- {
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c - 1] <= pivot
		data.Swap(b, c-1)
		b++
		c--
	}
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		dups := 0
		if !data.Less(pivot, hi-1) {
			data.Swap(c, hi-1)
			c++
			dups++
		}

		if !data.Less(m, pivot) {
			data.Swap(m, b-1)
			b--
			dups++
		}
		protect = dups > 1
	}

	if protect {

		for {
			for ; a < b && !data.Less(b-1, pivot); b-- {
			}
			for ; a < b && data.Less(a, pivot); a++ {
			}
			if a >= b {
				break
			}
			data.Swap(a, b-1)
			a++
			b--
		}
	}
	data.Swap(pivot, b-1)
	return b - 1, c
}

func medianOfThree(data Interface, m1, m0, m2 int) { //sort the three members
	if data.Less(m1, m0) {
		data.Swap(m1, m0)
	}
	if data.Less(m2, m1) {
		data.Swap(m2, m1)
		if data.Less(m1, m0) {
			data.Swap(m1, m0)
		}
	}
}

func siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}

		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

func heapSort(data Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	for i := hi - 1; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+1)
		siftDown(data, lo, i, first)
	}
}
