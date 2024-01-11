package example

import "fmt"

func AssertionMain() {
	var x interface{}
	x = 200

	value, ok := x.(int)

	fmt.Println(value, ok)
}
