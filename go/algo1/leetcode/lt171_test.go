package leetcode

import (
	"fmt"
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	fmt.Printf("%v\n", titleToNumber("A"))
	fmt.Printf("%v\n", titleToNumber("AA"))
	fmt.Printf("%v\n", titleToNumber("AB"))
	fmt.Printf("%v\n", titleToNumber("ZA"))
	fmt.Printf("%v\n", titleToNumber("AZ"))
}

func TestPointer(t *testing.T) {
	p := new([]int)
	fmt.Printf("%v\n", p)
}

func TestNum(t *testing.T) {
	head := &ListNode{}
	data := []int{1, 2, 6, 3, 4, 5, 6}
	val := 6
	// data := []int{1, 1}
	// val := 2
	n := head
	for _, v := range data {
		n.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		n = n.Next
	}
	n.Next = nil

	ret := removeElements(head.Next, val)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}

func TestDeleteNode(t *testing.T) {
	head := &ListNode{}
	data := []int{0, 1, 2, 3, 4, 5, 6}
	// data := []int{1, 1}
	// val := 2
	n := head
	for _, v := range data {
		n.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		n = n.Next
	}
	n.Next = nil

	head = head.Next

	deleteNode(head.Next.Next)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func TestContain(t *testing.T) {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func TestContainBy(t *testing.T) {
	fmt.Println(containsNearbyDuplicate([]int{1, 1}, 2))
}

func TestPal(t *testing.T) {
	s := "0123456789"
	fmt.Printf("%s\n", s[3:5])
}

func Test146(t *testing.T) {
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
}

func TestSize(t *testing.T) {
	data := make([]int, 3)
	fmt.Println(len(data))
}
func TestCount1(t *testing.T) {
	fmt.Println(count1(2))

}

func TestDiameter(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	n := diameterOfBinaryTree(root)
	fmt.Println(n)
}

func TestDiameter2(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	n := diameterOfBinaryTree(root)
	fmt.Println(n)
}

func TestTree110(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}
	fmt.Println(isBalanced(root))
}

func TestContinue(t *testing.T) {
	ans := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	fmt.Println(ans)
}

func TestLogSub(t *testing.T) {
	ans := lengthOfLongestSubstring("bbbbb")
	fmt.Printf("ans=%d\n", ans)
}

func TestThree(t *testing.T) {
	ans := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Printf("ans: %v\n", ans)
}

func TestAnagram(t *testing.T) {
	ans := findAnagrams("ababababab", "aab")
	fmt.Printf("%v\n", ans)
}

func TestTwoSum(t *testing.T) {
	ans := twoSum([]int{-1, -2, -3, -4, -5}, -8)
	fmt.Println(ans)
}

func TestMaxSubArr(t *testing.T) {
	ans := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(ans)
}

func TestMerge(t *testing.T) {
	ans := merge([][]int{
		{1, 4},
		{0, 4},
	})
	fmt.Printf("%v\n", ans)
}

func TestChann(t *testing.T) {
	ch := make(chan int, 1)
	_, ok := <-ch
	if ok {
		fmt.Printf("ok")
	} else {
		fmt.Println("no ok")
	}
}
