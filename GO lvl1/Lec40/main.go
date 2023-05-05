package main

import "fmt"

//Объявляем интерфейс, декларирующий поведенческий -паттерн (набор методов под реализацию)
type Worker interface {
	Work()
}

type Money interface {
	Mymoney()
}

//Объявляем структуру. Данная структура - претендент на удовлетворение интерфейса Worker
type Employee struct {
	name string
	age  int
}

//Реализуем метод Work(), чтобы структура Employee удовлетворяла интерфейсу Worker
func (e Employee) Work() {
	fmt.Println("Now Employee with name", e.name, "is working!")
}

//Функция, которая будет принимать аргументы типа Worker и что-то с ними делать
func Describer(w Worker) {
	fmt.Printf("Interface with type %T and value %v\n", w, w)
}
func Describer2(w Money) {
	fmt.Printf("Interface with type %T and value %v\n", w, w)
}

//Объявляем структура. Данная структура - второй пренедент на удовлетворение интерфейса Worker
type Student struct {
	name         string
	courseNumber int
}

func (s Student) Mymoney() {
	fmt.Println("Student with name", s.name, "want to money")
}

func (s Student) Work() {
	fmt.Println("Student with name", s.name, "is working!")
}

func main() {
	//Создадим экземпляр Employee
	emp := Employee{"ALEX", 33}
	var workerEmployee Worker = emp // Присваиваем сотрудника в переменную типа Worker
	workerEmployee.Work()
	Describer(workerEmployee) // В резульатте видим, что тип интерфейса определяется структурой, его реализующей,
	//а значение интерфейса - это соответственно значение состояний структуры

	//Создадим экземпляр Student
	std := Student{"Ivanov", 18}
	var workerStudent Worker = std
	var moneyStudent Money = std

	workerStudent.Work()
	Describer(workerStudent) // Принял внутренний тип Student, а значение - равно значению полей экземпляра
	moneyStudent.Mymoney()
	Describer2(moneyStudent) // Принял внутренний тип Student, а значение - равно значению полей экземпляра

	//Созаддим набор тех, кто умеет  Work()
	var workers []Worker = []Worker{workerStudent, workerEmployee}
	for _, worker := range workers {
		Describer(worker)
		// Данная функция вызывается у разных экземпляров благодря тому, кто для ее вызова
		//экземпляру нужно удовлетворить некому контракту (поведенческому паттерну). Если структура экземпляра поддерживает
		//данный паттерн - то у экземпляра 100% можно вызвать все необходимые для этого методы.
	}
}
