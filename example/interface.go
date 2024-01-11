package example

import (
	"fmt"
)

// 定义接口
type EatWhat interface {
	EatMeat(data interface{}) error
	LikeSleep() bool
}

// 定义类型结构1
type me struct {
}

// 定义类型结构2
type he struct {
}

// 定义类型结构3
type she struct {
}

func (I *me) EatMeat(data interface{}) error {
	fmt.Println("I like eat meat:!!!!: data：", data)

	return nil
}
func (I *me) LikeSleep() bool {
	return true
}

func (H *he) EatMeat(data interface{}) error {
	fmt.Println("he does not like meat!!!data:", data)
	return nil
}
func (H *he) LikeSleep() bool {
	return true
}

func (S *she) EatMeat(data interface{}) error {
	fmt.Println("she also likes meat!!!Data：", data)
	return nil
}
func (S *she) LikeSleep() bool {
	return false
}

func InterMain() {
	//实例化me结构体
	fm := new(me)
	fh := new(he)
	fs := new(she)

	//声明一个EatWhat的接口
	var ew EatWhat

	//将接口赋值结构体的实例化，即me类型
	ew = fm

	ew.EatMeat("dataaaaaa")
	fmt.Println(ew.LikeSleep())

	ew = fh
	ew.EatMeat("hhhhhhh")
	fmt.Println(ew.LikeSleep())

	ew = fs
	ew.EatMeat("ssssssss")
	fmt.Println(ew.LikeSleep())
}
