package lib_learn

import (
	"context"
	"fmt"
	"time"
)

func timeout() {
	ctx := context.Background()
	ctx1, cancel := context.WithCancel(ctx)
	go func(c context.Context) {
		for {
			select {
			case <-c.Done():
				fmt.Printf("ctx done: %v\n", c.Err())
				return
			default:
				fmt.Println("routing is running")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx1)

	time.Sleep(3 * time.Second)
	cancel()
	fmt.Println("canceled...")
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
	}
	fmt.Println("exit")
}

func readvalue() {
	ctx := context.Background()
	ctx1 := context.WithValue(ctx, "key1", "value1")
	fmt.Printf("%s\n", ctx1.Value("key1"))
}
