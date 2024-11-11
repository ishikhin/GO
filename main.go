
// 1

// package main
// import "fmt"

// var a int32 = 9
// var b = 90.0 //если не писать int(число) то присваивается макс кол во битов
// var s1, s2, s3 string

// var (
// 	c1 string = "Hello"
// 	c2 int = 56
// )

// var isGood = true

// func main() {
// 	c := 7 //только внутри функции
// 	fmt.Println(c)
// }





// 2

// package main
// import "fmt"

// var a int = 3
// var b int = 3
// var c = a / b //для интов это целое деление, для флоат = флоат. Разные разрядности делить нельзя.

// func main() {
	
// 	if a > b {
// 		fmt.Println("a > b")
// 	} else if a < b {
// 		fmt.Println("a < b")
// 	} else {
// 		fmt.Println("a = b")
// 	}
// }



// 3

// package main
// import "fmt"

// var a int = 3


// func main() {
// 	switch a {
// 	case 1:
// 		fmt.Println("1")

// 	case 2:
// 		fmt.Println("1")

// 	case 3:
// 		fmt.Println("Cреда")
// 	}
// 	}


// 4

// package main
// import "fmt"

// var numArray = [5]int{}  // [5]int{1, 2, 3, 4, 5} ЛИБО автоматическое определение размера - [...]int{}


// func main() { //Динамический массив:
// 	numArray[0] = 1
// 	numArray[1] = 2
// 	numArray[2] = 3
// 	numArray[3] = 4
// 	numArray[4] = 5
// 	fmt.Println(numArray) //fmt.Println(numArray[1]) 

// 	}




// 5

// package main
// import "fmt"


// func main() { 
// 	for i := 0; i < 100; i++ {
// 		if i%7 == 0 {
// 			fmt.Println(i)
// 		}
// 	}

// 	}

// 6
// package main
// import "fmt"

// var numArray = [5]int{1, 2, 3, 4, 5} 


// func main() {
// 	for i := 0; i < len(numArray); i++ {
// 		fmt.Println(numArray[i])
// 		}
// 	}

