package controllers

import (
	"fmt"
)

type USB interface {
	start(types string)
	stop(types string)
}

type Worker struct {
}

func (u *Worker) start(types string) {
	fmt.Println("computer start", types)
}

func (u *Worker) stop(types string) {
	fmt.Println("computer stop", types)
}

type Monkey struct {
	Name string
}

type Birder interface {
	Flying()
}

func (m Monkey) climbing() {
	fmt.Println(m.Name, "会爬树")
}

type LitMonkey struct {
	Monkey // 继承
}

func (l LitMonkey) Flying() {
	fmt.Println(l.Name, "通过学习会飞翔")
}

func InterfaceTest() {
	// InterfaceTest1()
	// InterfaceTest2()
	// InterfaceTest3()

	a := "类型"
	b := 34
	c := 11.9
	InterfaceTest4(a, b, c)
}

// InterfaceTest4 类型断言-判断传入变量的内容
func InterfaceTest4(items ...interface{}) {

	for t, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("bool param:%d, value:%v \n", t, x)
		case int64, int, int32:
			fmt.Printf("int param:%d, value:%v \n", t, x)
		case string:
			fmt.Printf("string param:%d, value:%v \n", t, x)
		case float64:
			fmt.Printf("float64 param:%d, value:%v \n", t, x)
		default:
			fmt.Println("无效的类型")
		}
	}

}

// InterfaceTest3 类型断言
func InterfaceTest3() {
	var a interface{}
	b := 34.9
	a = b //空接口可以接受任意类型

	// y, ok := a.(float64) // 类型断言
	if y, ok := a.(float64); ok {
		fmt.Println(y)
	} else {
		fmt.Println("断言失败")
	}

	fmt.Println(a)
}

// InterfaceTest2 继承和接口结合使用
func InterfaceTest2() {
	LitMonkey := LitMonkey{
		Monkey{
			Name: "猴子1",
		},
	}
	// LitMonkey.Name = "猴子"
	LitMonkey.climbing()
	LitMonkey.Flying()
}

func InterfaceTest1() {
	var computer USB = &Worker{}
	computer.stop("computer")
	computer.start("computer")

	var worker Worker
	var a1 USB = &worker
	a1.stop("phone")
	a1.start("phone")
}
