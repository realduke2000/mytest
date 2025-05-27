package main

import (
	"encoding/base64"
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

type Password struct {
}

func (p Password) Validate() {
	fmt.Println("Validate for Password")
}

type PasswordV2 struct {
	p Password
}

func TestEmbedded(t *testing.T) {

	teststr := "TutorialsPoint?java"
	fmt.Printf("StdURLDecoding:\t%s\n", base64.StdEncoding.EncodeToString([]byte(teststr)))
	fmt.Printf("URLDecoding:\t%s\n", base64.URLEncoding.EncodeToString([]byte(teststr)))
	fmt.Printf("RawURLDecoding:\t%s\n", base64.RawURLEncoding.EncodeToString([]byte(teststr)))

	name, err := base64.URLEncoding.DecodeString("VHV0b3JpYWxzUG9pbnQ/amF2YQ==")
	if err != nil {
		t.Failed()
	}
	fmt.Printf("URLDecoding:\t%s\n", string(name))

	name, err = base64.RawURLEncoding.DecodeString("VHV0b3JpYWxzUG9pbnQ_amF2YQ==")
	if err != nil {
		t.Failed()
	}
	fmt.Printf("RawURLDecoding:\t%s\n", string(name))
}
