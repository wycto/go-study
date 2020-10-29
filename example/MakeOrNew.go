package example

import (
	"fmt"
	"reflect"
)

type MakeOrNew struct {
}

func (receiver *MakeOrNew) MakeMap() {
	fmt.Println("\nMakeOrNew.MakeMap()")
	makeMap := make(map[string]string)
	makeMap["name"] = "魏义"
	fmt.Println("makeMap:", makeMap)
	fmt.Println("makeMapTypeOf:", reflect.TypeOf(makeMap))
}

func (receiver *MakeOrNew) NewMap() {
	fmt.Println("\nMakeOrNew.NewMap()")
	newMap := new(map[string]string)
	fmt.Println("makeMap:", newMap)
	fmt.Println("makeMapTypeOf:", reflect.TypeOf(newMap))
}

func (receiver MakeOrNew) MakeSlice() {
	fmt.Println("\nMakeOrNew.MakeSlice()")
	makeSlice := make([]string, 10)
	makeSlice[0] = "张三"
	makeSlice[1] = "李四"
	makeSlice[2] = "王二"
	makeSlice[3] = "麻子"
	makeSlice = append(makeSlice, "里斯")
	fmt.Println("makeSlice:", makeSlice)
	fmt.Println("makeSlice[10]:", makeSlice[10])
	fmt.Println("cap(makeSlice):", cap(makeSlice))
}

func (receiver MakeOrNew) CopySlice() {
	fmt.Println("\n\nMakeOrNew.CopySlice():")
	oneSlice := make([]string, 3)
	oneSlice[0] = "张三"
	oneSlice[1] = "李四"
	oneSlice[2] = "王二"
	fmt.Println("oneSlice:", oneSlice)

	twoSlice := make([]string, 5)
	twoSlice[2] = "魏义"
	twoSlice[3] = "黄冬"
	fmt.Println("twoSlice:", twoSlice)

	copy(oneSlice, twoSlice)
	fmt.Println("copy(oneSlice,twoSlice):oneSlice:", oneSlice)

	/*copy(twoSlice,oneSlice)
	fmt.Println("copy(twoSlice,oneSlice):twoSlice:",twoSlice)*/
}

func (receiver MakeOrNew) DeleteSlice() {
	fmt.Println("\n\nMakeOrNew.DeleteSlice():")
	oneSlice := make(map[string]string, 3)
	oneSlice["a"] = "张三"
	oneSlice["b"] = "李四"
	oneSlice["c"] = "王二"
	fmt.Println("oneSlice:", oneSlice)

	twoSlice := make(map[int]string, 5)
	twoSlice[2] = "魏义"
	twoSlice[3] = "黄冬"
	fmt.Println("twoSlice:", twoSlice)

	delete(oneSlice, "b")
	fmt.Println("delete(oneSlice,\"b\"):oneSlice:", oneSlice)

	delete(twoSlice, 3)
	fmt.Println("delete(twoSlice,3):twoSlice:", twoSlice)
}
