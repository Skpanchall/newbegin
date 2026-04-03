package main

import "fmt"

func main() {
	var choice int

	fmt.Println("1. Bio Maker")
	fmt.Println("2. Age Checker")
	fmt.Println("3. Work Status Checker")
	fmt.Println("4. Give Numbers and Print Even and Odd numbers")
	fmt.Println("5. functions")
	fmt.Println("6. arrya and slices")
	fmt.Println("7. sturcts")
	fmt.Println("8. Profile update")
	fmt.Println("9. API User Processor")
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
	} else if choice == 5 {
		FunctionBased()
	} else if choice == 6 {
		ArrayAndSlices()
	} else if choice == 7 {
		GetProductTotalValues()
	} else if choice == 8 {
		GetProfile()
	} else if choice == 9 {
		ApiUserProcess()
	} else {
		fmt.Println("Invalid choice")
	}

}
