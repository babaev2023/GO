package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

//Метод, в котором получаетль копируется и в его теле происходит работа с локальной копией
func (e Employee) SetName(newName string) {
	e.name = newName
}

//Метод, в котором получатель передается по ссылке (в теле метода работаем с ссылкой на экземпляр)
func (e *Employee) SetSalary(newSalary int) {
	e.salary = newSalary
}

//Используйте методы с PointerReciever'ом в ситуациях когда:
// 1) Изменения в работе метода над экземпляром должны быть видны на вызывающей стороне
// 2) Когда экземпляр достаточно "большой", то есть копировать его дорого
// 3) С ними может работать любой вид экземпляра

func main() {
	e := Employee{"Alex", 300}
	fmt.Println("Before setting parameters:", e)
	e.SetName("Bob")
	fmt.Println("After SetName call:", e)
	e.SetSalary(4500) //4. Вызов метода у ссылки на сотрудника
	// 5. Аналогично явному вызову (&e).SetSalary(9999)
	fmt.Println("After SetSalary call:", e)
}
