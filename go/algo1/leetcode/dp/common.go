package dp

import "fmt"

func displayIntArr2(data [][]int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			fmt.Printf("%d ", data[i][j])
		}
		fmt.Println()
	}
}