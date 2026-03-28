package main

import "fmt"

func EvenNumberPrint() {
	fmt.Println("Give me a number for Even Number Print till ")
	var n int
	fmt.Scanln(&n)
	var oddNumber, evenNumbers int
	fmt.Println("Here a Even numbers there :")
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			evenNumbers += 1
		} else {
			oddNumber += 1
		}
	}
	fmt.Printf("odd number : %d\n", oddNumber)
	fmt.Printf("Even number : %d", evenNumbers)
}
