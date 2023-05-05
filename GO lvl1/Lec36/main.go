package main

import "fmt"

type MyInt int

func (msdi *MyInt) IsEven() string {
	if *msdi%2 == 0 {
		return "Четное"
	}
	return "Нечетное"
}

func main() {
	num1 := MyInt(202)
	num2 := MyInt(201)
	fmt.Println("202:", num1.IsEven())
	fmt.Println("201:", num2.IsEven())

}
