package main

import "fmt"

func main() {
	var year int
	fmt.Println("Введите год: ")
	fmt.Scanln(&year)

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("Год", year, "является високосным")
	} else {
		fmt.Println("Год", year, "не является високосным")
	}
}
