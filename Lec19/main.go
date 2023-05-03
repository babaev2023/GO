package main

import (
	"fmt"
)

func main() {
	// Слайсы
	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[0:2] //Слайс иницилиазируется пустыми квадратными скобками
	fmt.Println("Slece[0:2]:", startSlice)
	// Слайс - на существующем массиве

	secondSlice := []int{10, 20, 30, 40}
	fmt.Println("SecondSlice:", secondSlice)

	//Изменение элеметнов среза
	originArr := [...]int{30, 40, 50, 60, 70, 80}
	firstSlice := originArr[1:4]
	for i, _ := range firstSlice {
		firstSlice[i]++
	}
	fmt.Println("OriginArr:", originArr)
	fmt.Println("FirstSlice:", firstSlice)

	// Один массив и два среза
	fSlice := originArr[:]
	sSlice := originArr[2:5]

	fmt.Println(originArr, fSlice, sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println(originArr, fSlice, sSlice)

	// Длина и емкость слайса
	wordsSlice := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSlice, "Lenght:", len(wordsSlice), "Capacity", cap(wordsSlice))
	wordsSlice = append(wordsSlice, "four")
	fmt.Println("slice:", wordsSlice, "Lenght:", len(wordsSlice), "Capacity", cap(wordsSlice))

	// Make
	sl1 := make([]int, 10, 15) // 10 длина 15 емкость
	fmt.Println(sl1)

	// Добавление элементов в СРЕЗ
	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords)
	anotherSlice := []string{"one1", "two1", "three1"}
	myWords = append(myWords, anotherSlice...)
	fmt.Println("myWords:", myWords)

	// Многомерный срез
	mSlice := [][]int{
		{1, 2},
		{3, 4, 5, 6},
		{10, 20, 30, 40},
		{},
	}
	fmt.Println(mSlice)
}
