package main

import "fmt"

//Площадь прямоугольника

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func Area(r *Rectangle) int {
	return r.length * r.width
}

func (r *Rectangle) SetWidth(newWidth int) {
	r.width = newWidth
}

func main() {
	//Значение исходное
	rectangleAsValue := Rectangle{10, 10}
	//Ссылка на исходное значение
	rectangleAsPointer := &rectangleAsValue
	fmt.Println("Call Area function for &rectangle:", Area(rectangleAsPointer))
	fmt.Println("Call Area method for &rectangle:", rectangleAsPointer.Area())

	//Вызываем метод у value - исходного значения
	fmt.Println("Call Area method for rectangle:", rectangleAsValue.Area())
	//Вызываем функцию с параметром value - исходное значение
	//Area(rectangleAsValue)

	//Распечатаем исходный прямоугольник
	fmt.Println("Before changing width:", rectangleAsValue)

	//Вызываю метод SetWidth у &rectangle
	rectangleAsPointer.SetWidth(100)
	fmt.Println("After change via method SetWidth for &rectangle:", rectangleAsValue)
	fmt.Println("Call Area method for &rectangle after change via method SetWidth:", rectangleAsPointer.Area())
	//Вызов метода SetWidth у rectangle
	rectangleAsValue.SetWidth(1000)
	fmt.Println("After change via method SetWidth for rectangle:", rectangleAsValue)
	fmt.Println("Call Area method for &rectangle after change via method SetWidth:", rectangleAsPointer.Area())
}
