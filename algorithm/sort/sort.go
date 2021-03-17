package sort

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

// 该函数会把 data[a, b) 区间的元素进行排序，下面称该区间为切片 slice
func quickSort(data Interface, a, b, maxDepth int) {
	// 如果切片长度不大于 12 ，则使用希尔排序，否则，使用下面的方法排序
	for b-a > 12 {
		if maxDepth == 0 { // 如果递归到最大深度，则使用堆排序
			heapSort(data, a, b)
			return
		}
		maxDepth--
		// doPivot 是快排核心算法，它取一点为轴，把不大于轴的元素放左边，大于轴的元素放右边，返回小于轴部分数据的最后一个下标，以及大于轴部分数据的第一个下标
		// 下标位置 a...mlo,pivot,mhi...b
		// data[a...mlo] <= data[pivot]
		// data[mhi...b] > data[pivot]
		mlo, mhi := doPivot(data, a, b)
		// 避免较大规模的子问题递归调用，保证栈深度最大为 maxDepth
		// 解释：因为循环肯定比递归调用节省时间，但是两个子问题只能一个进行循环，另一个只能用递归。这里是把较小规模的子问题进行递归，较大规模子问题进行循环。
		if mlo-a < b-mhi {
			quickSort(data, a, mlo, maxDepth)
			a = mhi // 相当于 quickSort(data, mhi, b)
		} else {
			quickSort(data, mhi, b, maxDepth)
			b = mlo // 相当于 quickSort(data, a, mlo)
		}
	}

	// 较小数据集使用希尔排序
	// 第一次步长为 6，第二次步长为 1（其实就是插入排序了）
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
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
	m := int(uint(lo+hi) >> 1) // trick:避免整型溢出的
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree(data, lo, lo+s, lo+2*s)
		medianOfThree(data, m, m-s, m+s)
		medianOfThree(data, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThree(data, lo, m, hi-1)

	// 以下代码达成目标为:
	// data[lo] = pivot (set up by ChoosePivot)
	// data[lo < i < a] < pivot
	// data[a <= i < b] <= pivot
	// data[b <= i < c] unexamined
	// data[c <= i < hi-1] > pivot
	// data[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && data.Less(a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !data.Less(pivot, b); b++ { // data[b] <= pivot
		}
		for ; b < c && data.Less(pivot, c-1); c-- { // data[c-1] > pivot
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c-1] <= pivot
		data.Swap(b, c-1)
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if !data.Less(pivot, hi-1) { // data[hi-1] = pivot
			data.Swap(c, hi-1)
			c++
			dups++
		}
		if !data.Less(b-1, pivot) { // data[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> data[m] <= pivot
		if !data.Less(m, pivot) { // data[m] = pivot
			data.Swap(m, b-1)
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		// data[a <= i < b] unexamined
		// data[b <= i < c] = pivot
		for {
			for ; a < b && !data.Less(b-1, pivot); b-- { // data[b] == pivot
			}
			for ; a < b && data.Less(a, pivot); a++ { // data[a] < pivot
			}
			if a >= b {
				break
			}
			// data[a] == pivot; data[b-1] < pivot
			data.Swap(a, b-1)
			a++
			b--
		}
	}
	// Swap pivot into middle
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
		data.Swap(first, first+i)
		siftDown(data, lo, i, first)
	}
}
