package main

import (
	"fmt"
)

func main() {
	date := 2023
	y := 2000
	fmt.Println("Введите год рождения")
	fmt.Scan(&date)

	switch y = date; {
	case y >= 1946 && y < 1965:
		fmt.Println("Привет, бумер!")
	case y >= 1965 && y < 1981:
		fmt.Println("Привет, представитель X!")
	case y >= 1981 && y < 1997:
		fmt.Println("Привет, миллениал!")
	case y >= 1997 && y < 2013:
		fmt.Println("Привет, зумер!")
	case y >= 2013:
		fmt.Println("Привет, альфа!")
	default:
		fmt.Println("Ошибка")
	}
}
