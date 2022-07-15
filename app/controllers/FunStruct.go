package controllers

import (
	"fmt"
	"strconv"
)

type Persons struct {
	Name string
	Age  int
}

type Opretore struct {
	Num1 float64
	Num2 float64
	Op   byte
}

func FunStructTest() {
	// test1()
	// FunStructTest2()
	op := Opretore{}
	res := op.FunStructTest3(10, 3.659, '*')
	fmt.Println(res)
}

func (o *Opretore) FunStructTest3(num1 float64, num2 float64, op byte) float64 {

	res := 0.0
	switch op {
	case '+':
		res = num1 + num2
	case '-':
		res = num1 - num2
	case '*':
		res = num1 * num2
	case '/':
		res = num1 / num2
	default:
		return res
	}

	result, err := strconv.ParseFloat(fmt.Sprintf("%.2f", res), 64)
	if err != nil {
		return res
	}

	return result
}

func (p *Persons) test1() {
	p.Name = "小黑"
	fmt.Printf("test1:%v \n", p.Name)
}

func (p Persons) test11() string {
	return fmt.Sprintf("this man name is %v, age %d", p.Name, p.Age)
}

func FunStructTest2() {
	p1 := Persons{"小王", 20}
	p1.test1()

	p2 := Persons{}
	p2.Age = 88
	p2.Name = "aaaaa"
	userMsg := p2.test11()
	fmt.Println(userMsg)

	var p3 Persons
	p3.Age = 22
	p3.Name = "发发"
	p3.test1()
}
