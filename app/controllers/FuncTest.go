package controllers

import "fmt"

func FuncTest() {
	//FuncTest1()
	FuncTest3()
}

func FuncTest1() {
	name := "addr"
	FuncTest2(&name)
	fmt.Println(name)
}

func FuncTest2(name *string) {
	*name += "wsq"
}

func FuncTest3() {
	f := func(x, y int) int {
		return x + y
	}(1, 22)

	fmt.Println(f)
}
