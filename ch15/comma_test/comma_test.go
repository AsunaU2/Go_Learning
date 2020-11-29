package comma_test

import (
	"bytes"
	"testing"
)

// comma 将字符串从右往左每3位增加逗号
func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	var buf bytes.Buffer
	n := len(s)

	for i := 0; i < n; i++ {

		if (n-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')

		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

// func twyg(s1 string, s2 string) bool {
// 	bRet := true

// 	for i := 0; i < 1; i++ {
// 		if s1 == s2 || len(s1) != len(s2) {
// 			bRet = false
// 			break
// 		}

// 		for j := 0; j < len(s1); j++ {
// 			for k, bSame := 0, false; k < len(s1); k++ {
// 				if s1[i] != s2[k] {
// 					bRet = false
// 					break
// 				}
// 			}
// 		}
// 	}

// 	return bRet
// }

func TestComma(t *testing.T) {
	t.Log(comma("12345sdfsdf26"))
}
