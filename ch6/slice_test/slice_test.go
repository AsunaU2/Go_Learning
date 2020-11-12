package slice_test

import "testing"

/*
切片：类似于可变长数组，但不是可变长数组
切片和数据区别：
1. 数据不可边长
2. 数组可比较，切片不行
3. 切片：a := []int{1,2,3,4}
   数组：a := [...]int{1,2,3,4}
数据结构
|--------|
|  *ptr  | - 切片指针
|   len  | - 可用数据个数
|   cap  | - 可用空间 相当于vector的capacity
*/

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// make创建切片
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
}

/*
切片的数据增长方式：
主要是cap的增长方式：和vector扩容机制一样，也会进行拷贝
*/
func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

/*
切片：
1. 本身是浅拷贝，会共享内存数据
2. 切片切出数据，但是空间会切到最后，也就是说CAP是从切片头开始计算直至这块存储空间的尾部
3. Slice和数组不同，相互间是不能比较的，仅能和空-nil进行比较
*/
func TestSliceShareMemory(t *testing.T) {
	year := []string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) // Q2长度是3, 但是cap是9

	summer := year[5:8]
	summer[0] = "Unknow"
	t.Log(Q2) // 改了Summer，但Q2也变了
	t.Log(summer)
	t.Log(year) // 同样year也变了
}

/*
切片不能相互比较，只能和nil比较
*/
func TestSliceComparing(t *testing.T) {
	if a := []int{1, 2, 3, 4}; a != nil {
		t.Log("oops!")
	}
}
