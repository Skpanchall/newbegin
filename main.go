package main

import "fmt"

func main() {
	var choice int

	fmt.Println("1. Bio Maker")
	fmt.Println("2. Age Checker")
	fmt.Println("3. Work Status Checker")
	fmt.Println("Enter choice:")

	fmt.Scanln(&choice)

	if choice == 1 {
		BioMaker()
	} else if choice == 2 {
		Agechecker()
	} else if choice == 3 {
		WorkStatusChecker()
	} else {
		fmt.Println("Invalid choice")
	}

}
