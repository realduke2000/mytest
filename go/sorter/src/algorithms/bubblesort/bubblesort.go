// bubblesort.go
package bubblesort

func Bubblesort(values []int){

	flag := true

	for i := len(values) - 1; i > 0; i-- {
		flag = true
		for j := 0; j < i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}

		if flag == true {
			break
		}
	}
}
