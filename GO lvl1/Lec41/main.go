package main

import "fmt"

func DoSomething(pretendent interface{}) {
	switch pretendent.(type) { // Пытаемся извлечь нижележащий тип
	case string:
		fmt.Println("This is string!")
	case int:
		fmt.Println("This is int!")
	default:
		fmt.Println("Unknown type! But i'm working!")
	}
}

func main() {
	DoSomething(10)
	DoSomething("Hello world!")
	DoSomething(true)
}
