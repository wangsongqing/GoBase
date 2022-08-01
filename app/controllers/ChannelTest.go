package controllers

import (
	"fmt"
)

type Cat1 struct {
	Name string
	Age  int
}

func ChannelTest() {
	// ChannelTest1()
	// ChannelTest2()
	// ChannelTest3()
	ChannelTest4()
}

func ChannelWrite(Channles1 chan int) {
	for i := 1; i <= 50; i++ {
		Channles1 <- i
		fmt.Println("写入数据：", i)
		// time.Sleep(time.Second)
	}
	close(Channles1)
}

func ChannelRead(Channle1 chan int, Channle2 chan bool) {
	for {
		v, ok := <-Channle1
		if !ok {
			break
		}
		fmt.Println("读取数据", v)
		// time.Sleep(time.Second)
	}

	Channle2 <- true
	close(Channle2)
}

func ChannelTest4() {
	Channle1 := make(chan int, 100)
	Channle2 := make(chan bool, 1)
	go ChannelWrite(Channle1)
	go ChannelRead(Channle1, Channle2)
	<-Channle2

	// time.Sleep(time.Second * 10)

	//for {
	//	_, ok := <-Channle2
	//	if !ok {
	//		break
	//	}
	//}
}

// ChannelTest3 遍历管道
func ChannelTest3() {
	chan1 := make(chan int, 100)
	for i := 1; i < 100; i++ {
		chan1 <- i * 2
	}

	close(chan1) // 关闭管道

	for v := range chan1 {
		fmt.Printf("值为: %d \n", v)
	}

}

func ChannelTest2() {
	allChan := make(chan interface{}, 4)
	allChan <- 10
	allChan <- "abc"
	cat2 := Cat1{
		Name: "小猫咪",
		Age:  5,
	}
	allChan <- cat2

	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Printf("allChan的类型是:%T, allChan的值是:%v \n", newCat, newCat)
	// fmt.Printf("newCat的名字是: %v", newCat.Name) // 编译不通过

	//类型断言
	cat3 := newCat.(Cat1)
	fmt.Printf("newCat的名字是: %v", cat3.Name)
}

func ChannelTest1() {

	var ChanInt chan int        // 声明channel
	ChanInt = make(chan int, 3) // 定义管道，容量最多加3个，超过会有异常
	ChanInt <- 10               // 向管道写入数据
	num := 11
	ChanInt <- num
	fmt.Printf("ChanInt的值:%v， ChanInt的地址:%p \n", ChanInt, &ChanInt)
	fmt.Printf("ChanInt的长度:%v， ChanInt的容量:%v \n", len(ChanInt), cap(ChanInt))

	num2 := <-ChanInt // 从管道取数据
	fmt.Printf("从管道取数据：%v \n", num2)
	fmt.Printf("ChanInt的长度:%v， ChanInt的容量:%v \n", len(ChanInt), cap(ChanInt))
}
