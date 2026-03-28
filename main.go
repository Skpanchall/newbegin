package main

import "fmt"

func main() {
	var choice int

	fmt.Println("1. Bio Maker")
	fmt.Println("2. Age Checker")
	fmt.Println("3. Work Status Checker")
	fmt.Println("3. Give Numbers and Print Even and Odd numbers")
	fmt.Println("Enter choice:")

	fmt.Scanln(&choice)

	if choice == 1 {
		BioMaker()
	} else if choice == 2 {
		Agechecker()
	} else if choice == 3 {
		WorkStatusChecker()
	} else if choice == 4 {
		EvenNumberPrint()
	} else {
		fmt.Println("Invalid choice")
	}

}
