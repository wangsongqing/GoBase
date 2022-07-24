package controllers

import (
	"encoding/json"
	"fmt"
)

type Man1 struct {
	Name string      `json:"name"`
	Age  int         `json:"age"`
	Addr interface{} `json:"addr"`
}

func JsonTest() {
	// JsonTest1()
	// JsonTest2()
	// JsonTest3()
	//JsonTest4()
	// JsonTest5()
	JsonTest6()
}

// JsonTest6 反序列化 slice
func JsonTest6() {
	var man1 []string
	str := "[\"林冲\",\"杨志\"]"
	err := json.Unmarshal([]byte(str), &man1)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}

	for _, v := range man1 {
		fmt.Println(v)
	}

}

// JsonTest5 反系列化map
func JsonTest5() {
	var map1 map[string]interface{}
	str := "{\"0\":{\"age\":\"42\",\"name\":\"宋江\"},\"1\":{\"age\":\"21\",\"name\":\"无用\"}}"
	err := json.Unmarshal([]byte(str), &map1)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	fmt.Println(map1)
}

// JsonTest4 反序列化结构体
func JsonTest4() {
	var man1 Man1
	str := "{\"Name\":\"武松\",\"Age\":30,\"Addr\":{\"大朗\":\"清河\",\"潘金莲\":\"阳光县\"}}"
	err := json.Unmarshal([]byte(str), &man1)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}

	fmt.Println(man1.Addr)
}

// JsonTest1 结构体序列化json
func JsonTest1() {
	Man1 := Man1{
		Name: "武松",
		Age:  30,
		Addr: map[string]string{"addr1": "清河", "addr2": "阳光县"}, // 切片的定义
	}

	data, err := json.Marshal(Man1)
	if err != nil {
		fmt.Println("序列化失败")
	}

	fmt.Println(string(data))
}

// JsonTest2 序列化map
func JsonTest2() {

	var a1 map[string]string
	a1 = make(map[string]string, 1)
	a1["name"] = "宋江"
	a1["age"] = "42"

	a2 := map[string]string{"name": "无用", "age": "21"}

	var man map[int]interface{}
	man = make(map[int]interface{})
	man[0] = a1
	man[1] = a2
	data, _ := json.Marshal(man)
	fmt.Println(string(data))
}

// JsonTest3 切片序列化
func JsonTest3() {

	shuihu := []string{"林冲", "杨志"}
	data, _ := json.Marshal(shuihu)
	fmt.Println(string(data))
}
