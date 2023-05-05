package main

import (
	"fmt"
	"unicode/utf8"
)

// Описание интерфейса
type BigWord interface {
	BigString() bool
}

// Наш претендент, которого надо научить делать BigString() bool
type MyString string

// Реализация BigString() у претендента MyString
func (mss MyString) BigString() bool {
	if utf8.RuneCountInString(string(mss)) > 10 {
		return true
	}
	return false
}

func main() {
	sample := MyString("ABC*766hjdhsjhdjsjdhjs")
	var interfaceSample BigWord // Объявление переменной, типа интерфейса BigWord
	interfaceSample = sample    // присваивание значения для переменной тип BigWord возможно,
	fmt.Println("BigString?", interfaceSample.BigString())

}
