package main

/*
指针学习篇
*/
import (
	"fmt"
	"os/exec"
	"time"
)

// "study/example"

func main() {

	var secondTime time.Duration;
	var shFileName string;

	fmt.Println("定时任务：")

	fmt.Println("请输入定时器时间间隔，秒数：")
	fmt.Scanln(&secondTime)

	fmt.Println("请输入sh命令文件名称：")
	fmt.Scanln(&shFileName)

	fmt.Println(time.Second)

	ticker := time.NewTicker(secondTime * time.Second)
	for _ = range ticker.C {
		//fmt.Println(time.Now())
		cmd := exec.Command("sh",shFileName)
		fmt.Println(cmd)
		fmt.Println(cmd.Run())
	}
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
	// panicAndRecover := example.PanicAndRecover{}
	/*panicAndRecover.ReceivePanic()
	panicAndRecover.RecoverPanic()*/
	// panicAndRecover.RecoverPanicTest()

	/*example: json*/
	/*user := json.User{}
	user.JsonEncodeStruct()
	user.JsonEncodeMap()

	jsonDecode := json.JsonDecode{}
	jsonDecode.JsonDecodeStruct()
	jsonDecode.JsonDecodeMap()*/

	/*example: point*/
	/* point.TestPoint()
	point.ArrPoint()
	point.PointArr() */
}