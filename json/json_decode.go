package json

import (
	"encoding/json"
	"fmt"
)

type JsonDecode struct {
	Name string `json:"name"` //字段映射Name name
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

/**
反序列化结构体
*/
func (receiver *JsonDecode) JsonDecodeStruct() {
	jsonStr := "{\"name\":\"JsonEncodeStruct\",\"age\":30,\"sex\":\"男\"}"

	err := json.Unmarshal([]byte(jsonStr), &receiver)
	if err != nil {
		fmt.Println("json decode err: ", err.Error())
	}
	fmt.Println("json decode data: ", receiver)
}

/**
反序列化，反序列化字符串，首先要有一个接收的东西，比如发系列化map，要有一个接收的map
*/
func (receiver *JsonDecode) JsonDecodeMap() {
	jsonStr2 := "{\"Age\":30,\"Name\":\"JsonEncodeMap\",\"Sex\":\"男\",\"User\":{\"name\":\"JsonEncodeStruct\",\"age\":30,\"sex\":\"男\"}}"
	myMap := make(map[string]interface{}) //interface代表任意类型，随便什么数据类型

	err := json.Unmarshal([]byte(jsonStr2), &myMap)
	if err != nil {
		fmt.Println("json decode err: ", err.Error())
	}
	fmt.Println("json decode data: ", myMap)
}
