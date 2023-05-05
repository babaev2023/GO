package main

import "fmt"

// Интерфейсы - явно декларированный набор сигнатур ПОВЕДЕНИЙ (набора методов)

// Объявление интерфейса
type FigureBuilder interface {
	//Набор сигнатур методов, которые необходимо реализовать в структуре-претенденте
	//У претендента должен быть метод Area() возвращающий int
	Area() int
	//У претендента должен быть метод Perimeter() возвращающий int
	Perimeter() int
}

//Декларируем претендентов
//Первый претендент - это прямоугольник
// У него есть 2 метода -
//Area() int
//Perimter() int
//Когда эти методы реализованы - RECTANGLE УДОВЛЕТВОРЯЕТ УСЛОВИЯМ ИНТЕРФЕЙСА FigureBuilder
//RECTANGLE РЕАЛИЗУЕТ ИНТЕРФЕЙС FigureBuilder
type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

// Второй претендент - это окружность
// У нее есть 2 метода -
//Area() int
//Perimter() int
//Когда эти методы реализованы - CIRCLE УДОВЛЕТВОРЯЕТ УСЛОВИЯМ ИНТЕРФЕЙСА FigureBuilder
//CIRCLE РЕАЛИЗУЕТ ИНТЕРФЕЙС FigureBuilder
type Circle struct {
	radius int
}

func (c Circle) Area() int {
	return 3 * c.radius * c.radius
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}

func main() {

	//Создадим по 3 экземпляра
	r1 := Rectangle{10, 10}
	r2 := Rectangle{20, 20}
	r3 := Rectangle{30, 30}
	c1 := Circle{10}
	c2 := Circle{20}
	c3 := Circle{30}

	//Посчитаем общую площадь этих  фигур
	rectangles := []Rectangle{r1, r2, r3}
	totalSumAreaOfRectangles := 0
	for _, rect := range rectangles {
		totalSumAreaOfRectangles += rect.Area()
	}

	circles := []Circle{c1, c2, c3}
	totalSumAreaOfCircles := 0
	for _, circ := range circles {
		totalSumAreaOfCircles += circ.Area()
	}

	fmt.Println("Total area: ", totalSumAreaOfRectangles+totalSumAreaOfCircles)
	fmt.Println("==>")
	//Воспользуемся интерфейсом, указанным выше
	figures := []FigureBuilder{r1, r2, r3, c1, c2, c3} //слайс экземпляров, удовлетворяющих интерфейсу FigureBuilder

	//Посчитаем общую площадь
	total := 0
	for _, fig := range figures {
		total += fig.Area()
	}
	fmt.Println("Total area: ", total)

}
