package tempconv

import "fmt"

// Celsius 摄氏度
type Celsius float64

// Fahrenheit 华氏度
type Fahrenheit float64

// Kelvin 开尔文
type Kelvin float64

// 温度常量定义
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°c", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g K", k)
}
