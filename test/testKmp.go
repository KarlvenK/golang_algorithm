package test

import "fmt"
import "algorithm/algorithm/myStr"

func Kmp() {
	s := "faklsdjfladsadf"
	t := "sdjf"
	fmt.Print(myStr.IndexKmp(s, t))
}
