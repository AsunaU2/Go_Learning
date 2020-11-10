package array_test

import "testing"

/*
1. 数组声明，见下面三种方式
2. 数组初始化，默认为0
3. 数组长度是固定的，不同于切片是可伸缩的，可用append追加
4. 数据比较时，同长度和同维度是可以比较的
*/
func TestArrayInit(t *testing.T) {
	var arr [3]int

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

/*
1. 数组遍历，2种方式
*/
func TestArrayTraval(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}

	// 方式1
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	/*
	   1. 通过range，会返回索引和索引下的值
	   2. 如果返回值不想用，但是又想通过编译，则用_进行占位
	*/
	for _, e := range arr3 {
		t.Log(e)
	}
}

/*
数组切片
切片仅仅是浅拷贝！
*/
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	arr3Sec := arr3[2:]
	t.Log(arr3Sec, len(arr3Sec), cap(arr3Sec))

	arr3[4] = 100
	t.Log(arr3Sec, len(arr3Sec), cap(arr3Sec))

}
