package main

import (
	"fmt"
	"strconv"
)

func main() {

	var decimal int = 43

	binary := ""

	fmt.Println("Введите десятичное число")
	fmt.Scan(&decimal)

	if decimal%2 == 0 {
		fmt.Println("четное")

	} else {
		fmt.Println("нечетное")

	}
outer:
	for {

		if decimal%2 == 0 {

			decimal = decimal / 2

			binary = strconv.Itoa(0) + binary

		} else {

			decimal = (decimal - 1) / 2

			binary = strconv.Itoa(1) + binary
		}

		if decimal == 0 {
			break outer
		}

	}

	fmt.Println("Перевод в двоичное:", binary)

}
