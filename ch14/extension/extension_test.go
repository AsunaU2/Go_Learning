package extension

import (
	"fmt"
	"testing"
)

// 定义Pet类
type Pet struct {
}

// 定义Pet类的public方法
func (p *Pet) Speak() {
	fmt.Print("...")
}

// 定义Pet类的Public方法
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Print(" ", host)
}

type Dog struct {
	p *Pet
}

// 定义Dog类的public方法
func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

// 定义Dog类的Public方法
func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println(" ", host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Yang")
}
