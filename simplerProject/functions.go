package main

import "fmt"

func printUserInfo(name string, age int) {
	fmt.Printf("so now your info is : \nname : %s \nage:%d  \n", name, age)

}

func CheckEvenOrOdd(number int) {
	if number%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

}
func UserInfo() {
	fmt.Println("what is your name ?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("age ? ")
	var age int
	fmt.Scanln(&age)
	printUserInfo(name, age)
}

func CheckEvenOrOddFHandler() {
	fmt.Println("give me a number")
	var number int
	fmt.Scanln(&number)
	CheckEvenOrOdd(number)
}

func SumAndMultiplication(a, b int) (int, int) {
	return a + b, a * b
}
func SumAndMultiplicationHandler() {
	fmt.Println("give me two number for sum and multiplication")
	var a, b int
	fmt.Scan(&a, &b)
	sum, multi := SumAndMultiplication(a, b)
	fmt.Println("Sum :", sum)
	fmt.Println("Multiplication :", multi)
}

func FunctionBased() {
	fmt.Println("Select which function you want to run :")
	fmt.Println("1. print user info")
	fmt.Println("2. check even or odd number")
	fmt.Println("3. give number and print sum and multiplication of two number")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		UserInfo()
	case 2:
		CheckEvenOrOddFHandler()
	case 3:
		SumAndMultiplicationHandler()
	}

}
