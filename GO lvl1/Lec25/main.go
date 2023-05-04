package main

import (
	"fmt"
	"math"
)

//Константы - это неизменяемые переменные

const (
	MAIN_PORT = "8001"
)

func main() {
	//Объвление одной константы
	const a = 10
	fmt.Println(a)
	//Объявление блока констант с областью видимости внутри функции main
	const (
		ipAddress = "192.168.0.1"
		port      = "80"
		dbName    = "postgres"
	)
	fmt.Println("ipAddress value:", ipAddress)
	fmt.Println("port:", port)
	fmt.Println("dbName:", dbName)
	fmt.Println(checkPortIsValid())

	// Константу никак нельзя поменять в ходе работы программы
	// const b = 200
	// b = 30

	// Значения констант ДОЛЖНЫ БЫТЬ ИЗВЕСТНЫ на момент компиляции
	var sqrt = math.Sqrt(25)
	//const sqrt = math.Sqrt(25) //Нельзя присвоить в константу что-либо, что является результатом вызова функции, метода
	fmt.Println("Var sqrt:", sqrt)

	//Типизированные и нетипизированные константы
	const ADMIN_EMAIL string = "admin@admin.com" // Указание типа - повышение читабельности кода

	//Нетипизирвоанные константы и их профит
	//При использовании нетипизированных констант мы разрешаем компилятору
	//исопльзовать неявное приведение типов в момент присваивания значеий констант в перменные
	const NUMERIC = 10
	var numInt8 int8 = NUMERIC
	var numInt32 int32 = NUMERIC
	var numInt64 int64 = NUMERIC
	var numFloat32 float32 = NUMERIC
	var numComplex complex64 = NUMERIC

	fmt.Printf("numInt8 value %v type %T\n", numInt8, numInt8)
	fmt.Printf("%v + %v is %v\n", numInt8, NUMERIC, numInt8+NUMERIC)
	fmt.Printf("numInt32 value %v type %T\n", numInt32, numInt32)
	fmt.Printf("numInt64 value %v type %T\n", numInt64, numInt64)
	fmt.Printf("numFloat32 value %v type %T\n", numFloat32, numFloat32)
	fmt.Printf("numComplex value %v type %T\n", numComplex, numComplex)
	//Константы в Go зашиваются в момент компиляции в RUNTIME программы и не выбрасываются до ее окончания

}

// Проверка порта
func checkPortIsValid() bool {
	if MAIN_PORT == "8001" {
		return true
	}
	return false
}
