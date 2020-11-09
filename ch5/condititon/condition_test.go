package condititon

import "testing"

/*
// go中if不同的要记下的是，可以
   if statement; condition {
   }
    即，可以先写一句赋值语句，再进行判断，和for有点像
 */
func TestInfMultiSec (t *testing.T) {
    if a:= 2; a == 1 {
        t.Log(a)
    } else {
        t.Log(a + 1)
    }
}