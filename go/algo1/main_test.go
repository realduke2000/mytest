package main

import (
	"fmt"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	r := TreeNode{
		0,
		&TreeNode{
			1,
			nil,
			nil,
		},
		&TreeNode{
			2,
			&TreeNode{
				3,
				nil,
				nil,
			},
			nil,
		},
	}
	ret := preorderTraversal(&r)
	for _, v := range ret {
		fmt.Println(v)
	}
}

func myapp(ret []int) {
	ret = append(ret, 5)
}

func TestArray(t *testing.T) {
	ret := make([]int, 5)
	ret[0] = 0
	ret[1] = 1
	ret[2] = 2
	myapp(ret[:])
	fmt.Printf("%v\n", ret)
}

func TestSize(t *testing.T) {
	var m1 map[int]int
	m2 := make(map[int]int, 5)
	fmt.Printf("m1=%d\n", len(m1))
	fmt.Printf("m2=%d\n", len(m2))
}
