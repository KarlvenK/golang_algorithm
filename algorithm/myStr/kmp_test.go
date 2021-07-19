package myStr

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func force(s, t string) int {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j >= len(t) {
		return i - len(t)
	}
	return -1
}

func TestKmp(tt *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 10000; i++ {
		var s, t bytes.Buffer
		for j := 0; j < 1000; j++ {
			s.WriteByte(byte(r.Intn(5)))
		}
		for j := 0; j < 10; j++ {
			t.WriteByte(byte(r.Intn(5)))
		}
		a, b := IndexKmp(s.String(), t.String()), force(s.String(), t.String())
		if a != b {
			tt.Fatal("noooooooo")
			return
		} else {
			if a != -1 {
				fmt.Println(a, b)
			}
		}
	}
}
