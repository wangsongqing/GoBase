package controllers

import (
	"fmt"
)

type Stu struct {
	Name    string
	Age     int
	Address string
}

// MapSlice map类型是引用类型
func MapSlice() {
	// stateMap1()
	// stateMap2()
	// statusMap3()
	// statusMap4()
	// statusMap5()
	// statusMap6()
	statusMap7()
}

// 定义map1
func stateMap1() {
	var map1 map[int]string
	map1 = make(map[int]string, 1) //出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。
	map1[1] = "11"
	map1[2] = "521"
	map1[3] = "250"
	fmt.Println("map1", map1)
}

// 定义map2
func stateMap2() {
	map1 := map[int]string{1: "a", 2: "b"}
	fmt.Println("map1", map1)
}

// for in 遍历map
func statusMap3() {
	map1 := map[int]string{1: "AA", 2: "BB"}
	for k, v := range map1 {
		fmt.Printf("k is %d, v is %s \n", k, v)
	}
}

// map赋值一定需要make,不然会出现异常
func statusMap4() {
	map1 := make(map[int]int, 10)
	map2 := make(map[int]int, 10)

	map2[1] = 11
	map2[2] = 22

	var num int
	for k, v := range map2 {
		map1[k] = v
		num += v
	}

	fmt.Printf("sum: %d \n", num)
	fmt.Println("map1", map1)
}

func statusMap5() {
	student := make(map[string]map[string]string)
	student["01"] = map[string]string{"name": "wsq", "sex": "男"}
	student["02"] = map[string]string{"name": "mak", "sex": "未知"}
	fmt.Println("student", student["02"])

	// 删除student[01]
	delete(student, "02")

	fmt.Println("student delete", student)

	fmt.Println("map的长度", len(student))
}

// map切片
func statusMap6() {
	var mapSlice []map[string]string
	mapSlice = make([]map[string]string, 2)

	if mapSlice[0] == nil {
		mapSlice[0] = make(map[string]string)
		mapSlice[0]["name"] = "牛魔王"
		mapSlice[0]["age"] = "1000"
	}

	if mapSlice[1] == nil {
		mapSlice[1] = make(map[string]string)
		mapSlice[1]["name"] = "猪八戒"
		mapSlice[1]["age"] = "300"
	}

	for k, v := range mapSlice {
		fmt.Printf("k:%d, v:%v \n", k, v)
	}

}

// map排序
func statusMap7() {
	//map1 := map[]
}
