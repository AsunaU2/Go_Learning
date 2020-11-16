package _string

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 默认初始化值为""

	s = "hello"
	t.Log(len(s))

	// s[1] = '3' //  Error, string内部是只读byte slice, 不支持更改

	s = "\xE4\xB8\xA5" // string可以存储任何二进制数据
	t.Log(s, len(s))

	s = "中华小当家"
	t.Log(s)

	/*
	   rune是关键字，它可以以unicode字符的形式来进行计算
	   即： len(string("中") = 3 // 3个byte
	        len([]rune(string("中") = 1 //  1个汉字
	*/
	c := []rune(s)
	t.Log(len(c))

	t.Logf("中 unicode %x", c[0])
	t.Logf("中 utf8 %x", s[0])

	// 此时索引值k是以byte进行计算的，所以是0, 3, 6, 9, 12
	for k, v := range s {
		t.Log(k, v)
	}
	// 此时索引值k是以unicode字符计算的，所以是0, 1, 2, 3, 4, 5
	for k, v := range c {
		t.Logf("key: %d, value: %[2]c %[2]d", k, v)
	}
}
