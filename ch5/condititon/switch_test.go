package condititon

import (
    "testing"
)

/*
1. 基本用法，与别的语言差不多，区别在于不用显式break，默认就是带break的
 */

func TestFuncSwitchBaseUse(t *testing.T) {
    var a int

    switch a {
    case 10:
        t.Log(10)
    case 5:
        t.Log(5)
    case 0:
        t.Log(0)
    default:
        t.Log("other")
    }
}

/*
fallthrough
加上fallthrough就跟c++中没有加break一样
 */

func TestSwitchFallthrough(t *testing.T) {
    var a int = 10
    switch a {
    case 10:
        t.Log(10)
        fallthrough // 会在命中10后继续命中5
    case 5:
        t.Log(5)

    }
}

/*
1. 多个case
 */

func TestSwitchMultiCase(t *testing.T) {
    var a int = 10
    switch a {
    case 10, 5:
        t.Log(10, " or ", 5)
    }
}

/*
1. switch后带statement,可以带一句声明，判断的则是后面那个
 */
func TestSwitchStatement(t *testing.T) {
    switch a := 10; a {
    case 10:
        t.Log(10)

    }
}

/*
1. 无条件表达式switch, 可以把条件放在case，相当与if用法了
 */
func TestSwitchEmpty(t *testing.T) {
    var a int = 10
    var b int = 20
    switch {
    case a > 6:
        t.Log("a > 6")
    case b > 6:
        t.Log("a < 6")
    }
}
