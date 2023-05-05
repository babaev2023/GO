package main

import "fmt"

//В Go принято называть интерфейсы , с окончанием на `er` (Describer, Worker, Pooller,....)
type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Printf("Student name: %s and age %d\n", std.name, std.age)
}

type Professor struct {
	name string
	age  int
}

func (prof *Professor) Describe() {
	fmt.Printf("Professor name %s and age %d\n", prof.name, prof.age)
}

func main() {
	var descr1 Describer
	stud1 := Student{"ALEX", 30}
	descr1 = stud1 //Student реализует интерфейс Describer
	descr1.Describe()

	stud2 := &Student{"JOHN", 20}
	descr1 = stud2
	descr1.Describe()

	var descr2 Describer
	prof1 := &Professor{"TRUE JOHN", 33}
	descr2 = prof1 // &Professor реализует интерфейс Describer
	descr2.Describe()

	prof2 := Professor{"Bob Marle", 65}
	prof2.Describe() // Здесь ссылку на &prof берет компилятор
	//descr2 = prof2   //Professor не реализует интерфейс Describer
	// Дело в том, что сам по себе интерфейс - не референсный тип

}
