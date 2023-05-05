// Если хочется скрестить 2 интерфейса и создать единый уровень абстракции
package main

import "fmt"

type PerimeterCalculator interface {
	Perimeter() int
}

type AreaCalculator interface {
	Area() int
}

//Соберем новый интерфейс из двух старых
type FigureFeatureCalculator interface {
	PerimeterCalculator // встраиваем интерфейсы
	AreaCalculator
	// Наслдеование интерфейсов
	//Perimeter() int
	//Area() int
}

type Rectangle struct {
	a, b  int
	color string
}

//Реализуем интерфейс PerimeterCalculator
func (r Rectangle) Perimeter() int {
	return 2 * (r.a + r.b)
}

//Реализуем второй интерфейс AreaCalculator
func (r Rectangle) Area() int {
	return r.a * r.b
}

func main() {
	var pCalc PerimeterCalculator
	fmt.Printf("Value %v and Type %T\n", pCalc, pCalc)
	var aCalc AreaCalculator
	rect := Rectangle{10, 20, "black"}
	pCalc = rect // Стурктура Rectangle удовлетворяет интерфейсу PerimeterCalculator
	fmt.Printf("Value %v and Type %T\n", pCalc, pCalc)
	fmt.Println("Perimeter:", pCalc.Perimeter())

	aCalc = rect // Структура Rectangle удовлетворяет интерфейсу AreaCalculator
	fmt.Println("Area:", aCalc.Area())

	var ffCalc FigureFeatureCalculator
	ffCalc = rect // Структура Rectangle удовлетворяет FigureFeatureCalculator
	fmt.Println("Perimeter:", ffCalc.Perimeter())
	fmt.Println("Area:", ffCalc.Area())

}
