package topK

import (
	"testing"
)

var nums []int

func init() {
	nums = Generate()
}

func TestGenerate(t *testing.T) {
	if len(nums) != tot {
		t.Fatal("cant't generate enough numbers")
	}
	t.Log(len(nums))
}

func TestSolve(t *testing.T) {
	_ = Solve(nums)
}

func TestForce(t *testing.T) {
	_ = Force(nums)
}
