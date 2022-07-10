package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1
	a := 1
	b := 1
	//fmt.Print(a, " ")
	t.Log(a)
	for i := 0; i < 10; i++ {
		//fmt.Print(" ", b)
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
		//	fmt.Println()
	}

}
