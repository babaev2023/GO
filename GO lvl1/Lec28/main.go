package main

import "fmt"

//Структура

//Определение структуры (явное определение)
type Student struct {
	firstName string
	lastName  string
	age       int
}

//Если имеется ряд состояний одного типа, можно сделать так
type AnotherStudent struct {
	firstName, lastName, groupName string
	age, courseNumber              int
}

//Структура с анонимными полями
type Human struct {
	firstName string
	lastName  string
	string
	int
	bool
}

func PrintStudent(std Student) {
	fmt.Println("==================")
	fmt.Println("FirstName:", std.firstName)
	fmt.Println("LastName:", std.lastName)
	fmt.Println("Age:", std.age)
}

func main() {

	//Создание представителей структуры
	stud1 := Student{
		firstName: "IVAN",
		age:       33,
		lastName:  "IVANOV"}
	PrintStudent(stud1)
	stud2 := Student{"DARIA", "IVANOVA", 23} // Порядок указания свойств - такой же как в структуре
	PrintStudent(stud2)

	//Что если не все поля структуры определить?
	stud3 := Student{
		firstName: "Only NAME"}
	PrintStudent(stud3)

	//Анонимные структуры (структура без имени)
	anonStudent := struct {
		age           int
		groupID       int
		proffesorName string
	}{
		age:           20,
		groupID:       7,
		proffesorName: "ALEX"}
	fmt.Println("AnonStudent:", anonStudent)
	fmt.Println("================++==============")

	//Доступ к состояниям и их модфикация
	studVova := Student{"Vova", "IVANOV", 21}
	fmt.Println("firstName:", studVova.firstName)
	fmt.Println("lastName:", studVova.lastName)
	fmt.Println("age 2020:", studVova.age)
	studVova.age += 2
	fmt.Println("new age 2022:", studVova.age)

	//Инициализация пустой структуры
	emptyStudent1 := Student{}
	var emptyStudent2 Student
	PrintStudent(emptyStudent1)
	PrintStudent(emptyStudent2)

	//Указатели на экземпляры структур
	studPointer := &Student{
		firstName: "IRINA",
		lastName:  "POINTER",
		age:       22}
	fmt.Println("Value studPointer:", studPointer)
	secondPointer := studPointer
	(*secondPointer).age += 3
	fmt.Println("Value afterPointerModify:", studPointer)
	studPointerNew := new(Student)
	fmt.Println(studPointerNew)

	//Работа с доступ к полям структур через указатель
	fmt.Println("Age via (*...).age:", (*studPointer).age)
	fmt.Println("Age via .age:", studPointer.age)
	//Неявно происходит разыменование указателя studpointer и запрос соотв поля

	//Создание экземпляра с анонимными полями структуры
	human := &Human{
		firstName: "Bob",
		lastName:  "John",
		string:    "Info 123456",
		int:       -1,
		bool:      true,
	}

	fmt.Println(human)
	fmt.Println("Anon field string:", human.string)
}
