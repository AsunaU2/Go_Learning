package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	// 表达方式1
	//var a int = 1
	//var b int = 1

	// 表达方式2
	//var (
	//	a int = 1
	//	b int = 1
	//)

	// 表达方式3
	a := 1
	b := 1

	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
