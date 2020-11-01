package point

import "fmt"

func TestPoint() {
	var count int = 20
	var countPoint *int //这里的*是指针
	countPoint = &count
	var name *int

	fmt.Println("count 变量: ", count)
	fmt.Println("count 变量的地址: ", &count)
	fmt.Println("countPoint 变量存储的地址：", countPoint)
	fmt.Println("countPoint 指针指向地址的值：", *countPoint) //这里的*是取值运算符
	fmt.Println("&countPoint：", &countPoint)         //这里的*是取值运算符
	fmt.Println("&*countPoint：", &*countPoint)       //这里的*是取值运算符
	fmt.Println("*&countPoint：", *&countPoint)       //这里的*是取值运算符
	fmt.Println("name 变量存储的地址：", name)
}

/**
数组指针
*/
func ArrPoint() {
	var a, b, c = 2, 3, 5
	arrPoint := [...]*int{&a, &b, &c}
	fmt.Println("数组指针：", arrPoint)
}

/**
指针数组
*/
func PointArr() {
	pointArr := [...]int{4, 6, 9}
	fmt.Println("指针数组：", &pointArr)
}

func init() {
	fmt.Println("this is point point.go init")
}
