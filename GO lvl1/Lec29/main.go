package main

import "fmt"

//Вложенные структуры
type University struct {
	age       int
	yearBased int
	infoShort string
	infoLong  string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}

//Встроенные структуры (когда мы добавляем поля одной структуры к другой)
type Professor struct {
	firstName  string
	lastName   string
	age        int
	greatWork  string
	University // В этом месте происходит добавление всех полей структуры Uni в Professor
}

func main() {
	//Создание экземпляров структур с вложением
	stud := Student{
		firstName: "IVAN",
		lastName:  "IVANOV",
		university: University{
			yearBased: 2022,
			infoShort: "Info",
			infoLong:  "Long Info"}}

	//Получение доступа к вложенным полям структур
	fmt.Println("FirstName:", stud.firstName)
	fmt.Println("LastName:", stud.lastName)
	fmt.Println("Year based Uni:", stud.university.yearBased)
	fmt.Println("Long info:", stud.university.infoLong)
	fmt.Println("========================")
	//5. Создание экземпляра с встраиванием структур
	prof := Professor{
		firstName: "ALEX",
		lastName:  "TRUE",
		age:       21,
		greatWork: "C++",
		University: University{
			yearBased: 2022,
			infoShort: "Info",
			age:       2021 - 1900,
		},
	}
	//Обращение к состояниям с встроенной структурой
	fmt.Println("FirstName:", prof.firstName)
	fmt.Println("Year based:", prof.yearBased)
	fmt.Println("Info Short:", prof.infoShort)
	fmt.Println("Age++", prof.age)
	fmt.Println("Age:", prof.University.age) //prof.age - получим доступ к полю ВЫШЕЛЕЖАЩЕЙ СТРУКТУРЫ

	//Сравнение экземпляров ==
	//При сравнении экзмеляров происходит сравнение всех их полей друг с другом
	profLeft := Professor{}
	profRight := Professor{}

	fmt.Println(profLeft == profRight)

}
