package main

import "study/json"

func main() {
	/*c := math.Add(12, 23)
	fmt.Println(c)
	user := math.User{}
	user.SetName("长位")
	fmt.Println(user.Name)
	fmt.Println("____________________")
	fmt.Println(math.GetN())*/

	//example MakeOrNew
	//makeOrNew := example.MakeOrNew{}
	//makeOrNew.MakeMap()
	//makeOrNew.NewMap()
	//makeOrNew.MakeSlice()
	//makeOrNew.CopySlice()
	//makeOrNew.DeleteSlice()

	/*example: panic and recover*/
	//panicAndRecover := example.PanicAndRecover{}
	/*panicAndRecover.ReceivePanic()
	panicAndRecover.RecoverPanic()*/
	//panicAndRecover.RecoverPanicTest()

	/*example: json*/
	user := json.User{}
	user.JsonEncodeStruct()
	user.JsonEncodeMap()

	jsonDecode := json.JsonDecode{}
	jsonDecode.JsonDecodeStruct()
	jsonDecode.JsonDecodeMap()
}
