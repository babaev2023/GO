package main

import (
	_ "Lec46/circle"
	"Lec46/rectangle"
	"fmt"
)

// Все импорты (вне зависимости , стандартные или пользовательские) сортируются по алфавиту
func init() {
	fmt.Println("Init function for main package!")
	fmt.Println("Name:", name, "Age:", age)
}

var (
	name string = "ALEX"
	age  int    = 30
)

func main() {
	r := rectangle.New(10, 10)
	fmt.Println(r)

}
