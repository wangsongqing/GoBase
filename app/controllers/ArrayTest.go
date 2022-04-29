package controllers

import "fmt"

func RunArray() {

	// state1()
	// state2()
	// state3()

	// arr := [3]int{1, 4, 6}
	// state4(&arr)
	state5()
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

// 数组传值最好传地址，因为传值会占用大量内存
func state4(arr *[3]int) {
	var num int
	for i := 0; i < len(arr); i++ {
		num += arr[i]
	}

	fmt.Printf("和为:%v", num)
}

func state5() {
	var arr [26]byte
	for i := 0; i < len(arr); i++ {
		arr[i] = 'A' + byte(i)
	}

	fmt.Println(arr)
}
