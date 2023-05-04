package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		found := false

		// проверяем, что кратно 3, и выводим 333
		if i%3 == 0 {
			fmt.Printf("333")
			found = true
		}
		// проверяем, что кратно 5, и выводим 555
		if i%5 == 0 {
			fmt.Printf("555")
			found = true
		}
		// если число кратно 3 и 5, выводим 333555

		if !found {
			// если не выполнилось ни одно из условий, выводим число
			fmt.Println(i)
			continue
		}

		fmt.Println()
	}
}
