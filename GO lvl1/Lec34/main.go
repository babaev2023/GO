package main

import "fmt"

type Rectangle struct {
	length, width int
}

//Реализуем функцию и метод для подсчет периметра прямогуольника

//Метод
func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

//Функция
func Perimeter(r Rectangle) int {
	return 2 * (r.length + r.width)
}

//Метод, который меняет значение состояния структуры на новое, но этот метод - Value Reciever
func (r Rectangle) SetLength(newLength int) {
	r.length = newLength
}

func main() {
	//Создаем экземпляр прямоугольника
	rectangleAsValue := Rectangle{10, 10}
	fmt.Println("Call function for rectangleAsValue:", Perimeter(rectangleAsValue))
	fmt.Println("Call method for rectangleAsValue:", rectangleAsValue.Perimeter())

	//Создадим указатель на прямоугольник
	rectangleAsPointer := &rectangleAsValue
	fmt.Println("Call method for rectangleAsPointer:", rectangleAsPointer.Perimeter())

	//Вызываем метод SetLength у экземпляра rectangleAsValue
	fmt.Println("Before call method SetLength:", rectangleAsValue)
	rectangleAsValue.SetLength(500)
	fmt.Println("Aftet call method SetLength:", rectangleAsValue)

	//Вызываем метод SetLength у ссылки на rectangleAsValue
	rectangleAsPointer.SetLength(50000)
	fmt.Println("After call method SetLength for &rectangle:", *rectangleAsPointer)

}
