package lib_learn

import "fmt"

func testchan() {
	var c1 chan bool = make(chan bool, 1)
	c1 <- true
	close(c1)
	// fmt.Println("write to closed chann")
	// c1 <- 1
	fmt.Println("read from closed chan")
	for i := 0; i < 3; i++ {
		select {
		case i := <-c1:
			fmt.Printf("c1: %v\n", i)
		default:
			fmt.Println("no block")
		}
	}
}
