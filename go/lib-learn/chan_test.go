package lib_learn

import (
	"fmt"
	"testing"
)

func Test_testchan(t *testing.T) {
	testchan()
}

type OpError struct {
	error
}

func (e *OpError) Error() string { return "operation error!" }

type RuntimeError struct {
	OpError
}

func TestErr(t *testing.T) {
	var e RuntimeError
	fmt.Printf("%s\n", e.Error())
}
