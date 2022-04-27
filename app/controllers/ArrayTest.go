package controllers

import "fmt"

func RunArray() {

	// state1()
	// state2()
	state3()

}

// 定义数组1
func state1() {
	var arr [3]int
	arr[0] = 11
	arr[1] = 22
	arr[2] = 33

	for i, v := range arr {
		fmt.Printf("key:%v, value:%v \n", i, v)
	}
}

// 定义数组2
func state2() {
	arr := [...]string{"name", "age"}
	for i, v := range arr {
		fmt.Printf("key:%v, value:%v \n", i, v)
	}
}

// 定义数组3
func state3() {
	arr := [2]int{66, 88}
	for k, v := range arr {
		fmt.Printf("key:%v, value:%v \n", k, v)
	}
}
