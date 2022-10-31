package controllers

import (
	"fmt"
	"strconv"
	"strings"
)

func StringTest() {

	//strs := "hello123哈喽"
	//
	//// strs := []rune(str)
	//for i := 0; i < len(strs); i++ {
	//	fmt.Println(string(strs[i]))
	//}

	// StringToNumber()
	// HasPrefix()
	// stringToByte()
	Contains()
}

func Contains() {
	str := "abcnubxc"
	fmt.Println(strings.Contains(str, "bx1"))
}

func stringToByte() {
	str := "times"
	list := []byte(str)
	fmt.Printf("type：%T, value: %v", list, list)
}

func HasPrefix() {
	str := "abcdft"
	fmt.Println(strings.HasPrefix(str, "abc1"))
}

func StringToNumber() {
	num := "123汉字"
	v, err := strconv.Atoi(num) // 字符串转int
	if err != nil {
		fmt.Println("类型错误")
		panic(err)
		return
	}
	fmt.Printf("type:%T, value:%v", v, v)
}
