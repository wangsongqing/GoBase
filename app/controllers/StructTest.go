package controllers

import (
	"fmt"
)

type Cat struct {
	Name  string
	Age   int
	Score float32
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
	test3()
}

// 结构体赋值方式
func test3() {
	// cat1 := Cat{"wsq", 11, 2.00}
	cat1 := Cat{}
	cat1.Name = "小马"
	fmt.Println(cat1)
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
