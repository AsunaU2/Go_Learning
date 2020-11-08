package type_test

import (
	"testing"
)

// 1. go不允许隐式类型转换
// 2. 哪怕是定义别名也不行，如 type MyInt int64 a 是不能赋值给int64类型的b的
// 3. 最大数据在math里
// math.MaxInt64
// math.MaxFloat64
// math.MaxUint32 等等
// go不支持指针运算
// string是值类型，且默认值

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Log("true")
	}
}
