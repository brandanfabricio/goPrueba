package users

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id   int
	Name string
}

type Employee struct {
	User
	Active bool
}
type Cashier interface{
	CalcTotal(item... float32) float32
	desastivar()
	reactive()
}

type Admin interface{
	DesactivarEmployeer(c Cashier)
	ActivarEmployeer(c Cashier)
}

func NewEmployee(name string) *Employee {
	return &Employee{
		User:   User{Id: rand.Intn(1000), Name: name},
		Active: true,
	}
}

func (e *Employee) DesactivarEmployeer(c Cashier) {
	c.desastivar()
}

func (e *Employee) ActivarEmployeer(c Cashier) {
	c.reactive()
}



func (e *Employee) CalcTotal(item ...float32)float32 {
	
	if !e.Active{
		fmt.Println("El Cajero esta desastivado")
		return 0
	} 
	var sum float32 

	for _, i := range item {
		sum +=i
	}
	return sum * 1.15
}

func (e *Employee)desastivar(){
	e.Active = false
}
func (e *Employee)reactive(){
	e.Active = true
}





