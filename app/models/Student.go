package models

import "fmt"

type studentMan struct {
	Name  string
	Age   int
	score float64
}

func NewStudent(name string, age int) *studentMan {
	return &studentMan{
		name,
		age,
		10.0,
	}
}

func (s *studentMan) GetScore() float64 {
	return s.score
}

type AllStudent struct {
	Name string
	Age  int
	sla  float64
}

func (a *AllStudent) GetInfo() {
	fmt.Println("this is a func")
}

// SmallStudent 继承 AllStudent
type SmallStudent struct {
	AllStudent
	MathScore float64
}

func (s *SmallStudent) SetSla(sla float64) {
	s.sla = sla
}
