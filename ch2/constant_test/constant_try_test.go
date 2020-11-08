package constant_test

import "testing"
import "fmt"

const (
	Monday = 1
	Tuesday
	Wednesday
	Thursday
	Friday
)

const (
	Readable = 1 << iota
	Writable
	Executable
	xixiable
)

const a = iota // https://www.cnblogs.com/OctoptusLian/p/9373421.html 描述iota用法

const (
	b = iota
)

const (
	m=1<<iota // 1 iota = 0
	n=2<<iota // 4(0010左移1位=0100) iota=1
	x=10 // 10 iota=2
	y = iota // 3 iota=3
	z=iota>>1 // 2(0100右移1位=0010） iota=5
	o // 2 向上复制，且恢复技术
	j // 3
)

func TestConstantTry(t *testing.T) {
	fmt.Println("m=",m)
	fmt.Println("n=",n)
	fmt.Println("x=",x)
	fmt.Println("y=",y)
	fmt.Println("z=",z)
	fmt.Println("o=",o)
	fmt.Println("j=",j)
}

func TestConstantTry2(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday)
}
