package main

import "fmt"

//Указатели - переменная, хранящая в качестве значения - адрес в памяти другой переменной

func main() {

	//Определение указателя на что-то
	var variable int = 30
	var pointer *int = &variable //&.... - операция взятия адреса в памяти
	//Выше у нас создан указатель на переменную variable
	//В pointer лежит строка - это место в памяти, где хранится int значение 30
	fmt.Printf("Type of pointer %T\n", pointer)
	fmt.Printf("Value of pointer %v\n", pointer)

	//А какое нулевое значение для указатели?
	var zeroPointer *int //zeroValue имеет значение nil
	fmt.Printf("Type %T and value %v\n", zeroPointer, zeroPointer)
	if zeroPointer == nil {
		zeroPointer = &variable
		fmt.Printf("After initializatoin type %T and value %v\n", zeroPointer, zeroPointer)
	}
	fmt.Println("=============================")
	//Разыменование указателя (получение значения): *pointer - возвращает значение, хранимое по адресу
	var numericValue int = 300000000
	pointerToNumeric := &numericValue

	fmt.Printf("Value in numericValue is %v Address is %v\n", *pointerToNumeric, pointerToNumeric)

	//Создание указателей на нулевые занчения типов
	zeroPoint := new(int) // Создает под капотом zeroValue для int, и возвращает адрес, где этот 0 хранится
	fmt.Printf("Value in *zeroPointer %v\nAddress is %v\n", *zeroPoint, zeroPoint)

	//Изменение значения хранимого по адресу через указатель
	zeroPointerToInt := new(int)
	fmt.Printf("Addres is %v and Value in zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt)
	*zeroPointerToInt += 40
	fmt.Printf("Addres is %v and New Value in zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt)
	fmt.Println("=============================")
	b := 333
	a := &b
	c := &b
	*a++
	*c += 100
	fmt.Println(b)

	//Передача указателей в функции
	// Колоссальный прирост производительности засчет того, что передается не значение (не копируется)
	// а передается лишь адрес в памяти, за которым уже хранится какое-то значение
	sample := 1

	fmt.Println("Origin value of sample:", sample)
	changeParam(&sample)
	fmt.Println("After changing sample is:", sample)

	//Возврат поинтера из функции
	ptr1 := returnPointer()
	ptr2 := returnPointer()
	fmt.Printf("Ptr1: %T and address %v and value %v\n", ptr1, ptr1, *ptr1)
	fmt.Printf("Ptr2: %T and address %v and value %v\n", ptr2, ptr2, *ptr2)

}

//Инициализация функции, возвращающей указатель
func returnPointer() *int {
	var numeric int = 3333
	return &numeric //В момент возврата Go перемещает данную переменную в кучу
}

//Определдение фукнции, принимающей параметр как указатель
func changeParam(val *int) {
	*val += 100
}
