package main

import "fmt"

//DEFER - оператор отложенного ВЫЗОВА функции/метода.

//Создадим отложенную фукнцию.

//С входными параметрами

func CheckDBCloseConncetion(a int) {
	fmt.Println("Check started......Value numIP in deferred:", a)
	fmt.Println("Check done! Connection refused!")
}

//Если функция принимает входные параметры и данная функция используется в связки с defer
// то :
// параметры расчитываются в момент передачи их в функцию
// А вызов функции с уже давно рассчитанным параметром осуществляется в момент
// завершения вышележащей функции

// В какой момент defer вызывается?
func OuterFunc() int {
	defer fmt.Println("I'm deferred print function!")
	fmt.Println("OuterFunc started!")
	var result = 10
	fmt.Println("OuterFunc finished. Ready to return value!")
	return result
}

func main() {
	defer fmt.Println("Step 1 defer")
	defer fmt.Println("Step 2 defer")
	defer fmt.Println("Step 3 defer")
	defer fmt.Println("Step 4 defer")

	var numIP = 10
	p := &numIP
	//defer CheckDBCloseConncetion(numIP) // defer означает, что данная функция будет вызвана при завершении main() с параметром,
	//значение которого расчитывается на момент 25-ой строки.
	*p++
	fmt.Println("Value numIP in main():", numIP)
	fmt.Println("Main function started")
	fmt.Println("Main function ended")
	fmt.Println("Value form OuterFunc on main side is:", OuterFunc())
}
