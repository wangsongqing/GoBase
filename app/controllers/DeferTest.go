package controllers

import "fmt"

func DeferTest() {
	// defer1 1
	// defer2 2
	//DeferTest1()

	i := DeferTest2()
	fmt.Println(*i)
}

func DeferTest1() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i //或者直接写成return
}

func DeferTest2() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}
