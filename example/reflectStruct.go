package example

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name  string
	Age   int
	Sex   string
	IsCan bool
}

func ReflectStructMain() {
	s1 := Stu{Name: "王一", Age: 18, Sex: "男", IsCan: false}
	s2 := Stu{Name: "王二", Age: 19, Sex: "女", IsCan: true}
	s3 := Stu{Name: "张三", Age: 20, Sex: "男", IsCan: false}

	//反射获取结构体的类型和值
	fmt.Println("s1的类型", reflect.TypeOf(s1))
	fmt.Println("s1的值", reflect.ValueOf(s1))

	fmt.Println("s2的类型", reflect.TypeOf(s2))
	fmt.Println("s2的值", reflect.ValueOf(s2))

	fmt.Println("s3的类型", reflect.TypeOf(s3))
	fmt.Println("s3的值", reflect.ValueOf(s3))

	fmt.Println("TypeOf()和Kind()方法输出的区别")
	fmt.Println("TypeOf(s1)：", reflect.TypeOf(s1))
	s1tp := reflect.TypeOf(s1)
	fmt.Println("Kind(s1)：", s1tp.Kind())

	/**
	func TypeOf(i interface{}) Type{
		eface := *(*emptyInterface)(unsafe.Pointer(&i))
		return toType(eface.type)
	}

	//Type接口提供了一系列方法
		type Type interface{
			Align()  int                      //对齐边界
			FieldAlign()  int
			Method(int) Method
			MethodByName(string)  (Method, bool)       //方法
			NumMethod()  int         //类型名称
			Name()  string
			PkgPath()  string         //包路径
			Slize()  uintptr
			String()  string
			Kind()  Kind
			Implements(u Type)  bool         //是否实现指定接口
			AssginableTo(u Type)  bool
			ConvertibleTo(u Type)  bool
			Comparable()  bool                //是否可比较
		}
		**/

}
