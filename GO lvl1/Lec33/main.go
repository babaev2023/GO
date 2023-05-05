package main

import "fmt"

type University struct {
	city string
	name string
}

//Данный метод явно связан только с структурой University
func (u *University) FullInfoUniversity() {
	fmt.Printf("University Name: %s and City: %s\n", u.name, u.city)
}

//В структуру Professor встроены поля структуры University (переходят и все методы тоже)
type Professor struct {
	fullName string
	age      int
	University
}

func main() {
	p := Professor{
		fullName: "TRUE JOHN",
		age:      33,
		University: University{
			city: "Moscow",
			name: "NEW"}}
	//Вызываем метод структуры University через экземпляр профессора
	p.FullInfoUniversity()
}
