package controllers

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
}

type Person struct {
	Name  string
	Age   int
	Score [2]float32
	ptr   *int
	slice []int
	map1  map[int]string
}

func StructTest() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	test6()
}

// 结构体转成JSON
func test6() {
	cat1 := Cat{"牛", 12, 23.00}
	jsonStr, err := json.Marshal(cat1)
	if err != nil {
		fmt.Println("转JSON失败")
	}
	fmt.Println(string(jsonStr))
}

func test5() {
	var c1 Cat
	c1.Name = "小花"

	// var c2 *Cat = &c1 和 c2 := &c1 效果一样
	c2 := &c1      // c2的值类型本身就是指针了
	c2.Name = "嘿嘿" // 因为指针修改c2会导致c1的Name发生变化

	fmt.Println(c1)
	fmt.Printf("c1地址:%p, c2地址:%p", &c1, c2)
}

// 结构体赋值方式
func test3() {
	// cat1 := Cat{"wsq", 11, 2.00}
	cat1 := Cat{}
	cat1.Name = "小马"
	fmt.Println(cat1)
}

func test4() {
	cat := &Cat{}
	// var cat1 = &Cat{}
	(*cat).Name = "小黑"
	// cat.Name = "小白"
	cat.Age = 10

	fmt.Println(cat)
}

func test2() {
	var Person1 Person
	Person1.Name = "wsq"
	Person1.Age = 20

	Person1.Score[0] = 100
	Person1.Score[1] = 99

	// Person1.ptr=&Person1

	Person1.slice = []int{11, 22}
	Person1.map1 = map[int]string{0: "w", 1: "s"}
	fmt.Println(Person1)
}

func test1() {
	var cat1 Cat
	cat1.Name = "小花"
	cat1.Age = 3
	fmt.Printf("Cat的地址:%p \n", &cat1)
	fmt.Printf("Cat Name 的地址:%p \n", &cat1.Name)
	fmt.Printf("Cat Nme is %s,  Age is %d", cat1.Name, cat1.Age)
}
