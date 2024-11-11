// №1 ДОДЕЛАТЬ
// package main
// import "fmt"
// var x int
// var s1 int
// var s2 int
// var count int = 0
// var count1 int = 0
// var decimal int = 1
// var decimal1 int = 1
// func main() {
//     fmt.Scan(&x)
//     fmt.Scan(&s1)
//     fmt.Scan(&s2)
//     k := x
//     for k != 0 {
//         count = count + k%s1*decimal
//         k = k/s1
//         decimal = decimal*10
//     }
//     i := count
//     for i != 0 {
//         count1 = count1 + i%s2*decimal1
//         i = i/s2
//         decimal1 = decimal1*10
//     }
//     fmt.Println(count1)

//  // Преобразование числа в десятичную систему
//  decimal, ok := new(big.Int).SetString(numberStr, sourceBase)
//  if !ok {
//   fmt.Println("Invalid number in source base.")
//   return
//  }

//  // Преобразование из десятичной в целевую систему
//  result := decimal.Text(targetBase)

//  fmt.Println(result)
// }

// №2
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var a, b, c, x1, x2 float64
// 	fmt.Scanln(&a, &b, &c)
// 	d := math.Pow(b, 2) - (4 * a * c)
// 	if d < 0 {
// 		realPart := -b / (2 * a)
// 		imaginaryPart := math.Sqrt(-d) / (2 * a)
// 		x1 := complex(realPart, imaginaryPart)
// 		x2 := complex(realPart, -imaginaryPart)
// 		fmt.Println(x1, x2)
// 	}
// 	if d >= 0 {
// 		x1 = (-b - math.Sqrt(d)) / (2 * a)
// 		x2 = (-b + math.Sqrt(d)) / (2 * a)
// 		fmt.Println(x1, x2)
// 	}
// }

