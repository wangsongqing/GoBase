package models

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
}

type SmallStudent struct {
	AllStudent
	MathScore float64
}
