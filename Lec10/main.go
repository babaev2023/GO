package main

import "fmt"

func main() {

	pricelist := map[string]int{"хлеб": 20, "молоко": 200, "масло": 100, "колбаса": 500, "соль": 10, "огурцы": 200, "сыр Гауди": 600, "ветчина": 700, "буженина": 900, "помидоры": 2500, "рыба": 300, "хамон": 1500}
	order := []string{"молоко", "соль", "хамон", "огурцы"}
	total := 0
	fmt.Println("Перечень товаров дороже 500:")
	for k, v := range pricelist {
		if v > 500 {
			fmt.Println(k)
		}
	}
	for _, v := range order {
		total += pricelist[v]
	}
	fmt.Println("Стоимость заказа ", total)
}
