package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
1. 函数可以返回多个值
2. 所有参数都是值传递: Slice、map、channel会有引用的错觉，原因是底层实现浅拷贝
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值
*/

/* 函数作为参数进行传递
该函数计算传入函数的计算时间
例子中函数传入一个整型值

其实是一个装饰器的写法
*/
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

/*
返回多个值
*/
func returnMultiValues() (int, int) {
	rand.Seed(3)
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

/*
可变长参数
*/
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5, 6))
}

/*
defer函数
1. 在函数返回前延迟执行
2. 主要用于资源回收，相当于python的异常处理的finally
3. 即使panic中断函数，也仍然会执行
*/

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}