// №3

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func Sortirovka(arr []int) []int {
// 	sort.Slice(arr, func(i, j int) bool {
// 		return abs(arr[i]) < abs(arr[j])
// 	})
// 	return arr
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// func main() {
// var n int
// fmt.Print("Введите количество элементов массива: ")
// fmt.Scan(&n)
// 	arr := make([]int, n)
// fmt.Println("Введите элементы массива:")
// for i := 0; i < n; i++ {
// 	fmt.Scan(&arr[i])
// 	}
// 	sortedArr := Sortirovka(arr)
// 	fmt.Println("Отсортированный массив по модулю:", sortedArr)
// }

//#4

// package main

// import "fmt"

// func mergeSortedArrays(arr1, arr2 []int) []int {
// 	merged := make([]int, 0, len(arr1)+len(arr2))
// 	i, j := 0, 0

// 	for i < len(arr1) && j < len(arr2) {
// 		if arr1[i] <= arr2[j] {
// 			merged = append(merged, arr1[i])
// 			i++
// 		} else {
// 			merged = append(merged, arr2[j])
// 			j++
// 		}
// 	}

// 	// Добавляем оставшиеся элементы из arr1
// 	for ; i < len(arr1); i++ {
// 		merged = append(merged, arr1[i])
// 	}

// 	// Добавляем оставшиеся элементы из arr2
// 	for ; j < len(arr2); j++ {
// 		merged = append(merged, arr2[j])
// 	}

// 	return merged
// }

// func main() {

// 	var n1, n2 int
// 	fmt.Print("Введите количество элементов массива №1: ")
// 	fmt.Scan(&n1)
// 	arr1 := make([]int, n1)
// 	fmt.Println("Введите элементы массива №1:")
// 	for i := 0; i < n1; i++ {
// 		fmt.Scan(&arr1[i])
// 	}

// 	fmt.Print("Введите количество элементов массива №2: ")
// 	fmt.Scan(&n2)
// 	arr2 := make([]int, n2)
// 	fmt.Println("Введите элементы массива №2:")
// 	for i := 0; i < n2; i++ {
// 		fmt.Scan(&arr2[i])
// 	}
// 	mergedArr := mergeSortedArrays(arr1, arr2)
// 	fmt.Println(mergedArr) // Вывод: [1 2 3 5 6 8 9]
// }

// // #5 ДОДЕЛАТЬ НЕ РАБОТАЕТ
// package main
// import "fmt"

// func findSubstring(text, substring string) int {
//  for i := 0; i <= len(text)-len(substring); i++ {
//   found := true
//   for j := 0; j < len(substring); j++ {
//    if text[i+j] != substring[j] {
//     found = false
//     break
//    }
//   }
//   if found {
//    return i
//   }
//  }
//  return -1
// }

// func main() {
//  var text, substring string
//  fmt.Print("Введите строку: ")
//  fmt.Scanln(&text)
//  fmt.Print("Введите подстроку: ")
//  fmt.Scanln(&substring)

//  index := findSubstring(text, substring)

//  if index == -1 {
//   fmt.Println("-1")
//  } else {
//   fmt.Printf("Подстрока найдена на позиции %d\n", index)
//  }
// }

// ----------------------------------------------------------------------------
// УСЛОВИЯ

// №1

// №2 ДОДЕЛАТЬ В КРАСИВОМ ВИДЕ
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func isPalindrom(number int) bool {
// 	strNumber := strconv.Itoa(number)
// 	left, right := 0, len(strNumber)-1

// 	for left < right {
// 		if strNumber[left] != strNumber[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }

// func main() {
// 	var number int
// 	fmt.Print("Введите число: ")
// 	fmt.Scanln(&number)

// 	if isPalindrom(number) {
// 		fmt.Println("Число является палиндромом")
// 	} else {
// 		fmt.Println("Число не является палиндромом")
// 	}
// }

// // №3 ДОДЕЛАТЬ ДО КРАСИВОГО ВИДА

// package main

// import "fmt"

// func max(a, b float64) float64 {
// 	if a > b {
// 	 return a
// 	}
// 	return b
//    }

//    func min(a, b float64) float64 {
// 	if a < b {
// 	 return a
// 	}
// 	return b
//    }

// func main() {
//  var start1, end1, start2, end2, start3, end3 float64
//  fmt.Scanln(&start1, &end1, &start2, &end2, &start3, &end3)

//  max1 := max(start1, start2)
//  max1 = max(max1, start3)
//  min2 := min(end1, end2)
//  min2 = min(min2, end3)

//  if max1 <= min2 {
//   fmt.Println("true")
//  } else {
//   fmt.Println("false")
//  }
// }

// №4 ДОДЕЛАТЬ ДО КРАСИВОГО ВИДА

// package main

// import "fmt"

// func main() {
// 	fmt.Print("Введите предложение: ")
// 	var input string
// 	fmt.Scanln(&input)

// 	}


// 	fmt.Println("Самое длинное слово:", longestWord)
// }

// №5 ДОДЕЛАТЬ ДО КРАСИВОГО ВИДА

// package main

// import "fmt"

// func isLeap(year int) bool {
//  if year%4 != 0 {
//   return false
//  } else if year%100 != 0 {
//   return true
//  } else if year%400 != 0 {
//   return false
//  } else {
//   return true
//  }
// }

// func main() {
//  var year int
//  fmt.Print("Введите год: ")
//  fmt.Scanln(&year)

//  if isLeap(year) {
//   fmt.Println("Год високосный")
//  } else {
//   fmt.Println("Год не високосный")
//  }
// }

// ________________________________________________________________________________
// ЦИКЛЫ

// №1 !!
// package main

// import (
// 	"fmt"
// )

// func fibonacci(n int) {
// 	a, b := 0, 1
// 	for b <= n {
// 		fmt.Print(b, " ")
// 		a, b = b, a+b
// 	}
// 	fmt.Println()
// }

// func main() {
// 	fmt.Println("Введите, до какого числа будет работать алгоритм Фибоначи: ") // Вывод приглашения ПЕРЕД считыванием
// var n int
// fmt.Scanln(&n)
// 	fibonacci(n)
// }

// #2 !!
// package main

// import "fmt"

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func findPrimes(start, end int) {
// 	for i := start; i <= end; i++ {
// 		if isPrime(i) {
// 			fmt.Print(i, " ")
// 		}
// 	}
// 	fmt.Println()
// }

// func main() {
// 	fmt.Println("Введите два числа диапазона нахождения простых чисел:")
// 	var a, b int
// 	fmt.Scanln(&a, &b)
// 	fmt.Println("Простые числа в диапазоне от", a, "до", b, ":")
// 	findPrimes(a, b)
// }

// #3 !!!
// package main

// import "fmt"

// func isArmstrong(n int) bool {
// 	temp := n
// 	sum := 0
// 	digits := 0
// 	for temp > 0 {
// 		digits++
// 		temp /= 10
// 	}
// 	temp = n
// 	for temp > 0 {
// 		digit := temp % 10
// 		sum += power(digit, digits)
// 		temp /= 10
// 	}
// 	return sum == n
// }

// func power(base, exp int) int {
// 	res := 1
// 	for i := 0; i < exp; i++ {
// 		res *= base
// 	}
// 	return res
// }

// func findArmstrong(start, end int) {
// 	for i := start; i <= end; i++ {
// 		if isArmstrong(i) {
// 			fmt.Print(i, " ")
// 		}
// 	}
// 	fmt.Println()
// }
// func main() {
// fmt.Println("Введите два числа диапазона чисел Армстронга:")
// var a, b int
// fmt.Scanln(&a, &b)
// 	fmt.Println("Числа Армстронга в диапазоне от", a, "до", b, ":")
// 	findArmstrong(a, b)
// }

// #4 !!!
// package main
// import "fmt"

// func reverseString(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 	 runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
//    }

// func gcd(a, b int) int {
// 	for b != 0 {
// 	 a, b = b, a%b
// 	}
// 	return a
//    }

//	func main() {
//		var s string
//		fmt.Scanln(&s)
//		fmt.Println("Введите строку: ")
//		fmt.Println("Реверс строки \"s\":", reverseString(s))
//	}
//
// №5
// package main

// import "fmt"

// func gcd(a, b int) int {
// 	for b != 0 {
// 		a, b = b, a%b
// 	}
// 	return a
// }
// func main() {
// 	fmt.Println("Введите два числа для нахождения НОДа:")
// 	var a, b int
// 	fmt.Scanln(&a, &b)
// 	fmt.Println("НОД чисел a и b: ", gcd(a, b))
// }
