package stack

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestStack(tt *testing.T) {
	s := New()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	cont := make([]int, 0)
	for i := 0; i < 1000000; i++ {
		if v := r.Int(); v&1 == 1 {
			s.Push(v)
			cont = append(cont, v)
		} else {
			if t := s.Pop(); t != nil {
				if t != cont[len(cont)-1] {
					tt.Fatal("stack failed to pass the test")
					return
				} else {
					cont = cont[:len(cont)-1]
				}
			}
		}
	}

	for !s.Empty() {
		temp := s.Pop()
		if temp != cont[len(cont)-1] {
			fmt.Println("stack failed to pass the test")
		} else {
			cont = cont[:len(cont)-1]
		}
	}
}
