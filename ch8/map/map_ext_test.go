package _map

import "testing"

func TestMapWithFunValue(t *testing.T) {
    m := map[int]func(op int) int{}
    m[1] = func(op int) int { return op }
    m[2] = func(op int) int { return op * op }
    m[3] = func(op int) int { return op * op * op }
    t.Log(m[1](2), m[2](2), m[3](2))
}

/*
golang中没有set  这里通过map模拟set
 */
func TestMapForSet(t *testing.T) {
    mySet := map[int]bool{}
    mySet[1] = true   // add
    mySet[2] = true
    mySet[3] = true

    exist := mySet[1] // 判断是否存在

    t.Log("mySet[1] is exist, it is ", exist)

    for key := range mySet {
        t.Log(key)
    }

    delete(mySet, 1)  // 删除
}