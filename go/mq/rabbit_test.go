package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_testMQConn(t *testing.T) {
	c1 := make(chan bool, 2)
	c2 := make(chan bool, 1)
	go func() {
		fmt.Println("sleep ")
		time.Sleep(1 * time.Second)
		c1 <- true
		time.Sleep(1 * time.Second)
		c2 <- false
		fmt.Println("send signal to c1")
	}()

	go func() {
		fmt.Println("sleep ")
		time.Sleep(1 * time.Second)
		c2 <- true
		time.Sleep(1 * time.Second)
		c1 <- false
		fmt.Println("send signal to c1")
	}()

	select {
	case <-c1:
		fmt.Println("c1")
	case <-c2:
		fmt.Println("c2")
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	}

	fmt.Println("finish")
}
