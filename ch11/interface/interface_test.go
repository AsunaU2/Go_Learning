package _interface

import "testing"

/*
1. Go的多态通过鸭子类型进行实现
2. 因此它的多态是非侵入式的
*/

// 定义接口
type Programmer interface {
	WriteHelloWorld() string
}

// 定义实现 - 结构体（类）
type GoProgrammer struct {
	Language string
}

// 定义类的方法 // 为何这样声明呢，答案很简单，让stProg对于WriteHelloWorld这个方法可见，当然了这里我没用到就是了
func (stProg *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello world\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
