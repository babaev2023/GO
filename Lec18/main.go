package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30}
	slice[0] = slice[0] * 10
	slice[1] = 200
	slice = append(slice, 200) //Добавление нового элемента
	for i, v := range slice {
		fmt.Println(i, v)

	}

}
