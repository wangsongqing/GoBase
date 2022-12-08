package controllers

import (
	"GoBase/app/pkg/helpers"
	"fmt"
	_ "sort"
)

func RunSlice() {
	// x := []int{5, 6, 7}
	// stateSlice1(x)
	// stateSlice2()
	// stateSlice3()
	//stateSlice4()
	//stateSlice5()
	//stateSlice6()
	//stateSlice7()
	//stateSlice8()
	//stateSlice9()
	//stateSlice10()
	//stateSlice11()
	//stateSlice12()
	stateSlice13()
	//stateSlice14()
}

func stateSlice1(x []int) {
	var sum int
	for _, v := range x {
		sum += v
	}

	fmt.Printf("和为: sum %d", sum)
}

// 定义切片方式1   通过让切片去引用一个数组定义切片
func stateSlice2() {
	var intArr = [5]int{11, 22, 33, 44, 55}
	var slice = intArr[1:3]
	fmt.Println("数组的元素1: ", intArr)
	fmt.Println("元素: ", slice)
	fmt.Println("长度: ", len(slice))
	fmt.Println("容量: ", cap(slice)) //切片的容量是可以变化的
	fmt.Println("intArr[22]的地址是:", &intArr[1])
	fmt.Printf("slice[0]的地址是: %p \n", &slice[0])
	slice[0] = 888
	fmt.Println("数组的元素2: ", intArr)
}

// 定义切片方式2 make方式可以指定切片的大小和容量
func stateSlice3() {
	var slice = make([]int, 2, 2)
	slice[0] = 11
	slice[1] = 22
	// slice[2] = 33  超出范围了
	fmt.Println("slice 的元素:", slice)
	fmt.Println("slice 的长度为:", len(slice))
	fmt.Println("slice 的容量为:", cap(slice))
	for k, v := range slice {
		fmt.Printf("k: %d, v: %d \n", k, v)
	}
}

// 定义切片方式3
func stateSlice4() {
	var slice = []string{"tom", "jimmy", "linda"}
	fmt.Println("slice:", slice)
	fmt.Printf("slice的大小: %d \n", len(slice))
	fmt.Printf("slice的容量: %d", cap(slice))
	fmt.Println("slice", slice)

}

// 切片的遍历
func stateSlice5() {
	var slice = []string{"tom", "jimmy", "linda"}
	slice = append(slice, "jack") // 切片追加数据
	for i := 0; i < len(slice); i++ {
		fmt.Printf("i=%d, v=%s \n", i, slice[i])
	}

	fmt.Println("-------------------------")

	for k, v := range slice {
		fmt.Printf("i=%d, v=%s \n", k, v)
	}
}

func stateSlice6() {
	var slice = []string{"tom", "jimmy", "linda"}

	slice2 := slice[0:1] //slice2和slice的地址是一样的，改变slice1   slice也会相应的改变

	slice2[0] = "jack"
	fmt.Println("slice2:", slice2)
	fmt.Println("slice1:", slice) // slice1: [jack jimmy linda]
}

// 拷贝切片 copy
func stateSlice7() {
	slice := []int{11, 22, 33}
	slice1 := []int{44, 55, 66}

	copy(slice, slice1)
	fmt.Println("slice=", slice)
	fmt.Println("slice1=", slice1)
}

// 追加切片
func stateSlice8() {
	slice := []int{11, 22, 33}
	slice = append(slice, 44)

	fmt.Println("slice=", slice)
}

// 字符串也可以切片截取
func stateSlice9() {
	str := "wangsong@gmail.com"
	str = str[5:]
	fmt.Println("str", str)
}

// 修改字符串某一些值
func stateSlice10() {
	str := "abcd"
	arr := []rune(str)
	arr[0] = '北'
	arr[1] = 'B'
	str = string(arr)

	fmt.Println("str", str)
}

// 数组查找
func stateSlice12() {
	arr := []int{11, 22, 44, 6}
	intKey := helpers.BinaryFind(arr, 111)
	fmt.Println("intKey", intKey)
}

func stateSlice13() {
	a := make([]interface{}, 2)
	a[0] = 111
	a[1] = "222"
	slice1 := append(a, "333", "444")
	fmt.Printf("a[0] type is %T \n", a[0])
	fmt.Printf("a[1] type is %T \n", a[1])
	fmt.Println(slice1)
}

func stateSlice14() {
	slice := make([]string, 2)
	slice[0] = "111"
	fmt.Println(slice)
}
