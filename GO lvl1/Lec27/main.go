package main

import "fmt"

//Указатели на массивы. Почему так делать не надо
func mutation(arr *[3]int) {
	arr[0] = 9
	arr[2] = 10
}

//Используйте лучше слайсы
func mutationSlice(sls []int) {
	sls[1] = 9
	sls[2] = 10
}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println("Arr before mutation:", arr)
	mutation(&arr)
	fmt.Println("Arr after mutation:", arr)

	newArr := [3]int{1, 2, 4}
	fmt.Println("newArr before mutationSlice:", newArr)
	mutationSlice(newArr[:])
	fmt.Println("newArr after mutationSlcie:", newArr)
}
