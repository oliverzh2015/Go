package fib

import "testing"

func TestExchange(t *testing.T) {
	a := 1
	b := 3
	//tmp := a
	//a = b
	//b = tmp
	a, b = b, a
	t.Log(a, b)
}
