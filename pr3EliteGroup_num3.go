package main

import "fmt"

func Pascal(prevRow []int, col int) int {
	if col == 0 || col == len(prevRow) {
		return 1
	}
	return prevRow[col-1] + prevRow[col]
}

func main() {
	var n int
	fmt.Println("Введите количество строк: ")
	fmt.Scanln(&n)

	row := []int{1}
	fmt.Println(row[0])

	for i := 1; i < n; i++ {
		newRow := make([]int, i+1)
		newRow[0] = 1
		newRow[i] = 1
		for j := 1; j < i; j++ {
			newRow[j] = Pascal(row, j)
		}
		row = newRow
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}
