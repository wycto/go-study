package main

func main() {
	/*var x float64 = 3.4
	reflect_type(x)*/

	/*var x float64 = 3.4
	reflect_value(x)*/

	/*var x float64 = 3.4
	// 反射认为下面是指针类型，不是float类型
	reflect_set_value(&x)
	fmt.Println("main:", x)*/

	u := User{1, "zs", 20}
	Poni(u)
}
