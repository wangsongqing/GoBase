package controllers

import (
	"GoBase/app/models"
	"fmt"
)

// 面向对象编程

type Student struct {
	Name   string
	Gender string
	Age    int
	Id     int
	Score  float64
}

type Visitor struct {
	Name  string
	Age   int
	Money float64
}

func (s *Student) say() string {
	return fmt.Sprintf("student 信息, name:%v, gender:%v, age:%d, id:%d, score:%v", s.Name, s.Gender, s.Age, s.Id, s.Score)
}

func (v *Visitor) VisitorSend() (string, int) {
	if v.Age > 18 && v.Age < 60 {
		return "成人", 20
	} else if v.Age > 60 {
		return "老年人", 0
	}
	return "未成年", 0
}

func OopTest() {
	// testoop1()
	// testoop2()
	// testoop3()
	// testoop4()
	// testoop5()
	testoop6()
}

// 继承
func testoop6() {
	user := models.SmallStudent{}
	user.Name = "jimmy"
	fmt.Println(user)
}

// 封装
func testoop5() {
	user := models.NewPerson("jimmy")
	user.SetSla(8000)
	user.SetAge(18)
	fmt.Printf("name:%v, age:%d, sla:%v", user.Name, user.GetAge(), user.GetSla())
}

// 面向对象编程
func testoop4() {
	account := models.Account{
		AccountNo: "12344321",
		Password:  "a123123a",
		Balance:   100.00,
	}

	var accountType int
	var accountMoney float64
	var accountPwd string
	fmt.Println("请输入操作类型")
	fmt.Scanln(&accountType)

	fmt.Println("请输入操作金额")
	fmt.Scanln(&accountMoney)

	fmt.Println("请输入账户密码")
	fmt.Scanln(&accountPwd)

	switch accountType {
	case 1:
		// 存款
		account.SaveMoney(accountMoney, accountPwd)
		fmt.Println("余额:", account.Balance)
	case 2:
		account.Withdrew(accountMoney, accountPwd)
		fmt.Println("余额:", account.Balance)
	default:
		fmt.Println("操作错误")
	}

}

// 工厂模式调用studentMan 因为首字母是小写，对于包来说方法就是私有的
func testoop3() {
	StudentMan := models.NewStudent(
		"MM",
		88,
	)

	fmt.Println(StudentMan)
	fmt.Printf("name:%v, age:%d \n", StudentMan.Name, StudentMan.Age)
	fmt.Println(StudentMan.GetScore())
}

func testoop2() {
	Visitor1 := Visitor{}
	for {
		fmt.Println("请输入名字:")
		fmt.Scanln(&Visitor1.Name)
		if Visitor1.Name == "n" {
			break
		}

		fmt.Println("请输入年龄:")
		fmt.Scanln(&Visitor1.Age)
		tags, m := Visitor1.VisitorSend()
		fmt.Printf("%v是%v，所以收费%d \n \n", Visitor1.Name, tags, m)
	}
	// Visitor1.Name = "小马"
	// Visitor1.Age = 66
	//tags, m := Visitor1.VisitorSend()
	//fmt.Printf("%v是%v，所以收费%d", Visitor1.Name, tags, m)
}

func testoop1() {
	Student1 := Student{
		"wsq",
		"语文",
		20,
		1,
		0.0,
	}
	Student1.Score = 88.00
	studentStr := Student1.say()
	fmt.Println(studentStr)
}
