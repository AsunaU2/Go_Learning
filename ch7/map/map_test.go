package map_test

import "testing"

/*
map可以make capacity, 但是不能打印：cap(m3)
*/
func TestInitMap(t *testing.T) {
    m1 := map[int]int{1: 1, 2: 4, 3: 9}
    t.Log(m1[2])
    t.Log("len m1: ", len(m1))
    m2 := map[int]int{}
    m2[4] = 16
    t.Log("len m2: ", len(m2))
    m3 := make(map[int]int, 10)
    t.Log("len m3: ", len(m3))
}

/*
map对于不存在的key处理比较有意思
别的语言都是直接报空如java, C++是需要用find确认key是否存在
map是直接初始化为0
所以如何判断一个map里是没有这个key，还是这个key对应的值本身就是0呢
需要用两个值来接map的返回值，第二个值代表这个key是否存在，bool类型

 */
func TestAccessNoExistingKey(t *testing.T) {
    m1 := map[int]int{}
    t.Log(m1[1])
    m1[2] = 0
    t.Log(m1[2])
    m1[3] = 54

    if v, ok := m1[3]; ok {
        t.Log("key 3's value is ", v)
    } else {
        t.Log("key 3 is not existing.")
    }
}

//map遍历
func TestTravelMap (t *testing.T) {
    m1 := map[int]int{1: 1, 2: 4, 3: 9}

    for k,v := range m1 {
        t.Log(k, v)
    }
}
