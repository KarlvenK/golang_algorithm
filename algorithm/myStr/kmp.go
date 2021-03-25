package myStr

func getNext(t string) []int {
	n := len(t)
	next := make([]int, n)
	next[0] = -1
	i, j := 0, -1

	for i < n-1 {
		if j == -1 || t[i] == t[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func IndexKmp(s, t string) int {
	i, j := 0, 0
	next := getNext(t)
	for i < len(s) && j < len(t) {
		if j == -1 || s[i] == t[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j >= len(t)-1 {
		return i - len(t)
	}
	return -1
}
