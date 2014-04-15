package qsort 

import (
	"testing"
	"math/rand"
	"time"
	"strconv"
)

func TestQSort(t *testing.T){
	len := 10
	max := 100
	values := make([]int, len)
	var msg string

	rand.Seed(time.Now().Unix())
	for i := 0; i < len; i++ {
		values[i] = rand.Intn(max)
	}

	msg = ""
	for i := 0; i< len; i++ {
		msg += strconv.Itoa(values[i]) + " "
	}

	msg = "Before sort: " + msg
	t.Log(msg)

	Qsort(values)

	msg = ""
	for i := 0; i< len; i++ {
		msg += strconv.Itoa(values[i]) + " "
	}
	msg = "After sort: " + msg
	t.Log(msg)

	for i := 0; i < len - 1; i++{
		if values[i] > values[i + 1] {
			t.Error("Sort failed, see log for details")
		}
	}
}

func TestQSort1(t *testing.T){
	len := 10
	var msg string

	values := []int {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	msg = ""
	for i := 0; i< len; i++ {
		msg += strconv.Itoa(values[i]) + " "
	}

	msg = "Before sort: " + msg
	t.Log(msg)

	Qsort(values)

	msg = ""
	for i := 0; i< len; i++ {
		msg += strconv.Itoa(values[i]) + " "
	}
	msg = "After sort: " + msg
	t.Log(msg)

	for i := 0; i < len - 1; i++{
		if values[i] > values[i + 1] {
			t.Error("Sort failed, see log for details")
		}
	}
}
