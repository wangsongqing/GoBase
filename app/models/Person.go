package models

type person struct {
	Name string
	age  int
	sla  float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// SetAge 设置是有属性
func (p *person) SetAge(age int) {
	p.age = age
}

func (p *person) SetSla(sla float64) {
	p.sla = sla
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) GetSla() float64 {
	return p.sla
}
