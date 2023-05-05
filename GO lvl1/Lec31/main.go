package main

import (
	"fmt"
	"math"
)

// Функция или метод

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width int
}

// Создадим функцию и метод Perimeter  для структуры Circle
func PerimeterCircle(c Circle) float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) Perimeter() float64 {
	return c.radius * 2 * math.Pi
}

// Создадим функцию и метод Perimeter для структуры Rectangle
func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

// В Go разрешено создавать методы с одинаковыми именами в пределах одной структуры. Главное, чтобы
// получать метода в разных структурах (где реализован мтеод со сходим именем) отличался.
func PerimeterRectangle(r Rectangle) int {
	return 2 * (r.length + r.width)
}

func main() {
	c := Circle{10.5}
	fmt.Println("Call function:", PerimeterCircle(c))
	fmt.Println("Call method:", c.Perimeter())
	fmt.Println("=======================================>")
	r := Rectangle{10, 20}
	fmt.Println("Call function for rectangle:", PerimeterRectangle(r))
	fmt.Println("Call method for rectangle:", r.Perimeter())
}
