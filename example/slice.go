package example

import "fmt"

func SliceMain() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("数组arr的值为", arr)
	fmt.Println("切片arr[0:1]的值为", arr[0:1])
	fmt.Println("切片arr[2:8]的值为", arr[2:8])
	fmt.Println("切片arr[4:9]的值为", arr[4:9])
	fmt.Println("切片arr[5:10]的值为", arr[5:10])
	fmt.Println("切片arr[5:10]中元素个数为", len(arr[5:10]))
	fmt.Println("切片最后一个元素arr[len(arr)]的值为", arr[len(arr[4:9])])
	fmt.Println("切片arr[2:]的值为", arr[2:])
	fmt.Println("切片arr[:8]的值为", arr[:8])
	fmt.Println("切片arr[:]的值为", arr[:])
}
