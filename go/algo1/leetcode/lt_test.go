package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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


func TestRoman(t *testing.T) {
	assert.Equal(t, 4, romanToInt("IV"))
	assert.Equal(t, 1, romanToInt("I"))
	assert.Equal(t, 9, romanToInt("IX"))
	assert.Equal(t, 11, romanToInt("XI"))
	assert.Equal(t, 3, romanToInt("III"))
	assert.Equal(t, 58, romanToInt("LVIII"))
	assert.Equal(t, 1994, romanToInt("MCMXCIV"))
}

func TestLt655(t *testing.T) {
	fmt.Printf("%d\n", 1<<2)
	var root TreeNode
	root.Val= 1
	root.Left = &TreeNode{
		Val: 2,
	}
	printTree(&root)
}
func TestLt58(t *testing.T) {
	n := lengthOfLastWord("a")
	assert.Equal(t, 1, n)
	n = lengthOfLastWord("")
	assert.Equal(t, 0, n)
	n = lengthOfLastWord(" ")
	assert.Equal(t, 0, n)
}

func TestPrefix(t *testing.T) {
	longestCommonPrefix([]string{"", "a"})
}

func TestStrstr(t *testing.T) {
	strStr("abc", "c")
}

func TestLt125(t *testing.T) {
	assert.Equal(t, true,isPalindrome125("A man, a plan, a canal: Panama"))
	assert.Equal(t, false,isPalindrome125("race a car"))
	assert.Equal(t, true,isPalindrome125(" "))
	

	assert.Equal(t, true,isPalindrome2Pointers("A man, a plan, a canal: Panama"))
	assert.Equal(t, false,isPalindrome2Pointers("race a car"))
	assert.Equal(t, true,isPalindrome2Pointers(" "))
	assert.Equal(t, true,isPalindrome2Pointers("a."))
}