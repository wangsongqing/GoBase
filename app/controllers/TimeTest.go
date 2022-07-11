package controllers

import (
	"fmt"
	"time"
)

func TimeTest() {
	// time.Millisecond * 1000 = 1s
	// getTime()
	// getNowTime()
	// sleepTimeData()
	RandRoom()
}

func getNowTime() {
	now := time.Now()
	fmt.Printf("年:%v \n", now.Year())
	fmt.Printf("月:%v \n", int(now.Month()))
	fmt.Printf("日:%v \n", now.Day())
	fmt.Printf("时:%v \n", now.Hour())
	fmt.Printf("分:%v \n", now.Minute())
	fmt.Printf("秒:%v \n", now.Second())
	// nowTime := strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "-" + strconv.Itoa(now.Day()) + " " + strconv.Itoa(now.Hour()) + ":" + strconv.Itoa(now.Minute()) + ":" + strconv.Itoa(now.Second())
	nowTime := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("Date:%v \n", nowTime)
	nowTime1 := fmt.Sprintf("date:%v", now.Format("2006-01-02 15:04:05"))
	fmt.Println(nowTime1)
}

func getTime() {
	for {
		var now = time.Now()
		fmt.Printf("time:%v, type:%T \n", now, now)
		time.Sleep(time.Millisecond * 1000)
	}
}

func sleepTimeData() {
	// time.Second = 1s
	i := 0
	for {
		i++
		if i == 10 {
			break
		}
		fmt.Println(i)
		time.Sleep(time.Second * 2)
	}
}

// RandRoom 随机数
func RandRoom() {
	//now := time.Now()
	initTime := time.Now().UnixMilli()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	initTime1 := time.Now().UnixMilli()

	fmt.Printf("耗费时间:%vs", initTime1-initTime)
}
