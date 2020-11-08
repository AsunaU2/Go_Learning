package operator_test

import "testing"

// 1. 长度不同的数组进行比较会得到编译错误： TestCompareArray

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c) // 错误 长度不同
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7 // 二进制 0111
	a = a &^ Readable
	t.Log(a)
}
