package main

import "fmt"

// Функцию необходимо определить + Функциию необходимо вызвать
//Сигнатура функций и их определение
// func functionName(params type) typeReturnValue {
//     //body
//    }

//Простейшая функция (определить функцию можно как до момента ее вызова в функции main,
// так и в любом месте пакета, главное чтобы она была определена в принципе где-нибудь)
func add(a int, b int) int {
	result := a + b
	return result
}

//Функция с однотипными параметрами
func mult(a, b, c, d int) int {
	result := a * b * c * d
	return result
}

//Возврат больше чем одного значения (returnType1, returnType2.....)
func rectangleParameters(length, width float64) (float64, float64) {
	var perimeter = 2 * (length + width)
	var area = length * width

	return perimeter, area
}

//Именованный возврат значений
func namedReturn(a, b int) (perimeter int32, area int) {
	perimeter = int32(2 * (a + b))
	area = a * b
	return // Не нужно указывать возвращаемые переменные
}

//При вызове оператора return функцию прекращает свое выполнение и возвращает что-то
func funcWithReturn(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}

	if a == b {
		return a, true
	}

	return 0, false
}

//Что вернется в случае, если return в принципе не указан (или он пустой)
func emptyReutrn(a int) {
	fmt.Println("I'M emptyReturn with parameter:", a)
}

//Variadic parameters (континуальные параметры)
func helloVariadic(a ...int) {
	fmt.Printf("value %v and type %T\n", a, a)
}

//Смешение параметров с variadic (
// 	1. Континуальынй параметр всегда самый последний
//  2. Variadic параметр - на всю функцию один (для удобочитаемости)
// )
func someStrings(a, b int, words ...string) {
	fmt.Println("Parameter:", a)
	fmt.Println("Parameter:", b)
	var result string
	for _, word := range words {
		result += word
	}
	fmt.Println("Result concat:", result)
}

//Передача слайса или использование variadic parameters?
func sumVariadic(nums ...int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

func sumSlice(nums []int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println("Hello world")
	// Вызов простейшей функции
	res := add(10, 20)
	fmt.Println("Result of add(10, 20):", res)
	fmt.Println("Result of mult(1, 2, 3, 4):", mult(1, 2, 3, 4))
	fmt.Println("==================================")
	per, area := rectangleParameters(10.5, 10.5)
	newPer, _ := rectangleParameters(10, 10)
	fmt.Println("Area of rect(10.5, 10.5):", area)
	fmt.Println("Perimeter of rect(10.5, 10.5):", per)
	fmt.Println("NewPer:", newPer)
	fmt.Println("==================================")
	secondPer, secondArea := namedReturn(10, 20)
	fmt.Println("namedReturn", secondArea, secondPer)

	emptyReutrn(10) //нельзя присвоить
	fmt.Println("==================================")
	helloVariadic(10, 20, 30, 40, 50, 60, 60)
	helloVariadic()
	someStrings(2, 3, "A", "B", "C")
	fmt.Println("==================================")
	sum1 := sumVariadic(10, 30, 40, 50, 60)
	sliceNumber := []int{10, 20, 30}
	sum2 := sumVariadic(sliceNumber...)
	fmt.Println(sum1, sum2)

	//Анонимная функция (синтаксис)
	anon := func(a, b int) int {
		return a + b
	}

	fmt.Println("Anon:", anon(20, 30))
	fmt.Println("BigFunction(10, 20):", bigFunction(10, 20))

}

//Анонимная функция внутри явной
func bigFunction(aArg, bArg int) int {
	return func(a, b int) int { return a + b + 1 }(aArg, bArg)
}
