package controllers

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var Mymap = make(map[int]int, 10)

// GoroutineTest go 协程
func GoroutineTest() {
	// go GoroutineTest1()
	// GoroutineTest2()
	// GoroutineTest3()
	// GoroutineTest5()
	GoroutineTest6()
}

func GoroutineTest1() {
	for i := 1; i <= 10; i++ {
		fmt.Println("main hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func GoroutineTest2() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func GoroutineTest3() {
	cpuNum := runtime.NumCPU() // 机器CPU的数量

	runtime.GOMAXPROCS(cpuNum - 1) // 使用CPU的数量

	fmt.Printf("cpu num  is %d", cpuNum)
}

// GoroutineTest4 用channel 解决map死锁问题
func GoroutineTest4(n int, Chan1 chan map[int]int) {
	map1 := make(map[int]int, 10)
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	map1[n] = res
	Chan1 <- map1
}

// GoroutineTest5 会报错，会出现死锁
//func GoroutineTest5() {
//	for i := 1; i <= 200; i++ {
//		go GoroutineTest4(i)
//	}
//
//	time.Sleep(time.Second * 4)
//
//	for i, v := range Mymap {
//		fmt.Printf("i:%d, v:%d \n", i, v)
//	}
//}

func GoroutineTest6() {
	var Chan1 = make(chan map[int]int, 200)

	for i := 1; i <= 200; i++ {
		go GoroutineTest4(i, Chan1)
	}

	for {
		if len(Chan1) >= 200 {
			close(Chan1)
			break
		}
	}

	for v := range Chan1 {
		fmt.Printf("value:%d \n", v)
	}
}
