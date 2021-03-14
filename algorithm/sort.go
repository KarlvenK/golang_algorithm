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

func quickSort() {

}
