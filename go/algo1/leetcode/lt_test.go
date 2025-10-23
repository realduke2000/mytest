package leetcode

import (
	"fmt"
	"testing"
)

func TestLt27(t *testing.T) {
	data := []int{0, 1, 2, 2, 3, 0, 4, 2}
	k := removeElement(data, 2)
	fmt.Println(k)
	fmt.Printf("data=%v\n", data)
}
func TestLt80(t *testing.T) {
	for i, data := range [][]int{
		{1, 2, 2},
		{1, 1, 1, 2, 2, 3},
		{0, 0, 1, 1, 1, 1, 2, 3, 3},
	} {
		ret := removeDuplicates80(data)
		fmt.Printf("i=%d, len=%d, data=%v\n", i, ret, data)
	}
}

func TestLt80Stack(t *testing.T) {
	for i, data := range [][]int{
		{1, 2, 2},
		{1, 1, 1, 2, 2, 3},
		{0, 0, 1, 1, 1, 1, 2, 3, 3},
	} {
		ret := removeDuplicatesStack(data)
		fmt.Printf("i=%d, len=%d, data=%v\n", i, ret, data)
	}
}

func TestLt121(t *testing.T) {
	n := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(n)
}

func TestLt55(t *testing.T) {
	// canJump([]int{2, 3, 1, 1, 4})
	// canJump([]int{0, 1})
	// canJump([]int{1, 2, 3})
	canJump([]int{3, 2, 1, 0, 4})
}
