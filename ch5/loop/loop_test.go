package loop_test

import "testing"


/*
// 无线循环的写法：
n:=0
for {
    t.Log(n)
    n++
}
*/

func TestWhileLoop(t *testing.T) {
    n :=0
    for n != 5 {
        t.Log(n)
        n++
    }
}
