package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

//Методы - функции, привязанные к определенным структурам
//func (s Struct) MethodName(parameters type) type {}
//      Reciever - получатель метода
func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Postion:", e.position)
	fmt.Printf("Salary : %d %s\n", e.salary, e.currency)
}

func main() {
	emp := Employee{
		name:     "ALEX",
		position: "Senior Golang developer",
		salary:   500,
		currency: "RUB",
	}
	//Вызов метода
	emp.DisplayInfo()
}
