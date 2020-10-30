package json

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"` //字段映射Name name
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func (receiver *User) JsonEncodeStruct() {
	receiver.Name = "JsonEncodeStruct"
	receiver.Age = 30
	receiver.Sex = "男"

	jsonData, _ := json.Marshal(receiver)
	fmt.Println("json data: ", string(jsonData))
}

func (receiver *User) JsonEncodeMap() {
	myMap := make(map[string]interface{}) //interface代表任意类型，随便什么数据类型
	myMap["Name"] = "JsonEncodeMap"
	myMap["Age"] = 30
	myMap["Sex"] = "男"
	myMap["User"] = receiver

	jsonData, _ := json.Marshal(myMap)
	fmt.Println("json data: ", string(jsonData))
}
