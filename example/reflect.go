package example

import (
	"fmt"
	"reflect"
)

func ReflectMain() {
	x := 3.1415
	y := 3
	z := "sheena"
	u := true

	//获取变量的类型
	xtype := reflect.TypeOf(x)
	ytype := reflect.TypeOf(y)
	ztype := reflect.TypeOf(z)
	utype := reflect.TypeOf(u)

	//Kind()方法的返回结果主要是用来进行类型判断的
	xkind := xtype.Kind()
	ykind := ytype.Kind()
	zkind := ztype.Kind()
	ukind := utype.Kind()

	fmt.Println("x的value为：", reflect.ValueOf(x))
	fmt.Println("y的value为：", reflect.ValueOf(y))
	fmt.Println("z的value为：", reflect.ValueOf(z))
	fmt.Println("u的value为：", reflect.ValueOf(u))

	//对x进行类型判断
	if xkind == reflect.String {
		fmt.Println("x的type是string")
	} else if xkind == reflect.Float64 {
		fmt.Println("x的type是float64")
	}

	//对y进行类型判断
	if ykind == reflect.Float64 {
		fmt.Println("y的type是Float64")
	} else if ykind == reflect.Int {
		fmt.Println("y的type是int")
	}

	//对z进行类型判断
	if zkind == reflect.Float64 {
		fmt.Println("z的type是Float64")
	} else if zkind == reflect.String {
		fmt.Println("z的type是string")
	}

	//对u进行类型判断
	if ukind == reflect.Bool {
		fmt.Println("u的type是bool")
	} else if ukind == reflect.String {
		fmt.Println("u的type是string")
	}
	//通过反射传入变量x的地址，并且通过Ele
	rex := reflect.ValueOf(&x).Elem()
	rey := reflect.ValueOf(&y).Elem()
	rez := reflect.ValueOf(&z).Elem()
	reu := reflect.ValueOf(&u).Elem()

	//判断是否可以修改变量x的值，若可以，则用SetFLoat64()方法进行修改
	if rex.CanSet() {
		rex.SetFloat(61.23466)
		fmt.Println("x修改后的value为：", reflect.ValueOf(x))
	} else {
		fmt.Println("该变量不能修改")
	}

	if rey.CanSet() {
		rey.SetInt(10000)
		fmt.Println("y修改后的value为：", reflect.ValueOf(y))
	} else {
		fmt.Println("该变量不能修改")
	}

	if rez.CanSet() {
		rez.SetString("hello world")
		fmt.Println("z修改后的value为：", reflect.ValueOf(z))
	} else {
		fmt.Println("该变量不能修改")
	}

	if reu.CanSet() {
		reu.SetBool(false)
		fmt.Println("u修改后的value为：", reflect.ValueOf(u))
	} else {
		fmt.Println("该变量不能修改")
	}

	fmt.Println("x Elem value为：", rex)
	fmt.Println("y Elem value为：", rey)
	fmt.Println("z Elem value为：", rez)
	fmt.Println("u Elem value为：", reu)

	//判断是否可以修改变量x的值，若可以，则用Set()方法进行修改
	if rex.CanSet() {
		ax := reflect.ValueOf(2344.23466) // 使用Set方法修改值，Set方法接收的是ValueOf的返回值
		rex.Set(ax)
		fmt.Println("x Set修改后的value为：", reflect.ValueOf(x))
	} else {
		fmt.Println("该变量不能修改")
	}

	if rey.CanSet() {
		ay := reflect.ValueOf(44444) // 使用Set方法修改值，Set方法接收的是ValueOf的返回值
		rey.Set(ay)
		fmt.Println("y Set修改后的value为：", reflect.ValueOf(y))
	} else {
		fmt.Println("该变量不能修改")
	}

	if rez.CanSet() {
		az := reflect.ValueOf("hello world44444") // 使用Set方法修改值，Set方法接收的是ValueOf的返回值
		rez.Set(az)
		fmt.Println("z Set修改后的value为：", reflect.ValueOf(z))
	} else {
		fmt.Println("该变量不能修改")
	}

	if reu.CanSet() {
		au := reflect.ValueOf(true) // 使用Set方法修改值，Set方法接收的是ValueOf的返回值
		reu.Set(au)
		fmt.Println("u Set修改后的value为：", reflect.ValueOf(u))
	} else {
		fmt.Println("该变量不能修改")
	}

}
