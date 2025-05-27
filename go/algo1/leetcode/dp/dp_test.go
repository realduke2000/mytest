package dp

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)
var (
	fadd = add
)

func TestDeleteAndEarn(t *testing.T) {
	ans := deleteAndEarn([]int{3, 4, 2})
	fmt.Println(ans)
}

func Test152(t *testing.T) {
	fmt.Println("hello")
	ans := maxProduct([]int{-1, -2, -9, -6})
	fmt.Println(ans)
}

func TestSquare(t *testing.T) {

	ans := numSquares(12)
	fmt.Println(ans)
	ans = numSquares(13)
	fmt.Println(ans)
}

func TestOutput(t *testing.T) {
	reqFilePath := "file1"
	reqFilePath = filepath.Join("/system/diagnostic/data-collect", reqFilePath)
	reqFilePath = filepath.Clean(reqFilePath)
	fmt.Println(reqFilePath)
	if strings.Index(reqFilePath, "/system/diagnostic/data-collect") == 0 {
		fmt.Println("legal!")
	}
}

func TestPrim(t *testing.T) {
	fmt.Println(countPrimes(10))
}

func TestCoin(t *testing.T) {
	ans := coinChange([]int{1, 2, 5}, 100)
	// ans := coinChange([]int{2}, 3)
	// ans := coinChange([]int{1}, 0)
	fmt.Println(ans)
}

func TestBase64(t *testing.T){
	s, _ := base64.RawURLEncoding.DecodeString("YXJmLnhtbA")
	fmt.Println(string(s))

	s, _ = base64.RawURLEncoding.DecodeString("YXJmLnhtbA=")
	fmt.Println(string(s))

	s, _ = base64.URLEncoding.DecodeString("YXJmLnhtbA==")
	fmt.Println(string(s))
}

func dg(n int) {
	if n == 1 {
		fmt.Print("A")
	} else {
		dg(n-1)
		m := byte('A' + n-1)
		fmt.Printf(string(m))
		dg(n-1)
	}
}

func TestDg(t *testing.T) {
	dg(4)
	fmt.Println()
}

func Test62(t *testing.T) {
	ans := uniquePaths(3, 7)
	fmt.Println(ans)
}
func Test64(t *testing.T) {
	input:=[][]int {[]int{1,3,1},[]int{1,5,1},[]int{4,2,1}}
	ans := minPathSum(input)
	fmt.Printf("%v\n", ans)
}

func Test5(t *testing.T) {
	assert.Equal(t, "bb", longestPalindrome("abb"))
	assert.Equal(t, "aaaa", longestPalindrome("aaaa"))
	assert.Equal(t, "ccc", longestPalindrome("ccc"))
	assert.Equal(t, "bab", longestPalindrome("babad"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
}


func add(a, b int) int {
	return a+b
}



func myadd() {
	fadd = func(a, b int) int {
		fmt.Println("inline functon")
		return 0
	}
	fmt.Println("exit myadd")
}

func TestNilFunc(t *testing.T) {
	myadd()
	fadd(1,2)
	fadd = nil
	fadd(1,2)
}


func Test63(t *testing.T) {
	// t.Run(
	// 	"test-1", func (t *testing.T)  {
	// 		data := make([][]int, 3)
	// 		data[0] = []int{0,0,0}
	// 		data[1] = []int{0,1,0}
	// 		data[2] = []int{0,0,0}
	// 		uniquePathsWithObstacles(data)
	// })
	// t.Run(
	// 	"test-2", func (t *testing.T)  {
	// 		data := make([][]int, 2)
	// 		data[0] = []int{0,0}
	// 		data[1] = []int{0,1}
	// 		uniquePathsWithObstacles(data)
	// })
	// t.Run(
	// 	"test-2", func (t *testing.T)  {
	// 		data := make([][]int, 2)
	// 		data[0] = []int{0}
	// 		data[1] = []int{1}
	// 		uniquePathsWithObstacles(data)
	// })
	t.Run(
		"test-2", func (t *testing.T)  {
			data := make([][]int, 3)
			data[0] = []int{0,0}
			data[1] = []int{1,1}
			data[2] = []int{0,0}
			uniquePathsWithObstacles(data)
	})
}

func TestNil(t *testing.T) {
	var certs []x509.Certificate
	fmt.Printf("len=%d\n", len(certs))
}

func TestDp120(t *testing.T) {
	triangle := make([][]int, 4)
	triangle[0] = []int{2}
	triangle[1]=[]int{3,4}
	triangle[2]=[]int{6,5,7}
	triangle[3] = []int{4,1,8,3}
	ans := minimumTotal(triangle)
	fmt.Println(ans)
}

func TestDp139(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		// {"leetcode", []string{"leet", "code"}, true},
		// {"applepenapple", []string{"apple", "pen"}, true},
		// {"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"aaaaaaa", []string{"aaaa", "aa"}, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("s=%s,wordDict=%v", tt.s, tt.wordDict), func(t *testing.T) {
			result := wordBreak(tt.s, tt.wordDict)
			assert.Equal(t, tt.expected, result)
		})
	}
}