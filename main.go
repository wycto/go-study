package main

import (
	"fmt"
	"study/test/math"
)

func main() {
	c := math.Add(12, 23)
	fmt.Println(c)
	user := math.User{}
	user.SetName("长位")
	fmt.Println(user.Name)
	fmt.Println("____________________")
	fmt.Println(math.GetN())
}
